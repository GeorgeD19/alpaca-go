package alpaca

import (
	"errors"
	"net/http"
	"time"

	"github.com/Jeffail/gabs"
)

// DefaultSchemaFieldMapping maps schema types to field types
var DefaultSchemaFieldMapping = map[string]string{
	"object":      "object",
	"repeatable":  "array",
	"string":      "text",
	"color":       "color",
	"search":      "search",
	"any":         "any",
	"array":       "array",
	"number":      "number",
	"information": "information",
	"boolean":     "checkbox",
}

// DefaultFormatFieldMapping maps format types across to field types
var DefaultFormatFieldMapping = map[string]string{
	"date":       "date",
	"datetime":   "datetime",
	"date-time":  "datetime",
	"email":      "email",
	"integer":    "integer",
	"ip-address": "ipv4",
	"lowercase":  "lowercase",
	"password":   "password",
	"phone":      "phone",
	"state":      "state",
	"time":       "time",
	"uppercase":  "uppercase",
	"url":        "url",
	"zipcode":    "zipcode",
}

// AlpacaOptions configures alpaca
type AlpacaOptions struct {
	Schema  string
	Data    string
	Request *http.Request
}

// Alpaca is the main operator of this package
type Alpaca struct {
	data            *gabs.Container
	schema          *gabs.Container
	options         *gabs.Container
	connector       string
	request         *http.Request
	FieldRegistry   []*Field
	MediaRegistry   []ImageFile
	UniqueIDCounter int
	output          string
}

// Chunk is used to construct a field path
type Chunk struct {
	Type      string
	Value     string
	Size      int
	Connector *Chunk
	Parent    *Chunk
	Field     *Field
}

// Field is a field of any kind
type Field struct {
	Data                *gabs.Container
	Options             *gabs.Container
	Schema              *gabs.Container
	Parent              *Field
	Children            []*Field
	ID                  string
	Key                 string
	Title               string
	SchemaType          string
	ChunkType           string
	OptionsType         string
	Type                string
	Path                []Chunk
	PathString          string
	Validate            string
	ShowingDefaultData  string
	PreviouslyValidated bool
	IsContainerField    bool
	Value               interface{}
	ValueType           string
	Default             interface{}
	DefaultType         string
	Order               int
	ReadOnly            bool
	notTopLevel         bool
	IsArrayChild        bool
	ArrayIndex          int
	ArrayValues         int
	Media               []ImageFile
}

// StandardFile type is a common base for files.
type StandardFile struct {
	Data     string
	Type     string
	Mime     string
	Field    string
	Name     string
	Created  time.Time
	FieldRef *Field
}

// ImageFile type extends File type to track width & height of image.
type ImageFile struct {
	StandardFile
	FieldKey string
	Width    int
	Height   int
}

var (
	ErrDefaultError  = errors.New("You must supply at least one argument.")
	ErrSchemaInvalid = errors.New("Invalid schema supplied.")
	ErrDataInvalid   = errors.New("Invalid data supplied.")
)
