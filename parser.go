package alpaca

import (
	"strconv"

	"github.com/Jeffail/gabs"
	"github.com/spf13/cast"
)

// ParseFieldPath reconstructs JSON based on the field path
func (a *Alpaca) ParseFieldPath(f *Field, chunk *Chunk, generated *gabs.Container) *gabs.Container {

	switch chunk.Type {
	case "array":
		if chunk.Connector != nil {
			if !generated.Exists(chunk.Value) {
				generated.ArrayOfSize(chunk.Field.ArrayValues, chunk.Value)
			}

			arrayVal := generated.S(chunk.Value)
			if chunk.Connector != nil {
				a.ParseFieldPath(f, chunk.Connector, arrayVal)
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
			if generated.Index(intVal).Data() != nil && chunk.Parent.Type == "object" {
				item.Merge(generated.Index(intVal))
			}

			generated.SetIndex(item.Data(), intVal)

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

	return generated
}

// Parse takes field registry and parses it into json string
func (a *Alpaca) Parse() string {
	result := gabs.New()

	if len(a.FieldRegistry) < 2 {
		if a.FieldRegistry[0].IsContainerField {
			return `""`
		}

		switch v := a.FieldRegistry[0].Value.(type) {
		case int:
			return cast.ToString(v)
		case float64:
			return cast.ToString(v)
		default:
			return `"` + cast.ToString(v) + `"`
		}

	}

	for _, f := range a.FieldRegistry {
		if f.Value != nil && cast.ToString(f.Value) != "" {
			a.ParseFieldPath(f, &f.Path[0], result)
		}
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
