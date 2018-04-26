package alpaca

import (
	"testing"

	"github.com/spf13/cast"
)

// TODO Address Field
// func TestAddressField(t *testing.T) {
// 	schema := `"schema": {
//         "title": "Home Address",
//         "type": "any"
//     },
//     "options": {
//         "type": "address",
//         "addressValidation": true,
//         "showMapOnLoad": true
// 	}`
// 	data := `{
// 		"street": ["308 Eddy Street", "Apartment #3"],
// 		"city": "Ithaca",
// 		"state": "NY",
// 		"zip": "14850"
// 	}`

// 	alpaca, _ := New(AlpacaOptions{Schema: schema, Data: data})
// 	if cast.ToString(alpaca.FieldRegistry[0].Value) != "test" {
// 		t.Fatalf("Should return test, instead returned %s", alpaca.FieldRegistry[0].Value)
// 	}
// }

func TestAnyField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "any",
			"title": "Any"
		}
	}`
	data := `"test"`

	alpaca, _ := New(AlpacaOptions{Schema: schema, Data: data})

	if cast.ToString(alpaca.FieldRegistry[0].Value) != "test" {
		t.Fatalf("Should return test, instead returned %s", alpaca.FieldRegistry[0].Value)
	}
}

func TestTextField(t *testing.T) {
	schema := `{"schema": {
        "minLength": 3,
        "maxLength": 8
	}}`
	data := `"Mint"`

	alpaca, _ := New(AlpacaOptions{Schema: schema, Data: data})

	if cast.ToString(alpaca.FieldRegistry[0].Value) != "Mint" {
		t.Fatalf("Should return test, instead returned %s", alpaca.FieldRegistry[0].Value)
	}
}

func TestTextField1(t *testing.T) {

	data := `"Mickey Mantle"`
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"label": "Name"
		}
	}`

	alpaca, _ := New(AlpacaOptions{Schema: schema, Data: data})

	if cast.ToString(alpaca.FieldRegistry[0].Value) != "Mickey Mantle" {
		t.Fatalf("Should return test, instead returned %s", alpaca.FieldRegistry[0].Value)
	}
}
