package alpaca

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// Array container field
func (a *Alpaca) Array(f *Field) {
	f.IsContainerField = true

	maxItems := 1
	if f.Schema.Exists("maxItems") {
		maxItems = cast.ToInt(f.Schema.S("maxItems").Data().(float64))
	}

	isInt := false
	intVal := 0
	if v, err := strconv.Atoi(f.Key); err == nil {
		intVal = v
		isInt = true
	}

	if f.Schema.Exists("items") {
		for x := 0; x < maxItems; x++ {
			a.ResolveItemSchemaOptions(f.Key, f, x)
		}
	} else if isInt {
		if f.SchemaType == "" {
			a.ResolveItemSchemaOptions(f.Key, f, intVal)
		}
	}

	if f.Type == "select" {
		if f.Schema.Exists("enum") {
			enum, err := f.Schema.S("enum").Children()
			if err == nil {
				// Enum fallback value
				for i := 0; i < len(enum); i++ {
					if cast.ToString(f.Value) == cast.ToString(enum[i].Data()) {
						f.EnumLabel = cast.ToString(enum[i].Data())
					}
				}
				if f.Options.Exists("optionLabels") {
					optionLabels, err := f.Options.S("optionLabels").Children()
					if err == nil {
						for i := 0; i < len(enum); i++ {
							if cast.ToString(f.Value) == cast.ToString(enum[i].Data()) {
								f.EnumLabel = cast.ToString(optionLabels[i].Data())
							}
							f.Enum = append(f.Enum, Enum{
								Value: enum[i].Data(),
								Label: optionLabels[i].Data(),
							})
						}
					}
				}
			}
		}
	}

	a.RegisterField(f)
}

// Tag control field
func (a *Alpaca) Tag(f *Field) {
	f.Value = strings.TrimSuffix(strings.TrimPrefix(f.Data.String(), `"`), `"`)
	a.RegisterField(f)
}

// Object container field
func (a *Alpaca) Object(f *Field) {
	f.IsContainerField = true

	properties, err := f.Schema.S("properties").ChildrenMap()
	if err == nil {
		for key := range properties {
			a.ResolvePropertySchemaOptions(key, f)
		}
	}
	a.RegisterField(f)
}

// Camera container field
func (a *Alpaca) Camera(f *Field) {
	if a.request != nil {

		maxImage := 10
		if f.Schema.Exists("maxImage") {
			maxImage = f.Schema.S("maxItems").Data().(int)
		}

		for x := 0; x < maxImage; x++ {
			a.RegisterMedia(f, x)
		}
	}

	a.RegisterField(f)
}

// Information container field
func (a *Alpaca) Information(f *Field) {
	f.IsContainerField = true
	a.RegisterField(f)
}

// Signature container field
func (a *Alpaca) Signature(f *Field) {
	if a.request != nil {
		for x := 0; x < 1; x++ {
			a.RegisterMedia(f, x)
		}
	}
	a.RegisterField(f)
}

// Datetime control field
func (a *Alpaca) Datetime(f *Field) {
	str := cast.ToString(f.Data.Data())
	layout := "2006-01-02T15:04:05.000"
	t, err := time.Parse(layout, str)
	if err != nil {
		f.Value = f.Data.Data()
	} else {
		f.Value = t.Format("2006-01-02 03:04:05")
	}
	a.RegisterField(f)
}

// Any control field
func (a *Alpaca) Any(f *Field) {
	a.RegisterField(f)
}

// Editor control field
func (a *Alpaca) Editor(f *Field) {
	JSON := []byte(f.Data.String())
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, JSON); err != nil {
		f.Value = f.Data.Data()
	}
	f.Value = strings.TrimSuffix(strings.TrimPrefix(cast.ToString(buffer), `"`), `"`)
	a.RegisterField(f)
}

// JSON control field
func (a *Alpaca) JSON(f *Field) {
	a.Editor(f)
}

// Lowercase control field
func (a *Alpaca) Lowercase(f *Field) {
	f.Value = strings.TrimSuffix(strings.TrimPrefix(strings.ToLower(f.Data.String()), `"`), `"`)
	a.RegisterField(f)
}

// Uppercase control field
func (a *Alpaca) Uppercase(f *Field) {
	f.Value = strings.TrimSuffix(strings.TrimPrefix(strings.ToUpper(f.Data.String()), `"`), `"`)
	a.RegisterField(f)
}
