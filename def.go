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
	"repeatable":  "object",
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

// ViewMessages contains messages used to flag up form issues
var ViewMessages = map[string]string{
	// date
	"invalidDate": "Invalid date for format {0}.",
	// editor
	"wordLimitExceeded":      "The maximum word limit of {0} has been exceeded.",
	"editorAnnotationsExist": "The editor has errors in it that must be corrected.",
	// email
	"invalidEmail": "Invalid Email address e.g. joe@blogs.com.",
	// integer
	"stringNotAnInteger": "This value is not an integer.",
	// IPv4
	"invalidIPv4": "Invalid IPv4 address, e.g. 192.168.0.1.",
	// json
	"stringNotAJSON": "This value is not a valid JSON string.",
	// map
	"keyNotUnique": "Keys of map field are not unique.",
	"keyMissing":   "Map contains an empty key.",
	// password
	"invalidPassword": "Invalid Password.",
	// phone
	"invalidPhone": "Invalid Phone Number, e.g. (123) 456-9999.",
	// time
	"invalidTime": "Invalid time.",
	// upload
	"chooseFile":                "Choose File...",
	"chooseFiles":               "Choose Files...",
	"dropZoneSingle":            "Click the Choose button or Drag and Drop a file here to upload...",
	"dropZoneMultiple":          "Click the Choose button or Drag and Drop files here to upload...",
	"dropZoneMultipleDirectory": "Click the Choose button or Drag and Drop files or a folder here to upload...",
	// url
	"invalidURLFormat": "The URL provided is not a valid web address.",
	// zipcode
	"invalidZipcodeFormatFive": "Invalid Five-Digit Zipcode (#####).",
	"invalidZipcodeFormatNine": "Invalid Nine-Digit Zipcode (#####-####).",
	// array
	"notEnoughItems": "The minimum number of items is {0}.",
	"tooManyItems":   "The maximum number of items is {0}.",
	"valueNotUnique": "Values are not unique",
	"notAnArray":     "This value is not an Array",
	// file
	"fileMissing": "This field should contain a file.",
	// number
	"stringValueTooSmall":          "The minimum value for this field is {0}.",
	"stringValueTooLarge":          "The maximum value for this field is {0}.",
	"stringValueTooSmallExclusive": "Value of this field must be greater than {0}.",
	"stringValueTooLargeExclusive": "Value of this field must be less than {0}.",
	"stringDivisibleBy":            "The value must be divisible by {0}.",
	"stringNotANumber":             "This value is not a number.",
	"stringValueNotMultipleOf":     "This value is not a multiple of {0}.",
	// object
	"tooManyProperties": "The maximum number of properties ({0}) has been exceeded.",
	"tooFewProperties":  "There are not enough properties ({0} are required).",
	// text
	"invalidPattern": "This field should have pattern {0}.",
	"stringTooShort": "This field should contain at least {0} numbers or characters.",
	"stringTooLong":  "This field should contain at most {0} numbers or characters.",
	// camera
	"imageMissing":   "This field should contain an image.",
	"imageTooMany":   "This field should contain at most {0} images.",
	"imageTooLittle": "This field should contain at least {0} images.",
	// signature
	"signatureMissing": "This field should contain a signature image",
	// list
	"noneLabel": "None",
}

type AlpacaOptions struct {
	Schema  string
	Data    string
	Request *http.Request
}

type Alpaca struct {
	data            *gabs.Container
	schema          *gabs.Container
	options         *gabs.Container
	connector       string
	request         *http.Request
	FieldRegistry   []*Field
	MediaRegistry   []ImageFile
	UniqueIDCounter int
}

type Chunk struct {
	Type      string
	Value     string
	Connector *Chunk
	Field     *Field
}

type Field struct {
	Data                *gabs.Container
	Options             *gabs.Container
	Schema              *gabs.Container
	Parent              *Field
	Children            []*Field
	ID                  string
	Key                 string
	Title               string
	Type                string
	Path                []Chunk
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
