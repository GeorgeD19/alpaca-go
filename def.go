package alpaca

import "time"

type Alpaca struct {
	Data   string
	Fields []*Field
	Media  map[string]File
}

// File type is a common base for files
type File struct {
	Data    string
	Type    string
	Mime    string
	Field   string
	Name    string
	Created time.Time
}

// ImageFile type extends File type to track width & height of image.
type ImageFile struct {
	File
	FieldKey string
	Width    int
	Height   int
}

type Field struct {
	Parent  *Field
	Schema  string
	Options string
	Data    string
	Name    string
	Key     string
	Path    string
	Type    string
	Value   interface{}
}
