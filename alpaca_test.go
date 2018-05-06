package alpaca

import (
	"testing"
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
		t.Fatalf("TestObjectField error: %s", err)
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
	if result != `["Vanilla", "Chocolate"]` {
		t.Fatalf(`Should return ["Vanilla", "Chocolate"], instead returned %s`, result)
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
	schema := `{
		"schema": {
			"type": "string"
		}
	}`
	data := `"Mickey Mantle"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestTextField1 error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"Mickey Mantle"` {
		t.Fatalf(`Should return "Mickey Mantle", instead returned %s`, alpaca.FieldRegistry[0].Value)
	}
}

// More Fields
// http://www.alpacajs.org/docs/fields/address.html
func TestAddressField(t *testing.T) {
	schema := `{
		"schema": {
			"title": "Home Address",
			"type": "any"
		},
		"options": {
			"type": "address"
		}
	}`
	data := `{"street":["street 1","street 2","street 3"],"city":"glasgow","state":"AL","zip":"23233"}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestAddressField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"street":["street 1","street 2","street 3"],"city":"glasgow","state":"AL","zip":"23233"}` {
		t.Fatalf(`Should return {"street":["street 1","street 2","street 3"],"city":"glasgow","state":"AL","zip":"23233"}, instead returned %s`, result)
	}
}

func TestChooserField(t *testing.T) {

}

func TestColorField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "color"
		}
	}`
	data := `"#ff8000"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestColorField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"#ff8000"` {
		t.Fatalf(`Should return "#ff8000", instead returned %s`, result)
	}
}

func TestColorPickerField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "colorpicker"
		}
	}`
	data := `"#ff8000"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestColorPickerField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"#ff8000"` {
		t.Fatalf(`Should return "#ff8000", instead returned %s`, result)
	}
}

func TestCKEditorField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "ckeditor"
		}
	}`
	data := `"Ice cream is a <b>frozen</b> dessert usually made from <i>dairy products</i>, such as milk and cream, and often combined with fruits or other ingredients and flavors."`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestCKEditorField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"Ice cream is a <b>frozen</b> dessert usually made from <i>dairy products</i>, such as milk and cream, and often combined with fruits or other ingredients and flavors."` {
		t.Fatalf(`Should return "Ice cream is a <b>frozen</b> dessert usually made from <i>dairy products</i>, such as milk and cream, and often combined with fruits or other ingredients and flavors.", instead returned %s`, result)
	}
}

func TestCountryField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "country"
		}
	}`
	data := `"gbr"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestCountryField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"gbr"` {
		t.Fatalf(`Should return "gbr", instead returned %s`, result)
	}
}

func TestCurrencyField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "currency"
		}
	}`
	data := `"413.21"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestCurrencyField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"413.21"` {
		t.Fatalf(`Should return "413.21", instead returned %s`, result)
	}
}

func TestDateField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "date"
		}
	}`
	data := `"05/03/2018"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestDateField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"05/03/2018"` {
		t.Fatalf(`Should return "05/03/2018", instead returned %s`, result)
	}
}

func TestDateTimeField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "datetime"
		}
	}`
	data := `"05/03/2018 00:00:06"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestDateTimeField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"05/03/2018 00:00:06"` {
		t.Fatalf(`Should return "05/03/2018 00:00:06", instead returned %s`, result)
	}
}

func TestEditorField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "editor"
		}
	}`
	data := `"{\n\t\"test\":\"test\"\n}"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestDateTimeField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"{\n\t\"test\":\"test\"\n}"` {
		t.Fatalf(`Should return "{\n\t\"test\":\"test\"\n}", instead returned %s`, result)
	}
}

func TestEmailField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "email"
		}
	}`
	data := `"test@test.com"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestEmailField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"test@test.com"` {
		t.Fatalf(`Should return "test@test.com", instead returned %s`, result)
	}
}

// Grid not supported - can't get any data to submit from client
// func TestGridField(t *testing.T) {

