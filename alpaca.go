package alpaca

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/buger/jsonparser"
)

type AlpacaOptions struct {
	Schema  string
	Data    string
	Request *http.Request
}

var DefaultSchemaFieldMapping = map[string]string{
	"object": "object",
}

var DefaultFormatFieldMapping = map[string]string{}

type Alpaca struct {
	data          []byte
	schema        []byte
	options       []byte
	connector     string
	request       *http.Request
	FieldRegistry []*Field
}

type Field struct {
	Initializing         bool
	Parent               Alpaca
	Data                 []byte
	Options              []byte
	Schema               []byte
	Connector            *Field
	SingleLevelRendering string
	ID                   string
	Key                  string
	Title                string
	Name                 string
	NameCalculated       string
	Type                 []byte
	Path                 string
	Validation           string
	Events               string
	ShowingDefaultData   string
	PreviouslyValidated  bool
	IsContainerField     bool
	PropertyId           string
	Value                interface{}
	ChildrenByPropertyId map[string]Alpaca
	childrenById         map[string]Alpaca
	Children             []Alpaca
	Debug                bool
	Order                int
	ReadOnly             bool
	notTopLevel          bool
	// Media                []ImageFile
}

var (
	ErrDefaultError = errors.New("You must supply at least one argument.")
)

// Quote surrounds a string in quotes so jsonparser can detect that it is a string
func Quote(a string) string {
	return "\"" + a + "\""
}

// New initalizes and returns new alpaca parser
func New(options AlpacaOptions) (*Alpaca, error) {

	if (AlpacaOptions{}) == options {
		return nil, ErrDefaultError
	}

	alpaca := &Alpaca{}

	schemaValue, schemaType, _, _ := jsonparser.Get([]byte(options.Schema), "schema")
	if schemaType != jsonparser.NotExist {
		alpaca.schema = schemaValue
	}

	optionsValue, optionsType, _, _ := jsonparser.Get([]byte(options.Schema), "options")
	if optionsType != jsonparser.NotExist {
		alpaca.options = optionsValue
	} else {
		alpaca.options = []byte("{}")
	}

	dataValue, dataType, _, _ := jsonparser.Get([]byte(options.Data))
	if dataType != jsonparser.NotExist {
		alpaca.data = dataValue
	} else {
		alpaca.data = []byte("{}")
	}

	// Kick off the field registration
	alpaca.CreateFieldInstance("", alpaca.data, alpaca.options, alpaca.schema, nil)

	return alpaca, nil
}

// Parse takes field registry and parses it into json
func (a *Alpaca) Parse() {

}

// CreateFieldInstance returns a new instance of the desired field based on the schema
func (a *Alpaca) CreateFieldInstance(key string, data []byte, options []byte, schema []byte, connector *Field) {
	fieldType := ""

	_, optionsType, _, _ := jsonparser.Get(options, "type")
	if optionsType == jsonparser.NotExist {

		// if nothing passed in, we can try to make a guess based on the type of data
		_, schemaType, _, _ := jsonparser.Get(schema, "type")
		if schemaType == jsonparser.NotExist {
			_, dataType, _, _ := jsonparser.Get(data)
			if dataType != jsonparser.NotExist {
				jsonparser.Set(schema, []byte(string(dataType)), "type")
			}
		}

		// if nothing passed in, fallback to defaults
		if schemaType == jsonparser.NotExist {
			fieldType = "object"
		}

		// using what we now about schema, try to guess the type
		guessedOptionType := a.GuessOptionsType(schema)
		if guessedOptionType != "" {
			options, _ = jsonparser.Set(options, []byte(Quote(guessedOptionType)), "type")
		}
	}

	optionsValue, optionsValueType, _, _ := jsonparser.Get(options, "type")

	if optionsValueType != jsonparser.NotExist {
		fieldType = string(optionsValue)
	}
	// TODO Add non container fields to field registry

	switch fieldType {
	case "object":

		a.Object(schema, options, data, connector)
		break
	case "information":

		break
	default:
		a.Any()
		break
	}
}

func (a *Alpaca) ResolvePropertySchemaOptions(key string, connector *Field) {

	aSchema := make([]byte, 0)
	propertiesValue, propertiesType, _, _ := jsonparser.Get(connector.Schema, "properties")
	if propertiesType != jsonparser.NotExist {
		propertyValue, propertyType, _, _ := jsonparser.Get(propertiesValue, key)
		if propertyType != jsonparser.NotExist {
			aSchema = propertyValue
		}
	}

	aOptions := make([]byte, 0)
	aOptions = connector.Options
	optionsValue, optionsType, _, _ := jsonparser.Get(aOptions, "fields")
	if optionsType != jsonparser.NotExist {
		propertyOptions, propertyOptionsType, _, _ := jsonparser.Get(optionsValue, string(key))
		if propertyOptionsType != jsonparser.NotExist {
			aOptions = propertyOptions
		}
	}

	// If field is found use that otherwise dive deeper
	aData := make([]byte, 0)
	aData = connector.Data
	propertyDataValue, propertyDataType, _, _ := jsonparser.Get(aData, string(key))
	if propertyDataType != jsonparser.NotExist {
		aData = propertyDataValue
	}

	a.CreateFieldInstance(key, aData, aOptions, aSchema, connector)
}

// Object container field
func (a *Alpaca) Object(schema []byte, options []byte, data []byte, connector *Field) {

	field := &Field{
		Schema:  schema,
		Options: options,
		Data:    data,
	}
	a.FieldRegistry = append(a.FieldRegistry, field)

	propertiesValue, propertiesType, _, _ := jsonparser.Get(schema, "properties")
	if propertiesType != jsonparser.NotExist {
		jsonparser.ObjectEach(propertiesValue, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			a.ResolvePropertySchemaOptions(string(key), field)
			return nil
		})
	}
}

// Information container field
func (a *Alpaca) Information() *Field {
	field := &Field{}
	fmt.Println("information")
	a.FieldRegistry = append(a.FieldRegistry)
	return field
}

// Any default fallback field
func (a *Alpaca) Any() *Field {
	field := &Field{}
	fmt.Println("any")
	a.FieldRegistry = append(a.FieldRegistry)
	return field
}

// GuessOptionsType determines field type
func (a *Alpaca) GuessOptionsType(schema []byte) string {
	optionsType := ""

	enumValue, enumType, _, _ := jsonparser.Get(schema, "enum")

	if enumType != jsonparser.NotExist {
		if enumType == jsonparser.Array {
			if len(enumValue) > 3 {
				optionsType = "select"
			} else {
				optionsType = "radio"
			}
		}
	} else {
		schemaTypeValue, schemaType, _, _ := jsonparser.Get(schema, "type")

		if schemaType != jsonparser.NotExist {
			mapValue, isset := DefaultSchemaFieldMapping[string(schemaTypeValue)]
			if isset {
				optionsType = mapValue
			}
		}
	}

	// check if it has format defined
	schemaFormatValue, schemaFormatType, _, _ := jsonparser.Get(schema, "format")
	if schemaFormatType != jsonparser.NotExist {
		mapValue, isset := DefaultFormatFieldMapping[string(schemaFormatValue)]
		if isset {
			optionsType = mapValue
		}
	}

	return optionsType
}
