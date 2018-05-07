package alpaca

import (
	"bytes"
	"encoding/hex"
	"image"
	"io"
	"mime"
	"strconv"
	"time"

	"github.com/bradfitz/slice"
	"github.com/spf13/cast"

	"github.com/Jeffail/gabs"
)

// New initalizes and returns new alpaca parser
func New(options AlpacaOptions) (*Alpaca, error) {

	if (AlpacaOptions{}) == options {
		return nil, ErrDefaultError
	}

	alpaca := &Alpaca{}

	schema, err := gabs.ParseJSON([]byte(options.Schema))
	if err != nil {
		return nil, ErrSchemaInvalid
	}

	data, err := gabs.ParseJSON([]byte(options.Data))
	if err != nil {
		return nil, ErrDataInvalid
	}

	alpaca.schema = schema.Search("schema")
	alpaca.options = schema.Search("options")
	if alpaca.options == nil {
		alpaca.options = gabs.New()
	}
	alpaca.data = data

	if options.Request != nil {
		alpaca.request = options.Request
	}

	// Kick off the field registration
	alpaca.CreateFieldInstance("", alpaca.data, alpaca.options, alpaca.schema, nil, 0, false)

	// Sort fields by ordering
	slice.Sort(alpaca.FieldRegistry[:], func(i, j int) bool {
		return alpaca.FieldRegistry[i].Order < alpaca.FieldRegistry[j].Order
	})

	return alpaca, nil
}

// ResolveItemSchemaOptions resolves the items in an array container field
func (a *Alpaca) ResolveItemSchemaOptions(key string, connector *Field, index int) {
	schema := gabs.New()
	if connector.Schema.Exists("items") {
		schema = connector.Schema.S("items")
	}

	options := gabs.New()
	if connector.Options.Exists("items") {
		options = connector.Options.S("items")
	}

	data := connector.Data.Index(index)

	a.CreateFieldInstance(cast.ToString(index), data, options, schema, connector, index, true)
}

// ResolvePropertySchemaOptions resolves the properties in an object container field
func (a *Alpaca) ResolvePropertySchemaOptions(key string, connector *Field) {

	schema := gabs.New()
	if connector.Schema.Exists("properties") && connector.Schema.S("properties").Exists(key) {
		schema = connector.Schema.S("properties").S(key)
	}

	options := gabs.New()
	options = connector.Options
	if connector.Options.Exists("fields") {
		if connector.Options.S("fields").Exists(key) {
			options = connector.Options.S("fields").S(key)
		} else {
			options = connector.Options.S("fields")
		}
	}

	data := gabs.New()
	if connector.Data.Exists(key) {
		data = connector.Data.S(key)
	}

	a.CreateFieldInstance(key, data, options, schema, connector, 0, false)
}

// GuessOptionsType determines field type
func (a *Alpaca) GuessOptionsType(schema *gabs.Container) string {
	optionType := ""

	if schema.Exists("enum") == true {

		children, err := schema.S("enum").Children()
		if err == nil {
			if len(children) > 3 {
				optionType = "select"
			} else {
				optionType = "radio"
			}
		}

	} else {
		fieldType := schema.S("type").Data()

		if fieldType != nil {
			if fieldType.(string) != "" {
				optionType = DefaultSchemaFieldMapping[fieldType.(string)]
			}
		}
	}

	// check if it has format defined
	if schema.Exists("format") == true {
		optionType = DefaultFormatFieldMapping[schema.S("format").Data().(string)]
	}

	return optionType
}

// GetSchemaType returns schema type of data.
func (a *Alpaca) GetSchemaType(data *gabs.Container) string {
	// seems to be returning an array even for strings, thus invalid
	if _, err := data.Children(); err == nil {
		return "array"
	}

	if _, err := data.ChildrenMap(); err == nil {
		return "object"
	}

	if _, err := strconv.Atoi(data.String()); err == nil {
		return "number"
	}

	if _, err := strconv.ParseBool(data.String()); err == nil {
		return "boolean"
	}

	return "string"
}

