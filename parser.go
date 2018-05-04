package alpaca

import (
	"strconv"

	"github.com/Jeffail/gabs"
	"github.com/spf13/cast"
)

// ParseFieldPath reconstructs JSON based on the field path
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
