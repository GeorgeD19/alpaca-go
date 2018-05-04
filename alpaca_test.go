package alpaca

import (
	"testing"

	"github.com/spf13/cast"
)

// Core Fields
// Any Field http://www.alpacajs.org/docs/fields/any.html
func TestAnyField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "any",
			"title": "Test Any Field"
		}
	}`
	data := `"test"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestAnyField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"test"` {
		t.Fatalf(`Should return test, instead returned %s`, result)
	}
}

// Array Field http://www.alpacajs.org/docs/fields/array.html
func TestArrayField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "object",
			"properties": {
				"name": {
					"title": "Full Name",
					"type":"string"
				},
				"age": {
					"title": "Age",
					"type": "number"
				},
				"nest": {
					"type": "array",
					"maxItems": 3,
					"items": {
						"type": "object",
						"properties": {
							"name1": {
								"title": "Name 1",
								"type":"string"
							}
						}
					}
				}
			}
		},
		"options": {
			"fields": {
				"name": {
					"order": 1
				},
				"age": {
					"order": 2
				}
			}
		}
	}`
	data := `{
		"nest": [
			{
				"name1": "test3"
			},
			{
				"name1": "test1"
			},
			{
				"name1": "test2"
			}
		],
		"name": "test3",
		"age": 4
	}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestAnyField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"age":4,"name":"test3","nest":[{"name1":"test3"},{"name1":"test1"},{"name1":"test2"}]}` {
		t.Fatalf(`Should return test, instead returned %s`, result)
	}
}

// Checkbox Field http://www.alpacajs.org/docs/fields/checkbox.html
func TestCheckboxField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string",
			"enum": ["sandwich", "chips", "cookie", "drink"]
		},
		"options": {
			"type": "checkbox"
		}
	}`
	data := `"sandwich,cookie,drink"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestCheckboxField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"sandwich,cookie,drink"` {
		t.Fatalf(`Should return "sandwich,cookie,drink", instead returned %s`, result)
	}
}

// File Field http://www.alpacajs.org/docs/fields/file.html
// func TestFileField(t *testing.T) {

// }

// Hidden Field http://www.alpacajs.org/docs/fields/hidden.html
func TestHiddenField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "object",
			"properties": {
				"name": {
					"type": "string"
				},
				"password": {
					"type": "string"
				},
				"token": {
					"type": "string"
				}
			}
		},
		"options": {
			"fields": {
				"name": {
					"type": "text",
					"label": "Username"
				},
				"password": {
					"type": "password",
					"label": "Password"
				},
				"token": {
					"type": "hidden"
				}
			},
			"form": {
				"buttons": {
					"submit": {
						"value": "Log In"
					}
				}
			}
		}
	}`
	data := `{
        "token": "12345"
    }`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestCheckboxField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"token":"12345"}` {
		t.Fatalf(`Should return {"token":"12345"}, instead returned %s`, result)
	}
}

// Number Field http://www.alpacajs.org/docs/fields/number.html
func TestNumberField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "object",
			"properties": {
				"latitude": {
					"minimum": -180,
					"maximum": 180,
					"title": "Latitude",
					"required": true
				},
				"longitude": {
					"minimum": -180,
					"maximum": 180,
					"title": "Longitude",
					"required": true
				}
			}
		},
		"options": {
			"fields": {
				"latitude": {
					"type": "number"
				},
				"longitude": {
					"type": "number"
				}
			}
		}
	}`
	data := `{
        "latitude": 0,
        "longitude": 0
    }`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestNumberField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"latitude":0,"longitude":0}` {
		t.Fatalf(`Should return {"latitude":0,"longitude":0}, instead returned %s`, result)
	}
}

// Object http://www.alpacajs.org/docs/fields/object.html
func TestObjectField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "object",
			"properties": {
				"name": {
					"title": "Full Name",
					"type":"string"
				},
				"age": {
					"title": "Age",
					"type": "number"
				},
				"nest": {
					"type": "object",
					"properties": {
						"name1": {
							"title": "Name 1",
							"type":"string"
						}
					}
				}
			}
		},
		"options": {
			"fields": {
				"name": {
					"order": 1
				},
				"age": {
					"order": 2
				}
			}
		}
	}`
	data := `{
		"nest":{
			"name1": "Mint"
		},
		"name": "John Matrix",
		"age": 15
	}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestTextAreaField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"age":15,"name":"John Matrix","nest":{"name1":"Mint"}}` {
		t.Fatalf(`Should return {"age":15,"name":"John Matrix","nest":{"name1":"Mint"}}, instead returned %s`, result)
	}
}

// Radio http://www.alpacajs.org/docs/fields/radio.html
func TestRadioField(t *testing.T) {
	schema := `{
		"schema": {
			"enum": ["Jimi Hendrix", "Mark Knopfler", "Joe Satriani", "Eddie Van Halen", "Orianthi"]
		},
		"options": {
			"type": "radio",
			"label": "Who is your favorite guitarist?",
			"removeDefaultNone": true,
			"vertical": false
		}
	}`
	data := `"Jimi Hendrix"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestRadioField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"Jimi Hendrix"` {
		t.Fatalf(`Should return "Jimi Hendrix", instead returned %s`, result)
	}
}

// Select http://www.alpacajs.org/docs/fields/select.html - TODO Fix
func TestSelectField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "array",
			"items": {
				"title": "Ice Cream",
				"type": "string",
				"enum": ["Vanilla", "Chocolate", "Strawberry", "Mint"]
			},
			"minItems": 2,
			"maxItems": 3
		},
		"options": {
			"label": "Ice cream",
			"helper": "Guess my favorite ice cream?",
			"type": "select",
			"size": 5,
			"noneLabel": "Pick a flavour of Ice Cream!"
		}
	}`
	data := `["Vanilla", "Chocolate"]`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestSelectField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"Jimi Hendrix"` {
		t.Fatalf(`Should return "Jimi Hendrix", instead returned %s`, result)
	}
}

// Text Area http://www.alpacajs.org/docs/fields/textarea.html
func TestTextAreaField(t *testing.T) {
	schema := `{
		"schema": {},
		"options": {
			"type": "textarea"
		}
	}`
	data := `"Mint"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestTextAreaField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"Mint"` {
		t.Fatalf(`Should return "Mint", instead returned %s`, result)
	}
}

// Text http://www.alpacajs.org/docs/fields/text.html
func TestTextField(t *testing.T) {
	schema := `{
		"schema": {},
		"options": {
			"type": "text"
		}
	}`
	data := `"Mint"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestTextField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"Mint"` {
		t.Fatalf(`Should return "Mint", instead returned %s`, result)
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

// func TestAddressField(t *testing.T) {
// 	schema := `{
// 		"schema": {
// 			"
// 		}
// 	}`
// }