// }

// Image field doesn't accept data, only renders it
func TestImageField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "image"
		}
	}`
	data := `"image"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestImageField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `""` {
		t.Fatalf(`Should return "", instead returned %s`, result)
	}
}

func TestIntegerField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "integer"
		}
	}`
	data := `17`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestEmailField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `17` {
		t.Fatalf(`Should return 17, instead returned %s`, result)
	}
}

func TestObjectIntegerField(t *testing.T) {
	schema := `{
		"schema": {
			"type":"object",
			"properties": {
				"integerfield": {
					"type": "string"
				}
			}
		},
		"options": {
			"fields": {
				"integerfield": {
					"type": "integer"
				}
			}
		}
	}`
	data := `{"integerfield":17}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestObjectIntegerField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"integerfield":17}` {
		t.Fatalf(`Should return {"integerfield":17}, instead returned %s`, result)
	}
}

func TestIPV4Field(t *testing.T) {
	schema := `{
		"schema": {
			"format": "ip-address"
		}
	}`
	data := `"128.253.180.2"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestIPV4Field error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"128.253.180.2"` {
		t.Fatalf(`Should return "128.253.180.2", instead returned %s`, result)
	}
}

func TestJSONField(t *testing.T) {
	schema := `{
		"options": {
			"type": "json"
		}
	}`
	data := `{"test":"test2"}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestJSONField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"test":"test2"}` {
		t.Fatalf(`Should return {"test":"test2"}, instead returned %s`, result)
	}
}

func TestLowerCaseField(t *testing.T) {
	schema := `{
		"schema": {
			"format": "lowercase"
		}
	}`
	data := `"Ice cream is wonderful."`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestLowerCaseField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"ice cream is wonderful."` {
		t.Fatalf(`Should return "ice cream is wonderful.", instead returned %s`, result)
	}
}

func TestMapField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "array",
			"items": {
				"type": "object",
				"properties": {
					"_key": {
						"title": "User ID",
						"type": "string"
					},
					"firstName": {
						"title": "First Name",
						"type": "string"
					},
					"lastName": {
						"title": "Last Name",
						"type": "string"
					},
					"gender": {
						"title": "Gender",
						"type": "string",
						"enum": ["Male", "Female"]
					}
				}
			}
		},
		"options": {
			"type": "map",
			"toolbarSticky": true,
			"items": {
				"fields": {
					"_key": {
						"size": 60,
						"helper": "This value serves as a unique key into the associative array."
					}
				}
			}
		}
	}`
	data := `{
        "john316": {
            "firstName": "Tim",
            "lastName": "Tebow",
            "gender": "Male"
        },
        "ladygaga": {
            "firstName": "Stefani",
            "lastName": "Germanotta",
            "gender": "Female"
        }
    }`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `19` {
		t.Fatalf(`Should return 19, instead returned %s`, result)
	}
}

// Option Tree Field http://www.alpacajs.org/docs/fields/optiontree.html
func TestOptionTreeField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "number",
			"title": "What number would like for your sports jersey?"
		},
		"options": {
			"type": "optiontree",
			"tree": {
				"selectors": {
					"sport": {
						"schema": {
							"type": "string"
						},
						"options": {
							"type": "select",
							"noneLabel": "Pick a Sport..."
						}
					},
					"team": {
						"schema": {
							"type": "string"
						},
						"options": {
							"type": "select",
							"noneLabel": "Pick a Team..."
						}
					},
					"player": {
						"schema": {
							"type": "string"
						},
						"options": {
							"type": "select",
							"noneLabel": "Pick a Player..."
						}
					}
				},
				"order": ["sport", "team", "player"],
				"data": [{
					"value": 23,
					"attributes": {
						"sport": "Basketball",
						"team": "Chicago Bulls",
						"player": "Michael Jordan"
					}
				}, {
					"value": 33,
					"attributes": {
						"sport": "Basketball",
						"team": "Chicago Bulls",
						"player": "Scotty Pippen"
					}
				}, {
					"value": 4,
					"attributes": {
						"sport": "Football",
						"team": "Green Bay Packers",
						"player": "Brett Favre"
					}
				}, {
					"value": 19,
					"attributes": {
						"sport": "Baseball",
						"team": "Milwaukee Brewers",
						"player": "Robin Yount"
					}
				}, {
					"value": 99,
					"attributes": {
						"sport": "Hockey",
						"player": "Wayne Gretzky"
					}
				}],
				"horizontal": true
			}
		}
	}`
	data := `19`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `19` {
		t.Fatalf(`Should return 19, instead returned %s`, result)
	}
}

