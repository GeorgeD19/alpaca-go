package alpaca

import (
	"strconv"

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
	alpaca.CreateFieldInstance("", alpaca.data, alpaca.options, alpaca.schema, nil)

	return alpaca, nil
}

// Parse takes field registry and parses it into json
func (a *Alpaca) Parse() string {
	return "Nothing for you!"
}

// Validate takes field registry and validates it against passed data
func (a *Alpaca) Validate() {

}

func (a *Alpaca) ResolvePropertySchemaOptions(key string, connector *Field) {

	schema := gabs.New()
	if connector.Schema.Exists("properties") && connector.Schema.S("properties").Exists(key) {
		schema = connector.Schema.S("properties").S(key)
	}

	options := gabs.New()
	options = connector.Options
	if connector.Options.Exists("fields") && connector.Options.S("fields").Exists(key) {
		options = connector.Options.S("fields").S(key)
	}

	data := gabs.New()
	if connector.Data.Exists(key) {
		data = connector.Data.S(key)
	}

	a.CreateFieldInstance(key, data, options, schema, connector)
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
		fieldType := schema.Search("type").Data().(string)
		if fieldType != "" {
			optionType = DefaultSchemaFieldMapping[fieldType]
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

// CreateFieldInstance returns a new instance of the desired field based on the schema
func (a *Alpaca) CreateFieldInstance(key string, data *gabs.Container, options *gabs.Container, schema *gabs.Container, connector *Field) {

	if !options.Exists("type") {

		// if nothing passed in, we can try to make a guess based on the type of data
		if !schema.Exists("type") {
			schemaType := a.GetSchemaType(data)
			if schemaType != "" {
				schema.Set(schemaType, "type")
			}
		}

		// if nothing passed in, fallback to defaults
		if !schema.Exists("type") {
			schema.Set("object", "type")
		}

		// using what we now about schema, try to guess the type
		optionType := a.GuessOptionsType(schema)
		if optionType != "" {
			options.Set(optionType, "type")
		}
	}

	fieldType := options.Search("type").Data().(string)

	f := &Field{
		Schema:  schema,
		Options: options,
		Data:    data,
		Key:     key,
		Type:    fieldType,
		Parent:  connector,
	}
	f.GetAttributes()

	switch fieldType {
	case "address":
		a.Address(f)
		break
	case "ckeditor":
		a.CKEditor(f)
		break
	case "color":
		a.Color(f)
		break
	case "colorpicker":
		a.ColorPicker(f)
		break
	case "country":
		a.Country(f)
		break
	case "currency":
		a.Currency(f)
		break
	case "date":
		a.Date(f)
		break
	case "datetime":
		a.DateTime(f)
		break
	case "editor":
		a.Editor(f)
		break
	case "email":
		a.Email(f)
		break
	case "grid":
		a.Grid(f)
		break
	case "image":
		a.Image(f)
		break
	case "integer":
		a.Integer(f)
		break
	case "ipv4":
		a.IPv4(f)
		break
	case "json":
		a.JSON(f)
		break
	case "lowercase":
		a.Lowercase(f)
		break
	case "map":
		a.Map(f)
		break
	case "optiontree":
		a.OptionTree(f)
		break
	case "password":
		a.Password(f)
		break
	case "personalname":
		a.PersonalName(f)
		break
	case "phone":
		a.Phone(f)
		break
	case "pickacolor":
		a.PickAColor(f)
		break
	case "search":
		a.Search(f)
		break
	case "state":
		a.State(f)
		break
	case "summernote":
		a.Summernote(f)
		break
	case "table":
		a.Table(f)
		break
	case "tablerow":
		a.TableRow(f)
		break
	case "tag":
		a.Tag(f)
		break
	case "time":
		a.Time(f)
		break
	case "tinymce":
		a.TinyMCE(f)
		break
	case "token":
		a.Token(f)
		break
	case "upload":
		a.Upload(f)
		break
	case "uppercase":
		a.Uppercase(f)
		break
	case "url":
		a.URL(f)
		break
	case "zipcode":
		a.Zipcode(f)
		break
	case "array":
		a.Array(f)
		break
	case "file":
		a.File(f)
		break
	case "hidden":
		a.Hidden(f)
		break
	case "number":
		a.Number(f)
		break
	case "object":
		a.Object(f)
		break
	case "text":
		a.Text(f)
		break
	case "textarea":
		a.TextArea(f)
		break
	case "camera":
		a.Camera(f)
		break
	case "information":
		a.Information(f)
		break
	case "repeatable":
		a.Repeatable(f)
		break
	case "signature":
		a.Signature(f)
		break
	case "checkbox":
		a.Checkbox(f)
		break
	case "chooser":
		a.Chooser(f)
		break
	case "radio":
		a.Radio(f)
		break
	case "select":
		a.Select(f)
		break
	case "any":
		a.Any(f)
	default:
		a.Any(f)
	}
}