// GetAttributes extracts generic attributes from fields
func (f *Field) GetAttributes() {
	if f.Schema.Exists("title") {
		f.Title = cast.ToString(f.Schema.S("title").Data())
	}

	// So parent can see children as well as the children seeing parent
	if f.Parent != nil {
		f.Parent.Children = append(f.Parent.Children, f)
	}

	if f.Options.Exists("order") {
		f.Order = cast.ToInt(f.Options.S("order").Data())
	}

	if f.Data.Data() != nil {
		f.Value = f.Data.Data()
	}

	if f.Schema.Exists("readonly") {
		f.ReadOnly = cast.ToBool(f.Schema.S("readonly").Data())
	}

	if f.Schema.Exists("default") {
		f.Default = f.Schema.S("default").Data()
	}

	if f.ReadOnly {
		f.Value = f.Default
	}
}

// RegisterField field adds the field to the field registry
func (a *Alpaca) RegisterField(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

func (a *Alpaca) RegisterMedia(f *Field, index int) {

	// This won't work since f.Key ignores all the chunks before hand. We need to regen the entire path
	fileName := f.Key + "_image_" + strconv.Itoa(index)

	file, _, err := a.request.FormFile(fileName)
	CreatedDevice := a.request.FormValue(fileName + "_created")

	if err == nil {
		defer file.Close()

		foundFile := ImageFile{}
		var Buf bytes.Buffer
		io.Copy(&Buf, file)
		contents := Buf.Bytes()
		content := hex.EncodeToString(contents)
		foundFile.Data = content

		file, _, _ := a.request.FormFile(fileName)
		config, format, _ := image.DecodeConfig(file)
		foundFile.Name = fileName
		foundFile.Width = config.Width
		foundFile.Height = config.Height
		foundFile.Type = format
		foundFile.Mime = mime.TypeByExtension("." + format)
		foundFile.FieldKey = f.Key
		foundFile.FieldRef = f

		layout := "2006-01-02 15:04:05"
		t, err := time.Parse(layout, CreatedDevice)
		if err != nil {
			foundFile.Created = time.Now()
		} else {
			foundFile.Created = t
		}

		a.MediaRegistry = append(a.MediaRegistry, foundFile)
		f.Media = append(f.Media, foundFile)

		Buf.Reset()
	}
}

// CreateFieldInstance returns a new instance of the desired field based on the schema
func (a *Alpaca) CreateFieldInstance(key string, data *gabs.Container, options *gabs.Container, schema *gabs.Container, connector *Field, arrayIndex int, arrayChild bool) {

	fieldType := ""

	if options.Exists("type") == false {

		// if nothing passed in, we can try to make a guess based on the type of data
		if schema.Exists("type") == false {
			schemaType := a.GetSchemaType(data)

			if schemaType != "" {
				fieldType = schemaType
			}
		}

		// if nothing passed in, fallback to defaults
		if schema.Exists("type") == false {
			fieldType = "object" // fallback
		}

		optionType := a.GuessOptionsType(schema)
		if optionType != "" {
			fieldType = optionType
		}
	} else {
		fieldType = options.S("type").Data().(string)
	}

	f := &Field{
		Schema:       schema,
		Options:      options,
		Data:         data,
		Key:          key,
		Type:         fieldType,
		Parent:       connector,
		IsArrayChild: arrayChild,
		ArrayIndex:   arrayIndex,
		ArrayValues:  0,
	}

	f.GetAttributes()

	if connector != nil {
		for _, chunk := range connector.Path {
			f.Path = append(f.Path, chunk)
		}
	}

	if f.IsArrayChild && f.Value != nil {
		f.Parent.ArrayValues++
	}

	f.Path = append(f.Path, Chunk{Type: f.Type, Value: f.Key, Field: f})

	for i := range f.Path {
		if i > 0 {
			f.Path[i-1].Connector = &f.Path[i]
			f.Path[i].Parent = &f.Path[i-1]
		}
	}

	// Not all field types are required for definition, many share the same basic behaviour as Any
	switch f.Type {
	case "array", "repeatable":
		a.Array(f)
		break
	case "object":
		a.Object(f)
		break
	case "camera":
		a.Camera(f)
		break
	case "lowercase":
		a.Lowercase(f)
		break
	case "uppercase":
		a.Uppercase(f)
		break
	case "information", "image":
		a.Information(f)
		break
	case "signature":
		a.Signature(f)
		break
	case "editor":
		a.Editor(f)
		break
	case "json":
		a.JSON(f)
		break
	default:
		a.Any(f)
	}
}