// Password Field http://www.alpacajs.org/docs/fields/password.html
func TestPasswordField(t *testing.T) {
	schema := `{
		"schema": {
			"format": "password"
		}
	}`
	data := `"password"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"password"` {
		t.Fatalf(`Should return "password", instead returned %s`, result)
	}
}

func TestPersonalNameField(t *testing.T) {
	schema := `{
		"options": {
			"type": "personalname"
		}
	}`
	data := `"Oscar Zoroaster Phadrig Isaac Norman Henkel Emmannuel Ambroise Diggs"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"Oscar Zoroaster Phadrig Isaac Norman Henkel Emmannuel Ambroise Diggs"` {
		t.Fatalf(`Should return "Oscar Zoroaster Phadrig Isaac Norman Henkel Emmannuel Ambroise Diggs", instead returned %s`, result)
	}
}

// Phone Field http://www.alpacajs.org/docs/fields/phone.html
func TestPhoneField(t *testing.T) {
	schema := `{
		"schema": {
			"format": "phone"
		}
	}`
	data := `"2145324635"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"2145324635"` {
		t.Fatalf(`Should return "2145324635", instead returned %s`, result)
	}
}

func TestPickAColorField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "pickacolor"
		}
	}`
	data := `"#bb9977"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"#bb9977"` {
		t.Fatalf(`Should return "#bb9977", instead returned %s`, result)
	}
}

func TestSearchField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "search"
		}
	}`
	data := `"Where for art thou Romeo?"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"Where for art thou Romeo?"` {
		t.Fatalf(`Should return "Where for art thou Romeo?", instead returned %s`, result)
	}
}

func TestStateField(t *testing.T) {
	schema := `{
		"options": {
			"type": "state"
		}
	}`
	data := `"AR"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"AR"` {
		t.Fatalf(`Should return "AR", instead returned %s`, result)
	}
}

func TestSummernoteEditorField(t *testing.T) {

}

func TestTableField(t *testing.T) {

}

// Tag fields convert comma seperated string into an array of values, all done on client
func TestTagField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "tag"
		}
	}`
	data := `["great", "wonderful", "ice cream"]`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `["great", "wonderful", "ice cream"]` {
		t.Fatalf(`Should return ["great", "wonderful", "ice cream"], instead returned %s`, result)
	}
}

// TODO Write a test for the Tag field which converts a comma delimited string into an array

func TestTokenField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "string"
		},
		"options": {
			"type": "token",
			"tokenfield": {
				"autocomplete": {
					"source": ["marty", "doc", "george", "biff", "lorraine", "mr. strickland"],
					"delay": 100
				},
				"showAutocompleteOnFocus": true
			}
		}
	}`
	data := `"marty,doc,george,biff"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"marty,doc,george,biff"` {
		t.Fatalf(`Should return "marty,doc,george,biff", instead returned %s`, result)
	}
}

func TestTimeField(t *testing.T) {
	schema := `{
		"schema": {
			"format": "time"
		}
	}`
	data := `"00:00:35"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"00:00:35"` {
		t.Fatalf(`Should return "00:00:35", instead returned %s`, result)
	}
}

func TestTinyMCEField(t *testing.T) {

}

func TestUploadField(t *testing.T) {

}

