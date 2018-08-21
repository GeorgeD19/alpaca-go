package alpaca

import (
	"strconv"

	"github.com/Jeffail/gabs"
	"github.com/spf13/cast"
)

// ParseFieldPath reconstructs JSON based on the field path
func (a *Alpaca) ParseFieldPath(f *Field, chunk *Chunk, generated *gabs.Container) *gabs.Container {

	switch chunk.Type {
	case "repeatable", "array", "select":
		if chunk.Connector != nil {
			if chunk.Field.ArrayValues > 0 {

				if chunk.Value != "" {
					if !generated.Exists(chunk.Value) {
						generated.ArrayOfSize(chunk.Field.ArrayValues, chunk.Value)
					}

					arrayVal := generated.S(chunk.Value)
					if chunk.Connector != nil {
						if chunk.Connector.Type != "array" && chunk.Connector.Type != "select" && chunk.Connector.Type != "repeatable" && chunk.Connector.Type != "object" {
							intVal := 0
							if v, err := strconv.Atoi(chunk.Connector.Value); err == nil {
								intVal = v
							}

							arrayVal.SetIndex(f.Value, intVal)
						} else {
							a.ParseFieldPath(f, chunk.Connector, arrayVal)
						}
					}
				} else {
					generated.Set(f.Parent.Data.Data())
				}

			}
		}
		break
	case "object":
		isInt := false
		intVal := 0
		if v, err := strconv.Atoi(chunk.Value); err == nil {
			intVal = v
			isInt = true
		}

		if isInt {
			arrayValue := gabs.New()

			item := a.ParseFieldPath(f, chunk.Connector, arrayValue)
			itemParsed, _ := gabs.ParseJSON([]byte(item.String()))
			indexData := generated.Index(intVal).Data()
			// This shouldn't work, but it does. Something wrong with chunk types
			if indexData != nil && (chunk.Parent.Type == "object" || chunk.Parent.Type == "repeatable" || chunk.Parent.Type == "array") {
				item2Parsed, _ := gabs.ParseJSON([]byte(generated.Index(intVal).String()))
				itemParsed.Merge(item2Parsed)
			}
			generated.SetIndex(itemParsed.Data(), intVal)

		} else {
			if chunk.Connector != nil {
				if chunk.Value != "" {
					generated.Set(a.ParseFieldPath(f, chunk.Connector, generated.S(chunk.Value)).Data(), chunk.Value)
				} else {
					return a.ParseFieldPath(f, chunk.Connector, generated)
				}
			}
		}
		break
	default:
		if generated == nil {
			generated = gabs.New()
		}
		generated.Set(f.Value, chunk.Value)
	}
	// TODO Merge with existing arrays e.g. [{"location": "test"}] and [{"control": "test"}] becomes [{"control":"test","location": "test"}]
	// fmt.Println(generated.String())
	return generated
}

// Parse takes field registry and parses it into json string
func (a *Alpaca) Parse() string {

	if a.output == "" {

		result := gabs.New()

		if len(a.FieldRegistry) < 2 {
			if a.FieldRegistry[0].IsContainerField {
				return `""`
			}

			// return cast.ToString(a.FieldRegistry[0].Value)

			switch v := a.FieldRegistry[0].Value.(type) {
			case int:
				return cast.ToString(v)
			case float64:
				return cast.ToString(v)
			default:
				if a.FieldRegistry[0].Type != "select" && a.FieldRegistry[0].Type != "json" && a.FieldRegistry[0].Type != "tag" {
					return `"` + cast.ToString(v) + `"`
				}
				return cast.ToString(v)

			}

		}

		for _, f := range a.FieldRegistry {
			// fmt.Println(f.PathString)
			if f.Value != nil && cast.ToString(f.Value) != "" {
				a.ParseFieldPath(f, &f.Path[0], result)
			}
		}

		a.output = result.String()
	}

	return a.output
}

// GetPathString returns combined path string - decrepit
func (f *Field) GetPathString() (path string) {
	result := ""
	fPath := f.Path
	for _, chunk := range fPath {
		if chunk.Parent != nil {
			parent := chunk.Parent
			if parent.Type == "array" || parent.Type == "repeatable" {
				result += "[" + chunk.Value + "]"
			} else {
				if result != "" {
					result += "." + chunk.Value
				} else {
					result = chunk.Value
				}
			}
		}
	}

	return result
}
