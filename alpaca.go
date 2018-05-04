package alpaca

import (
	"strconv"

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

func (a *Alpaca) ParseFieldPath(f *Field, chunk *Chunk) *gabs.Container {
	result := gabs.New()

	switch chunk.Type {
	case "array":
		if chunk.Connector != nil {
			result.Array(chunk.Value)
			result.ArrayAppend(a.ParseFieldPath(f, chunk.Connector).Data(), chunk.Value)
		}
		break
	case "object":
		if chunk.Connector != nil {
			isInt := false
			if _, err := strconv.Atoi(chunk.Value); err == nil {
				isInt = true
			}

			if chunk.Value != "" && !isInt {
				result.Set(a.ParseFieldPath(f, chunk.Connector).Data(), chunk.Value)
			} else {
				return a.ParseFieldPath(f, chunk.Connector)
			}
		}
		break
	default:
		result.Set(f.Value, chunk.Value)
	}

	return result
}

// Parse takes field registry and parses it into json string
func (a *Alpaca) Parse() string {
	result := gabs.New()

	if len(a.FieldRegistry) < 2 {
		return `"` + cast.ToString(a.FieldRegistry[0].Value) + `"`
	}

	results := make([]*gabs.Container, 0)
	for _, f := range a.FieldRegistry {
		if f.Value != nil && cast.ToString(f.Value) != "" {
			results = append(results, a.ParseFieldPath(f, &f.Path[0]))
		}
	}

	for _, generated := range results {
		result.Merge(generated)
	}

	return result.String()
}

// PathString returns combined path string - decrepit
func (f *Field) PathString(depth int) (path string, chunks []string) {
	result := make([]string, 0)
	strResult := ""

	for _, chunk := range f.Path {
		if chunk.Type != "object" {
			if strResult != "" {
				strResult += "." + chunk.Value
				result = append(result, chunk.Value)
			} else {
				strResult = chunk.Value
				result = append(result, chunk.Value)
			}
		} else if chunk.Type == "array" {
			strResult += "[" + chunk.Value + "]"
		}
	}

	return strResult, result
}

func (a *Alpaca) RegisterField(f *Field) {
	a.FieldRegistry = append(a.FieldRegistry, f)
}

func (a *Alpaca) ResolveItemSchemaOptions(key string, connector *Field, index int) {
	schema := gabs.New()
	if connector.Schema.Exists("items") {
		schema = connector.Schema.S("items")
	}

	options := connector.Options
	data := connector.Data.Index(index)

	a.CreateFieldInstance(cast.ToString(index), data, options, schema, connector, index, true)
}

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

	// Here we can do some fancy logic if the field is an array field
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

	f.Path = append(f.Path, Chunk{Type: f.Type, Value: f.Key, Field: f})

	for i, _ := range f.Path {
		if i > 0 {
			f.Path[i-1].Connector = &f.Path[i]
		}
	}

	// Not all field types are required for definition, many share the same basic behaviour as Any
	switch f.Type {
	case "array":
		a.Array(f)
		break
	case "object":
		a.Object(f)
		break
	case "camera":
		a.Camera(f)
		break
	case "information":
		a.Information(f)
		break
	case "signature":
		a.Signature(f)
		break
	default:
		a.Any(f)
	}
}
