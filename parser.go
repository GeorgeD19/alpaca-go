package alpaca

import (
	"fmt"
	"strconv"

	"github.com/Jeffail/gabs"
	"github.com/spf13/cast"
)

// FirstFieldPath returns the first chunk of the path we are currently on
func (a *Alpaca) FirstFieldPath(chunk *Chunk) *Chunk {
	if chunk.Parent != nil {
		return a.FirstFieldPath(chunk.Parent)
	}
	return chunk
}

// LastFieldPath returns the first chunk of the path we are currently on
func (a *Alpaca) LastFieldPath(chunk *Chunk) *Chunk {
	if chunk.Connector != nil {
		return a.LastFieldPath(chunk.Connector)
	}
	return chunk
}

func (a *Alpaca) DiveDeeper(chunk *Chunk, generated *gabs.Container) *gabs.Container {

	switch chunk.Type {
	case "array":
	case "object":
		// If integer we need to use index
		if generated.Exists(chunk.Value) {
			if chunk.Connector != nil {
				return a.DiveDeeper(chunk.Connector, generated.S(chunk.Value))
			} else {
				return generated.S(chunk.Value)
			}
		} else {
			if chunk.Connector != nil {
				return a.DiveDeeper(chunk.Connector, generated.S(chunk.Value))
			} else {
				return generated
			}
		}
		break
	default:
		if generated.Exists(chunk.Value) {
			return generated.S(chunk.Value)
		}
	}
	return generated

}

func (a *Alpaca) DiveFieldPath(chunk *Chunk, generated *gabs.Container) *gabs.Container {
	// generated will be starting at the left, and chunk will be starting at the right

	root := a.FirstFieldPath(chunk)

	// We get the root field for the current drunk and for every connector dive deeper down the generated container then we can affect it instead
	generated = a.DiveDeeper(root, generated)

	return generated
}

// ParseFieldPath reconstructs JSON based on the field path
func (a *Alpaca) ParseFieldPath(f *Field, chunk *Chunk, generated *gabs.Container) *gabs.Container {

	// Dive to the current position within the generated container
	// currentPath := a.DiveFieldPath(chunk, generated)
	// result := a.DiveFieldPath(chunk, generated)

	switch chunk.Type {
	case "array":
		if chunk.Connector != nil {
			if !generated.Exists(chunk.Value) {
				generated.ArrayOfSize(chunk.Field.ArrayValues, chunk.Value)
			}

			arrayVal := generated.S(chunk.Value)

			// fmt.Println(generated.String())
			if chunk.Connector != nil {
				a.ParseFieldPath(f, chunk.Connector, arrayVal)
			}
			// 		result.ArrayAppend(a.ParseFieldPath(f, chunk.Connector, generated).Data(), chunk.Value)
		}
		break
	case "object":
		isInt := false
		intVal := 0
		if v, err := strconv.Atoi(chunk.Value); err == nil {
			intVal = v
			isInt = true
		}

		if !generated.Exists(chunk.Value) && !isInt && chunk.Value != "" {
			generated.Set("", chunk.Value)
		}

		if isInt {
			// generated.Index(intVal)
			arrayValue := gabs.New()
			// generated.S(chunk.Parent.Value).SetIndex

			item := a.ParseFieldPath(f, chunk.Connector, arrayValue)
			// TOFIX - Only merge if part of the same object
			if generated.Index(intVal).Data() != nil && chunk.Field.Type == "object" {
				item.Merge(generated.Index(intVal))
			}

			generated.SetIndex(item.Data(), intVal)

		} else {
			// struggles with nested objects
			item := generated.S(chunk.Value)
			if item == nil {
				item = generated
			}
			if chunk.Connector != nil {
				a.ParseFieldPath(f, chunk.Connector, item)
			}
		}

		// if chunk.Connector != nil {
		// if chunk.Value != "" && !isInt {
		// generated.Set(a.ParseFieldPath(f, chunk.Connector, generated).Data(), chunk.Value)
		// } else {
		// return a.ParseFieldPath(f, chunk.Connector, generated)
		// }
		// }
		break
	default:
		generated.Set(f.Value, chunk.Value)
	}

	return generated
}

// Parse takes field registry and parses it into json string
func (a *Alpaca) Parse() string {
	result := gabs.New()

	if len(a.FieldRegistry) < 2 {
		return `"` + cast.ToString(a.FieldRegistry[0].Value) + `"`
	}

	for _, f := range a.FieldRegistry {

		if f.Value != nil && cast.ToString(f.Value) != "" {
			fmt.Println(f.Value)
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
