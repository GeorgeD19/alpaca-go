package alpaca

import (
	"bytes"
	"encoding/hex"
	"image"
	"io"
	"mime"
	"strconv"
	"time"

	"github.com/spf13/cast"
)

// Array container field
func (a *Alpaca) Array(f *Field) {
	f.IsContainerField = true

	maxItems := 10
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

	// If we have a request object we can attempt to grab image data
	if a.request != nil {

		maxImage := 10
		if f.Schema.Exists("maxImage") {
			maxImage = f.Schema.S("maxItems").Data().(int)
		}

		for x := 0; x < maxImage; x++ {
			FileName := f.Key + "_image_" + strconv.Itoa(x)

			file, _, err := a.request.FormFile(FileName)
			CreatedDevice := a.request.FormValue(FileName + "_created")

			if err == nil {
				defer file.Close()

				FoundFile := ImageFile{}
				var Buf bytes.Buffer
				io.Copy(&Buf, file)
				contents := Buf.Bytes()
				content := hex.EncodeToString(contents)
				FoundFile.Data = content

				file, _, _ := a.request.FormFile(FileName)
				config, format, _ := image.DecodeConfig(file)
				FoundFile.Name = FileName
				FoundFile.Width = config.Width
				FoundFile.Height = config.Height
				FoundFile.Type = format
				FoundFile.Mime = mime.TypeByExtension("." + format)
				FoundFile.FieldKey = f.Key
				FoundFile.FieldRef = f

				layout := "2006-01-02 15:04:05"
				t, err := time.Parse(layout, CreatedDevice)
				if err != nil {
					FoundFile.Created = time.Now()
				} else {
					FoundFile.Created = t
				}

				a.MediaRegistry = append(a.MediaRegistry, FoundFile)
				f.Media = append(f.Media, FoundFile)

				Buf.Reset()
			}
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
	// TODO Pull across code from existing alpaca package
	a.RegisterField(f)
}

// Any control field
func (a *Alpaca) Any(f *Field) {
	a.RegisterField(f)
}
