package alpaca

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"image"
	"io"
	"math"
	"mime"
	"os"
	"strconv"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

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

	for _, field := range alpaca.FieldRegistry {
		if field.Parent != nil && field.Parent.IsArrayChild {
			field.IsArrayChild = true
		}
	}

	// Sort fields by ordering - This won't work as it ignores array ordering
	slice.Sort(alpaca.FieldRegistry[:], func(i, j int) bool {
		return alpaca.FieldRegistry[i].DepthOrder < alpaca.FieldRegistry[j].DepthOrder
	})

	return alpaca, nil
}

// ResolveItemSchemaOptions resolves the items in an array container field
func (a *Alpaca) ResolveItemSchemaOptions(key string, connector *Field, index int) {

	isInt := false
	if _, err := strconv.Atoi(key); err == nil {
		isInt = true
	}

	schema := gabs.New()
	if isInt {
		schema = connector.Schema
	}
	if connector.Schema.Exists("items") {
		schema = connector.Schema.S("items")
	}

	options := gabs.New()
	if isInt {
		options = connector.Options
	}
	if connector.Options.Exists("items") {
		options = connector.Options.S("items")
	}

	data := connector.Data.Index(index)
	if isInt {
		data = connector.Data
	}

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

// GetDepthOrder
func GetDepthOrder(f *Field, o float64) float64 {
	if f.Parent != nil {
		o += GetDepthOrder(f.Parent, o)
	}

	if v, err := strconv.Atoi(f.Key); err == nil {
		o += float64(v) / math.Pow(10, float64(f.Depth))
	} else {
		o += f.Order / math.Pow(10, float64(f.Depth))
	}

	return o
}

// GetAttributes extracts generic attributes from fields
func (f *Field) GetAttributes() {

	if f.Default != "" {
		f.Value = f.Default
	}

	if f.Schema.Exists("title") {
		f.Title = cast.ToString(f.Schema.S("title").Data())
	}

	// So parent can see children as well as the children seeing parent
	if f.Parent != nil {
		f.Parent.Children = append(f.Parent.Children, f)
	}

	if f.Parent != nil {
		f.Depth = f.Parent.Depth
		f.Depth++
	}

	if f.Options.Exists("order") {
		// f.Order += cast.ToFloat64(f.Options.S("order").Data()) / math.Pow(10, float64(f.Depth))
		f.Order = cast.ToFloat64(f.Options.S("order").Data()) // / math.Pow(10, float64(f.Depth))
	}

	f.DepthOrder = GetDepthOrder(f, 0)

	if f.Data.Data() != nil {
		f.Value = f.Data.Data()
	}

	if f.Schema.Exists("readonly") {
		f.ReadOnly = cast.ToBool(f.Schema.S("readonly").Data())
	}

	if f.Schema.Exists("default") {
		f.Default = f.Schema.S("default").Data()
	}

	// if f.ReadOnly {
	// 	f.Value = f.Default
	// }
}

// RegisterField field adds the field to the field registry
func (a *Alpaca) RegisterField(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

func (a *Alpaca) RegisterMedia(f *Field, index int) {

	fileName := f.PathString + "_image_" + strconv.Itoa(index)
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
		config, format, err := image.DecodeConfig(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", fileName, err)
		}

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

		fmt.Println(foundFile.Type)
		fmt.Println(foundFile.Mime)
		fmt.Println(foundFile.Width)
		fmt.Println(foundFile.Height)

		a.MediaRegistry = append(a.MediaRegistry, foundFile)
		f.Media = append(f.Media, foundFile)

		Buf.Reset()
	}
}

// CreateFieldInstance returns a new instance of the desired field based on the schema
func (a *Alpaca) CreateFieldInstance(key string, data *gabs.Container, options *gabs.Container, schema *gabs.Container, connector *Field, arrayIndex int, arrayChild bool) {

	optionsType := ""
	schemaType := ""

	if schema.Exists("type") != false {
		schemaType = schema.S("type").Data().(string)
	}

	if options.Exists("type") == false {

		// if nothing passed in, we can try to make a guess based on the type of data
		if schema.Exists("type") == false {
			schemaType = a.GetSchemaType(data)

			if schemaType != "" {
				optionsType = schemaType
			}
		}

		// if nothing passed in, fallback to defaults
		if schema.Exists("type") == false {
			optionsType = "object" // fallback
		} else {
			schemaType = schema.S("type").Data().(string)
		}

		optionType := a.GuessOptionsType(schema)
		if optionType != "" {
			optionsType = optionType
		}
	} else {
		optionsType = options.S("type").Data().(string)
	}

	f := &Field{
		Schema:       schema,
		Options:      options,
		Data:         data,
		DataString:   data.String(),
		Key:          key,
		Type:         optionsType,
		SchemaType:   schemaType,
		ChunkType:    optionsType,
		Parent:       connector,
		IsArrayChild: arrayChild,
		ArrayIndex:   arrayIndex,
		ArrayValues:  0,
	}

	if optionsType == "select" && schemaType != "" {
		f.ChunkType = schemaType
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

	f.Path = append(f.Path, Chunk{Type: f.ChunkType, Value: f.Key, Field: f})

	for i := range f.Path {
		if i > 0 {
			f.Path[i-1].Connector = &f.Path[i]
			f.Path[i].Parent = &f.Path[i-1]
		}
	}

	f.PathString = f.GetPathString()

	// Not all field types are required for definition, many share the same basic behaviour as Any
	switch f.Type {
	case "array", "repeatable", "select", "checkbox":
		a.Array(f)
		break
	case "datetime":
		a.Datetime(f)
		break
	case "object":
		a.Object(f)
		break
	case "tag":
		a.Tag(f)
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
