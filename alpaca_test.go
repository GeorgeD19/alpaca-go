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
	if result != `["Vanilla","Chocolate"]` {
		t.Fatalf(`Should return ["Vanilla","Chocolate"], instead returned %s`, result)
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
// Unsupported
// func TestAddressField(t *testing.T) {
// 	schema := `{
// 		"schema": {
// 			"title": "Home Address",
// 			"type": "any"
// 		},
// 		"options": {
// 			"type": "address"
// 		}
// 	}`
// 	data := `{"street":["street 1","street 2","street 3"],"city":"glasgow","state":"AL","zip":"23233"}`

// 	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
// 	if err != nil {
// 		t.Fatalf("TestAddressField error: %s", err)
// 	}

// 	result := alpaca.Parse()
// 	if result != `{"street":["street 1","street 2","street 3"],"city":"glasgow","state":"AL","zip":"23233"}` {
// 		t.Fatalf(`Should return {"street":["street 1","street 2","street 3"],"city":"glasgow","state":"AL","zip":"23233"}, instead returned %s`, result)
// 	}
// }

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

func TestObjectMixedField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "object",
			"properties": {
				"integerfield": {
					"type": "string"
				},
				"selectfield": {
					"type": "array",
					"items": {
						"title": "Ice Cream",
						"type": "string",
						"enum": ["Vanilla", "Chocolate", "Strawberry", "Mint"]
					},
					"minItems": 2,
					"maxItems": 3
				}
			}
		},
		"options": {
			"fields": {
				"integerfield": {
					"type": "integer"
				},
				"selectfield": {
					"type": "select"
				}
			}
		}
	}`

	data := `{"integerfield":17,"selectfield":["Vanilla","Chocolate"]}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestObjectIntegerField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"integerfield":17,"selectfield":["Vanilla","Chocolate"]}` {
		t.Fatalf(`Should return {"integerfield":17,"selectfield":["Vanilla","Chocolate"]}, instead returned %s`, result)
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

// Unsupported
// func TestMapField(t *testing.T) {
// 	schema := `{
// 		"schema": {
// 			"type": "array",
// 			"items": {
// 				"type": "object",
// 				"properties": {
// 					"_key": {
// 						"title": "User ID",
// 						"type": "string"
// 					},
// 					"firstName": {
// 						"title": "First Name",
// 						"type": "string"
// 					},
// 					"lastName": {
// 						"title": "Last Name",
// 						"type": "string"
// 					},
// 					"gender": {
// 						"title": "Gender",
// 						"type": "string",
// 						"enum": ["Male", "Female"]
// 					}
// 				}
// 			}
// 		},
// 		"options": {
// 			"type": "map",
// 			"toolbarSticky": true,
// 			"items": {
// 				"fields": {
// 					"_key": {
// 						"size": 60,
// 						"helper": "This value serves as a unique key into the associative array."
// 					}
// 				}
// 			}
// 		}
// 	}`
// 	data := `{
//         "john316": {
//             "firstName": "Tim",
//             "lastName": "Tebow",
//             "gender": "Male"
//         },
//         "ladygaga": {
//             "firstName": "Stefani",
//             "lastName": "Germanotta",
//             "gender": "Female"
//         }
//     }`

// 	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
// 	if err != nil {
// 		t.Fatalf("error: %s", err)
// 	}

// 	result := alpaca.Parse()
// 	if result != `19` {
// 		t.Fatalf(`Should return 19, instead returned %s`, result)
// 	}
// }

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
	if result != `["great","wonderful","ice cream"]` {
		t.Fatalf(`Should return ["great","wonderful","ice cream"], instead returned %s`, result)
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