func TestUpperCaseField(t *testing.T) {
	schema := `{
		"schema": {
			"format": "uppercase"
		}
	}`
	data := `"Ice cream is wonderful!"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"ICE CREAM IS WONDERFUL!"` {
		t.Fatalf(`Should return "ICE CREAM IS WONDERFUL!", instead returned %s`, result)
	}
}

func TestURLField(t *testing.T) {
	schema := `{
		"options": {
			"type": "url"
		},
		"schema": {
			"format": "uri"
		}
	}`
	data := `"http://www.alpacajs.org"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"http://www.alpacajs.org"` {
		t.Fatalf(`Should return "http://www.alpacajs.org", instead returned %s`, result)
	}
}

// Zip Code Field http://www.alpacajs.org/docs/fields/zipcode.html
func TestZipCodeField(t *testing.T) {
	schema := `{
		"options": {
			"type": "zipcode",
			"format": "five"
		}
	}`
	data := `"53221"`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `"53221"` {
		t.Fatalf(`Should return "53221", instead returned %s`, result)
	}
}

// Forms
func TestORSORDERForm(t *testing.T) {
	schema := `{
		"options": {
			"fields": {
				"actions_taken": {
					"order": 3
				},
				"client_managers_name": {
					"order": 1
				},
				"details_of_problem": {
					"order": 2
				},
				"fix_term": {
					"order": 4,
					"type": "select"
				},
				"manager_director_name": {
					"order": 5
				},
				"reference": {
					"order": 7
				},
				"signature": {
					"order": 6,
					"type": "signature"
				}
			}
		},
		"schema": {
			"properties": {
				"actions_taken": {
					"required": true,
					"title": "What actions have been taken?",
					"type": "string"
				},
				"client_managers_name": {
					"required": true,
					"title": "Client Manager's name",
					"type": "string"
				},
				"details_of_problem": {
					"required": true,
					"title": "Details of problem",
					"type": "string"
				},
				"fix_term": {
					"enum": ["Yes", "No"],
					"required": true,
					"title": "Is the problem fixed for the long term?",
					"type": "string"
				},
				"manager_director_name": {
					"required": true,
					"title": "SecuriGroup Manager / Director name",
					"type": "string"
				},
				"reference": {
					"default": "SL Feb 17 Ref:C069",
					"readonly": true,
					"title": "Form reference",
					"type": "string"
				},
				"signature": {
					"required": true,
					"title": "Signature",
					"type": "string"
				}
			},
			"type": "object"
		}
	}`

	data := `{
		"actions_taken": "None",
		"client_managers_name": "Winnie The Poo",
		"details_of_problem": "Humpty Dumpty fell off the wall.",
		"fix_term": "Yes",
		"manager_director_name": "Tigger",
		"signature": "[Signature]"
	}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"actions_taken":"None","client_managers_name":"Winnie The Poo","details_of_problem":"Humpty Dumpty fell off the wall.","fix_term":"Yes","manager_director_name":"Tigger","reference":"SL Feb 17 Ref:C069","signature":"[Signature]"}` {
		t.Fatalf(`Should return {"actions_taken":"None","client_managers_name":"Winnie The Poo","details_of_problem":"Humpty Dumpty fell off the wall.","fix_term":"Yes","manager_director_name":"Tigger","reference":"SL Feb 17 Ref:C069","signature":"[Signature]"}, instead returned %s`, result)
	}
}

