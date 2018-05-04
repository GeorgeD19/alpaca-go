package alpaca

import (
	"github.com/spf13/cast"
)

// Array container field
func (a *Alpaca) Array(f *Field) {
	f.IsContainerField = true

	maxItems := 1
	if f.Schema.Exists("maxItems") {
		maxItems = cast.ToInt(f.Schema.S("maxItems").Data().(float64))
	}

	if f.Schema.Exists("items") {
		for x := 0; x < maxItems; x++ {
			a.ResolveItemSchemaOptions(f.Key, f, x)
		}
	}

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

// Any control field
func (a *Alpaca) Any(f *Field) {
	a.RegisterField(f)
}
