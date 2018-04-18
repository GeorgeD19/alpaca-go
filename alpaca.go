package alpaca

import (
	"fmt"

	"github.com/spf13/cast"

	"github.com/buger/jsonparser"
)

// We want to build a list of fields

// Then we want to validate each fields dependencies

func FieldContainer() {

}

func FieldAny() {

}

// GetSchemaType returns schema type of a variable.
func GetSchemaType(data string) string {
	var schemaType = "any"

	// if isArray(Data) == true {
	// 	SchemaType = "array"
	// } else if isObject(Data) == true {
	// 	SchemaType = "object"
	// } else if isNumber(Data) == true {
	// 	SchemaType = "number"
	// } else if isBoolean(Data) == true {
	// 	SchemaType = "boolean"
	// } else {
	// 	SchemaType = "string"
	// }

	return schemaType
}

func (f *Field) Construct() []*Field {
	fmt.Println(f.Type)
	switch f.Type {
	case "object":
		f.GetProperties()
		break
	}
}

func (f *Field) GetProperties() []*Field {
	properties, dataType, _, _ := jsonparser.Get([]byte(f.Schema), "properties")
	if dataType != jsonparser.NotExist {
		jsonparser.ObjectEach([]byte(properties), func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)

			return nil
		})
	}
}

func CreateFieldInstance(schema string, options string, data string) []*Field {
	field := new(Field)
	field.Schema = schema
	field.Options = options
	field.Data = data

	fieldType, dataType, _, _ := jsonparser.Get([]byte(options), "type")
	if dataType == jsonparser.NotExist {
		fieldType, dataType, _, _ = jsonparser.Get([]byte(schema), "type")
		if dataType == jsonparser.NotExist {
			// Rely on data to determine type
			fieldType = []byte(GetSchemaType(data))
		}
	}

	field.Type = cast.ToString(fieldType)
	fields = append(fields, field.Construct())

	return fields
}

func CheckDependencies() {

}

// New returns a pointer to a new Alpaca instance. Its methods are subsequently
// called to produce a single Alpaca object.
func New(logic string, data string) (a *Alpaca) {
	return AlpacaNew(logic, data)
}

func AlpacaNew(logic string, data string) (a *Alpaca) {
	a = new(Alpaca)
	schema := "{}"
	options := "{}"

	jsonparser.ObjectEach([]byte(logic), func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if string(key) == "schema" {
			schema = string(value)
		}

		if string(key) == "options" {
			options = string(value)
		}

		return nil
	})

	a.Fields = append(a.Fields, CreateFieldInstance(schema, options, data))

	return
}

// func Apply(logic string, data string) (a *Alpaca) {
// 	a = new(Alpaca)
// 	schema := "{}"
// 	options := "{}"

// 	jsonparser.ObjectEach([]byte(logic), func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
// 		if string(key) == "schema" {
// 			schema = string(value)
// 		}

// 		if string(key) == "options" {
// 			options = string(value)
// 		}

// 		a.CreateFieldInstance(schema, options, data)

// 		// fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
// 		return nil
// 	})

// 	return
// }