func TestOEVLate1Form(t *testing.T) {
	schema := `{
		"schema": {
			"properties": {
				"oev_incomplete_information": {
					"description": "This OEV was overdue from last month, you know how important that these are completed on time!",
					"type": "information"
				},
				"oev_incomplete_reason": {
					"required": true,
					"title": "Why was the OEV not completed within the allotted month?",
					"type": "string"
				},
				"oev_completion_date": {
					"title": "When in the next 7 days are you planning to do this OEV?",
					"required": true,
					"type": "string"
				},
				"reference": {
					"default": "SL Feb 17 Ref:C060",
					"readonly": true,
					"title": "Form reference",
					"type": "string"
				}
			},
			"type": "object"
		},
		"options": {
			"fields": {
				"oev_incomplete_information": {
					"order": 1
				},
				"oev_incomplete_reason": {
					"order": 2
				},
				"oev_completion_date": {
					"order": 3,
					"type": "date"
				},
				"reference": {
					"order": 4
				}
			}
		}
	}`

	data := `{
		"oev_incomplete_reason": "My dog ate it.",
		"oev_completion_date": "15/10/01"
	}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"oev_completion_date":"15/10/01","oev_incomplete_reason":"My dog ate it.","reference":"SL Feb 17 Ref:C060"}` {
		t.Fatalf(`Should return {"oev_completion_date":"15/10/01","oev_incomplete_reason":"My dog ate it.","reference":"SL Feb 17 Ref:C060"}, instead returned %s`, result)
	}
}

func TestNOTESForm(t *testing.T) {
	schema := `{
		"options": {
			"fields": {
				"notes": {
					"order": 1,
					"type": "textarea"
				},
				"reference": {
					"order": 2
				}
			}
		},
		"schema": {
			"properties": {
				"notes": {
					"required": true,
					"title": "Please update some details.",
					"type": "string"
				},
				"reference": {
					"default": "SL Feb 17 Ref:C056",
					"readonly": true,
					"title": "Form reference",
					"type": "string"
				}
			},
			"type": "object"
		}
	}`

	data := `{
		"notes": "Some notes."
	}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"notes":"Some notes.","reference":"SL Feb 17 Ref:C056"}` {
		t.Fatalf(`Should return {"notes":"Some notes.","reference":"SL Feb 17 Ref:C056"}, instead returned %s`, result)
	}
}

func TestPearsonHazardReportForm(t *testing.T) {
	schema := `{
		"schema": {
			"type": "object",
			"properties": {
				"date": {
					"type": "string",
					"title": "Date",
					"required": true
				},
				"time": {
					"type": "string",
					"title": "Time",
					"required": true
				},
				"site": {
					"type": "string",
					"title": "Site",
					"enum": ["Holborn", "Fretwell Road", "Lowton House", "Jordan Hill"],
					"required": true
				},
				"hazard_type": {
					"type": "string",
					"title": "Hazard Type",
					"required": true
				},
				"location": {
					"type": "string",
					"title": "Location",
					"required": true
				},
				"risk_profile": {
					"type": "string",
					"title": "Risk Profile + Resolution Time",
					"enum": ["Low - 6 Hours", "Medium - 1 Hour", "High - Immediately"],
					"required": true
				},
				"picture": {
					"type": "string",
					"title": "Picture",
					"required": true
				},
				"comments": {
					"type": "string",
					"title": "comments"
				},
				"form_reference": {
					"type": "string",
					"title": "Form Reference",
					"default": "PS HR Dec 17 Ref:001",
					"readonly": true
				}
			}
		},
		"options": {
			"fields": {
				"date": {
					"type": "date"
				},
				"time": {
					"type": "time"
				},
				"site": {
					"removeDefaultNone": false,
					"sort": false
				},
				"picture": {
					"type": "camera"
				},
				"comments": {
					"type": "textarea"
				}
			}
		}
	}`

	data := `{
		"date": "12/04/18",
		"time": "00:12:40",
		"site": "Holborn",
		"hazard_type": "Some hazard type",
		"location": "Some location",
		"risk_profile": "Low - 6 Hours",
		"picture": "[Image]",
		"comments": "Bleh"
	}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"comments":"Bleh","date":"12/04/18","form_reference":"PS HR Dec 17 Ref:001","hazard_type":"Some hazard type","location":"Some location","picture":"[Image]","risk_profile":"Low - 6 Hours","site":"Holborn","time":"00:12:40"}` {
		t.Fatalf(`Should return {"comments":"Bleh","date":"12/04/18","form_reference":"PS HR Dec 17 Ref:001","hazard_type":"Some hazard type","location":"Some location","picture":"[Image]","risk_profile":"Low - 6 Hours","site":"Holborn","time":"00:12:40"}, instead returned %s`, result)
	}
}

