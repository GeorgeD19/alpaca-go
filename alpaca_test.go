package alpaca

import (
	"testing"

	"github.com/spf13/cast"
)

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

	data := `"Mint"`
	schema := `{"schema": {
        "minLength": 3,
        "maxLength": 8
    }}`

	alpaca, _ := New(AlpacaOptions{Schema: schema, Data: data})

	if cast.ToString(alpaca.FieldRegistry[0].Value) != "Mint" {
		t.Fatalf("Should return test, instead returned %s", alpaca.FieldRegistry[0].Value)
	}
}