func TestAshdownPhillipsSiteReportForm(t *testing.T) {
	schema := `{
		"schema": {
			"type": "object",
			"properties": {
				"site": {
					"type": "string",
					"title": "Select Site",
					"required": true,
					"enum": ["Charter House", "Dover Street", "Shepherds Building"]
				},
				"form_type": {
					"title": "Form Type",
					"type": "string",
					"enum": ["Incident Report", "Locker Clearance", "Another Type", "Extra Type"],
					"required": true
				},
				"incident_report_completed_by": {
					"type": "string",
					"title": "Report Completed By",
					"required": true
				},
				"incident_report_position": {
					"type": "string",
					"title": "Position",
					"required": true
				},
				"incident_report_datetime": {
					"type": "string",
					"title": "Date and Time of Incident",
					"required": true
				},
				"incident_report_incident_type": {
					"type": "string",
					"enum": ["Fire", "Flood", "Intruder", "Near Miss", "Suspicious Behaviour", "Theft", "Violence", "Other"],
					"title": "Incident Type",
					"required": true
				},
				"incident_report_incident_other": {
					"type": "string",
					"title": "Please specify"
				},
				"incident_report_emergency_services_called": {
					"type": "string",
					"title": "Emergency Services Called?",
					"enum": ["No", "Yes"]
				},
				"incident_report_emergency_services_ambulance_service": {
					"title": "Ambulance Service",
					"type": "string",
					"enum": ["Yes", "No"]
				},
				"incident_report_emergency_services_ambulance_called_at": {
					"type": "string",
					"title": "Called at",
					"required": true
				},
				"incident_report_emergency_services_ambulance_arrived_at": {
					"type": "string",
					"title": "Arrived at",
					"required": true
				},
				"incident_report_emergency_services_ambulance_left_at": {
					"type": "string",
					"title": "Left at",
					"required": true
				},
				"incident_report_emergency_services_ambulance_callsign": {
					"type": "string",
					"title": "Callsign",
					"required": true
				},
				"incident_report_emergency_services_fire_service": {
					"title": "Fire Service",
					"type": "string",
					"enum": ["Yes", "No"]
				},
				"incident_report_emergency_services_fire_service_called_at": {
					"type": "string",
					"title": "Called at",
					"required": true
				},
				"incident_report_emergency_services_fire_service_arrived_at": {
					"type": "string",
					"title": "Arrived at",
					"required": true
				},
				"incident_report_emergency_services_fire_service_left_at": {
					"type": "string",
					"title": "Left at",
					"required": true
				},
				"incident_report_emergency_services_fire_service_callsign": {
					"type": "string",
					"title": "Callsign",
					"required": true
				},
				"incident_report_emergency_services_police_service": {
					"title": "Police",
					"type": "string",
					"enum": ["Yes", "No"]
				},
				"incident_report_emergency_services_police_called_at": {
					"type": "string",
					"title": "Called at",
					"required": true
				},
				"incident_report_emergency_services_police_arrived_at": {
					"type": "string",
					"title": "Arrived at",
					"required": true
				},
				"incident_report_emergency_services_police_left_at": {
					"type": "string",
					"title": "Left at",
					"required": true
				},
				"incident_report_emergency_services_police_callsign": {
					"type": "string",
					"title": "Callsign",
					"required": true
				},
				"incident_report_incident_details": {
					"type": "string",
					"title": "Incident Details"
				},
				"incident_report_who": {
					"type": "string",
					"title": "Who?",
					"description": "(Who was involved/witnessed the incident. Give description if no names)"
				},
				"incident_report_where": {
					"type": "string",
					"title": "Where?",
					"description": "(Exactly where did the Incident happen, be as precise as possible)"
				},
				"incident_report_what": {
					"type": "string",
					"title": "What?",
					"description": "(What happened? Start from the beginning and work through events as they occurred)"
				},
				"incident_report_images": {
					"type": "array",
					"maxItems": 10,
					"items": {
						"type": "object",
						"title": "Image",
						"properties": {
							"name": {
								"type": "string",
								"title": "Name",
								"required": true
							},
							"image": {
								"type": "string",
								"title": "camera",
								"required": true
							}
						}
					}
				},
				"incident_report_follow_up": {
					"type": "string",
					"title": "Follow up since the incident?",
					"description": "(Have the Police requested CCTV? Has management been contacted already?)"
				},
				"incident_report_single_image": {
					"type": "string",
					"title": "Report Image"
				},
				"locker_clearance_locker_number": {
					"type": "string",
					"title": "Locker Number",
					"required": true
				},
				"locker_clearance_cleaner_signature_name": {
					"type": "string",
					"title": "Cleaner Signature & Name",
					"required": true
				},
				"form_reference": {
					"type": "string",
					"title": "Form Reference",
					"default": "SL Feb 17 Ref:C003",
					"readonly": true
				}
			},
			"dependencies": {
				"incident_report_completed_by": ["form_type"],
				"incident_report_position": ["form_type"],
				"incident_report_datetime": ["form_type"],
				"incident_report_incident_type": ["form_type"],
				"incident_report_incident_other": ["form_type", "incident_report_incident_type"],
				"incident_report_emergency_services_called": ["form_type"],
				"incident_report_emergency_services_ambulance_service": ["form_type", "incident_report_emergency_services_called"],
				"incident_report_emergency_services_ambulance_called_at": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_ambulance_service"],
				"incident_report_emergency_services_ambulance_arrived_at": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_ambulance_service"],
				"incident_report_emergency_services_ambulance_left_at": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_ambulance_service"],
				"incident_report_emergency_services_ambulance_callsign": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_ambulance_service"],
				"incident_report_emergency_services_fire_service": ["form_type", "incident_report_emergency_services_called"],
				"incident_report_emergency_services_fire_service_called_at": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_fire_service"],
				"incident_report_emergency_services_fire_service_arrived_at": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_fire_service"],
				"incident_report_emergency_services_fire_service_left_at": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_fire_service"],
				"incident_report_emergency_services_fire_service_callsign": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_fire_service"],
				"incident_report_emergency_services_police_service": ["form_type", "incident_report_emergency_services_called"],
				"incident_report_emergency_services_police_called_at": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_police_service"],
				"incident_report_emergency_services_police_arrived_at": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_police_service"],
				"incident_report_emergency_services_police_left_at": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_police_service"],
				"incident_report_emergency_services_police_callsign": ["form_type", "incident_report_emergency_services_called", "incident_report_emergency_services_police_service"],
				"incident_report_incident_details": ["form_type"],
				"incident_report_who": ["form_type"],
				"incident_report_where": ["form_type"],
				"incident_report_what": ["form_type"],
				"incident_report_single_image": ["form_type"],
				"incident_report_follow_up": ["form_type"],
				"locker_clearance_locker_number": ["form_type"],
				"locker_clearance_cleaner_signature_name": ["form_type"]
			}
		},
		"options": {
			"fields": {
				"site": {
					"removeDefaultNone": false
				},
				"form_type": {
					"removeDefaultNone": false
				},
				"incident_report_completed_by": {
					"dependencies": {
						"form_type": ["Incident Report"]
					}
				},
				"incident_report_position": {
					"dependencies": {
						"form_type": "Incident Report"
					}
				},
				"incident_report_datetime": {
					"dependencies": {
						"form_type": "Incident Report"
					}
				},
				"incident_report_incident_type": {
					"removeDefaultNone": false,
					"dependencies": {
						"form_type": "Incident Report"
					}
				},
				"incident_report_incident_other": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_incident_type": "Other"
					}
				},
				"incident_report_emergency_services_called": {
					"removeDefaultNone": true,
					"dependencies": {
						"form_type": "Incident Report"
					}
				},
				"incident_report_emergency_services_ambulance_service": {
					"removeDefaultNone": true,
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes"
					}
				},
				"incident_report_emergency_services_ambulance_called_at": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_ambulance_service": "Yes"
					}
				},
				"incident_report_emergency_services_ambulance_arrived_at": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_ambulance_service": "Yes"
					}
				},
				"incident_report_emergency_services_ambulance_left_at": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_ambulance_service": "Yes"
					}
				},
				"incident_report_emergency_services_ambulance_callsign": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_ambulance_service": "Yes"
					}
				},
				"incident_report_emergency_services_fire_service": {
					"removeDefaultNone": true,
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes"
					}
				},
				"incident_report_emergency_services_fire_service_called_at": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_fire_service": "Yes"
					}
				},
				"incident_report_emergency_services_fire_service_arrived_at": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_fire_service": "Yes"
					}
				},
				"incident_report_emergency_services_fire_service_left_at": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_fire_service": "Yes"
					}
				},
				"incident_report_emergency_services_fire_service_callsign": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_fire_service": "Yes"
					}
				},
				"incident_report_emergency_services_police_service": {
					"removeDefaultNone": true,
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes"
					}
				},
				"incident_report_emergency_services_police_called_at": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_police_service": "Yes"
					}
				},
				"incident_report_emergency_services_police_arrived_at": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_police_service": "Yes"
					}
				},
				"incident_report_emergency_services_police_left_at": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_police_service": "Yes"
					}
				},
				"incident_report_emergency_services_police_callsign": {
					"dependencies": {
						"form_type": "Incident Report",
						"incident_report_emergency_services_called": "Yes",
						"incident_report_emergency_services_police_service": "Yes"
					}
				},
				"incident_report_incident_details": {
					"dependencies": {
						"form_type": "Incident Report"
					}
				},
				"incident_report_who": {
					"dependencies": {
						"form_type": "Incident Report"
					}
				},
				"incident_report_where": {
					"dependencies": {
						"form_type": "Incident Report"
					}
				},
				"incident_report_what": {
					"dependencies": {
						"form_type": "Incident Report"
					}
				},
				"incident_report_single_image": {
					"type": "camera",
					"dependencies": {
						"form_type": "Incident Report"
					}
				},
				"incident_report_follow_up": {
					"dependencies": {
						"form_type": "Incident Report"
					}
				},
				"incident_report_images": {
					"type": "repeatable",
					"toolbarSticky": true,
					"items": {
						"fields": {
							"image": {
								"type": "camera"
							}
						}
	
					}
				},
				"locker_clearance_locker_number": {
					"dependencies": {
						"form_type": ["Locker Clearance"]
					}
				},
				"locker_clearance_cleaner_signature_name": {
					"type": "signature",
					"dependencies": {
						"form_type": ["Locker Clearance"]
					}
				}
			}
		}
	}`

	data := `{
		"site": "Holborn",
		"form_type": "Incident Report",
		"incident_report_completed_by": "Some person",
		"incident_report_position": "Some position",
		"incident_report_datetime": "21/11/20",
		"incident_report_incident_type": "Other",
		"incident_report_incident_other": "Other type"
	}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"comments":"Bleh","date":"12/04/18","form_reference":"PS HR Dec 17 Ref:001","hazard_type":"Some hazard type","location":"Some location","picture":"[Image]","risk_profile":"Low - 6 Hours","site":"Holborn","time":"00:12:40"}` {
		t.Fatalf(`Should return {"comments":"Bleh","date":"12/04/18","form_reference":"PS HR Dec 17 Ref:001","hazard_type":"Some hazard type","location":"Some location","picture":"[Image]","risk_profile":"Low - 6 Hours","site":"Holborn","time":"00:12:40"}, instead returned %s`, result)
	}
}
