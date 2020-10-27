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

func TestArrayCameraField(t *testing.T) {
	schema := `{"schema":{"type":"object","properties":{"property_address":{"type":"string","title":"Property Address","enum":["Colliers - Fife Energy Park, High Street, Methil, KY8 3RA","N/A"],"required":true},"site_name":{"type":"string","title":"Site Name","required":true},"date_time":{"type":"string","title":"Date / Time"},"location":{"type":"string","title":"Location","enum":["External","Internal"],"required":true},"building_intact":{"type":"string","title":"Building intact (including windows)? ","enum":["Yes","No"],"required":true},"building_intact_comments":{"type":"string","title":"Comments"},"site_secure":{"type":"string","title":"Is the site secure","enum":["Yes","No"],"required":true},"site_secure_comments":{"type":"string","title":"Comments"},"damaged_cladding":{"type":"string","title":"Is there damage to the cladding / brickwork?","enum":["Yes","No"],"required":true},"damaged_cladding_comments":{"type":"string","title":"Comments"},"gutters_downpipes":{"type":"string","title":"Damage to gutters or down-pipes?","enum":["Yes","No"],"required":true},"gutters_downpipes_comments":{"type":"string","title":"Comments"},"graffiti_vandalism":{"type":"string","title":"Graffiti or vandalism damage?","enum":["Yes","No"],"required":true},"graffiti_vandalism_comments":{"type":"string","title":"Comments"},"unlawful_entry":{"type":"string","title":"Signs of unlawful entry?","enum":["Yes","No"],"required":true},"unlawful_entry_comments":{"type":"string","title":"Comments"},"over_grown_landscape":{"type":"string","title":"Is there overgrown landscaping?","enum":["Yes","No"],"required":true},"over_grown_landscape_comments":{"type":"string","title":"Comments"},"exterior_light":{"type":"string","title":"Is there exterior lighting?","enum":["Yes","No"],"required":true},"exterior_light_comments":{"type":"string","title":"Comments"},"exterior_cctv":{"type":"string","title":"Is there exterior CCTV?","enum":["Yes","No"],"required":true},"exterior_cctv_comments":{"type":"string","title":"Comments"},"public_hazards":{"type":"string","title":"Are there any hazards to the public?","enum":["Yes","No"],"required":true},"public_hazards_comments":{"type":"string","title":"Comments"},"external_photos":{"type":"array","title":"External Photos","maxItems":10,"items":{"type":"object","properties":{"subject":{"type":"string","title":"Subject"},"photo":{"type":"string","title":"Photo"}}}},"windows_doors":{"type":"string","title":"Condition of windows and doors","required":true},"comments":{"type":"string","title":"Comments"},"list_of_electrical":{"type":"array","title":"Electrical Devices","maxItems":10,"items":{"type":"object","properties":{"electrical_device":{"type":"string","title":"Device"},"is_it_on_off":{"type":"string","title":"Switched on or off?","enum":["On","Off"]},"electrical_device_comments":{"type":"string","title":"Comments"},"electrical_device_photo":{"type":"string","title":"Photo"}}}},"electricals_comments":{"type":"string","title":"Comments"},"health_and_safety":{"type":"string","title":"Health and Safety issues / cleaning"},"issue_comment":{"type":"string","title":"Comments"},"internal_inspection":{"type":"string","title":"Internal inspection all intact?","enum":["Yes","No"]},"internal_inspection_repeatable":{"type":"array","title":"Internal Inspection","maxItems":10,"items":{"type":"object","properties":{"heating_system":{"type":"string","title":"Heating system operational?","enum":["Yes","No"]},"heating_system_comments":{"type":"string","title":"Comments"},"intruder_alarm":{"type":"string","title":"Intruder alarm system operational?","enum":["Yes","No"]},"intruder_alarm_comments":{"type":"string","title":"Comments"},"fire_detection":{"type":"string","title":"Fire detection system operational?","enum":["Yes","No"]},"fire_detection_comments":{"type":"string","title":"Comments"},"letter_box":{"type":"string","title":"Letter box sealed?","enum":["Yes","No"]},"letter_box_comments":{"type":"string","title":"Comments"},"inernal_photo":{"type":"string","title":"Internal photo"},"internal_subject":{"type":"string","title":"Subject"}},"dependencies":{"heating_system_comments":["heating_system"],"intruder_alarm_comments":["intruder_alarm"],"fire_detection_comments":["fire_detection"],"letter_box_comments":["letter_box"]}}},"gas_isolated":{"type":"string","title":"Gas isolated?","enum":["Yes","No"]},"gas_meter":{"type":"array","title":"Gas meter reading","maxItems":10,"items":{"type":"object","properties":{"meter_number":{"type":"string","title":"Gas meter number"},"meter_reading":{"type":"number","title":"Gas meter reading","required":true},"meter_photo":{"type":"string","title":"Photo"}}}},"electricity_isolated":{"type":"string","title":"Electricity isolated","enum":["Yes","No"]},"electric_meter":{"type":"array","title":"Electricity meter","maxItems":10,"items":{"type":"object","properties":{"electric_number":{"type":"string","title":"Electricity meter number"},"electric_meter":{"type":"number","title":"Electricity meter reading","required":true},"electric_photo":{"type":"string","title":"Photo"}}}},"water_isolated":{"type":"string","title":"Water isolated?"},"post_left":{"type":"string","title":"Post left on site?","enum":["Yes","No"]},"post_left_comment":{"type":"string","title":"Comments"},"documentation_left":{"type":"string","title":"Documentation left on site?","enum":["Yes","No"]},"documentation_left_comments":{"type":"string","title":"Comments"},"valuable_items":{"type":"string","title":"Valuable items left on site?","enum":["Yes","No"]},"valuable_items_comments":{"type":"string","title":"Comments"},"internal_photos_repeatable":{"type":"array","title":"Internal Photos","maxItems":10,"items":{"type":"object","properties":{"internal_subject":{"type":"string","title":"Subject"},"internal_photo":{"type":"string","title":"Photo"}}}},"valuable_items_photo":{"type":"string","title":"Photo"},"further_notes":{"type":"string","title":"Further notes"},"change_in_occupation":{"type":"string","title":"Are there any changes in occupation to the building","enum":["Yes","No"]},"change_in_occupation_comments":{"type":"string","title":"Comments"},"security_on_premises":{"type":"string","title":"Condition and security of vacant premises?"},"other_comments":{"type":"string","title":"Any other comments"},"person_reporting":{"type":"string","title":"Person doing the inspection","required":true},"signature":{"type":"string","title":"Signature","required":true},"form_ref":{"type":"string","title":"Form Reference","readonly":true,"default":"SL Feb 17 Ref:C107"}},"dependencies":{"site_name":["property_address"],"building_intact":["location"],"building_intact_comments":["building_intact"],"site_secure":["location"],"site_secure_comments":["site_secure"],"damaged_cladding":["location"],"damaged_cladding_comments":["damaged_cladding"],"gutters_downpipes":["location"],"gutters_downpipes_comments":["gutters_downpipes"],"graffiti_vandalism":["location"],"graffiti_vandalism_comments":["graffiti_vandalism"],"unlawful_entry":["location"],"unlawful_entry_comments":["unlawful_entry"],"over_grown_landscape":["location"],"over_grown_landscape_comments":["over_grown_landscape"],"exterior_light":["location"],"exterior_light_comments":["exterior_light"],"exterior_cctv":["location"],"exterior_cctv_comments":["exterior_cctv"],"public_hazards":["location"],"public_hazards_comments":["public_hazards"],"external_photos":["location"],"windows_doors":["location"],"comments":["location"],"list_of_electrical":["location"],"electricals_comments":["location"],"health_and_safety":["location"],"issue_comment":["location"],"internal_inspection":["location"],"internal_inspection_repeatable":["internal_inspection"],"gas_isolated":["location"],"gas_meter":["location"],"electricity_isolated":["location"],"electric_meter":["location"],"water_isolated":["location"],"post_left":["location"],"post_left_comment":["post_left"],"documentation_left":["location"],"documentation_left_comments":["documentation_left"],"valuable_items":["location"],"valuable_items_comments":["valuable_items"],"internal_photos_repeatable":["loation"],"valuable_items_photo":["valuable_items"],"change_in_occupation_comments":["change_in_occupation"]}},"options":{"fields":{"property_address":{"type":"radio","optionLabels":["Colliers - Fife Energy Park, High Street, Methil, KY8 3RA","N/A"],"vertical":true,"sort":false,"hideNone":true,"order":0},"site_name":{"type":"text","dependencies":{"property_address":["Other"]},"order":1},"date_time":{"type":"datetime","order":2},"location":{"type":"radio","optionLabels":["External","Internal"],"vertical":true,"sort":false,"hideNone":true,"order":3},"building_intact":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["External"]},"hideNone":true,"order":4},"building_intact_comments":{"type":"text","dependencies":{"building_intact":["No"]},"order":5},"site_secure":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["External"]},"hideNone":true,"order":6},"site_secure_comments":{"type":"text","dependencies":{"site_secure":["No"]},"order":7},"damaged_cladding":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["External"]},"hideNone":true,"order":8},"damaged_cladding_comments":{"type":"text","dependencies":{"damaged_cladding":["Yes"]},"order":9},"gutters_downpipes":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["External"]},"hideNone":true,"order":10},"gutters_downpipes_comments":{"type":"text","dependencies":{"gutters_downpipes":["Yes"]},"order":11},"graffiti_vandalism":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["External"]},"hideNone":true,"order":12},"graffiti_vandalism_comments":{"type":"text","dependencies":{"graffiti_vandalism":["Yes"]},"order":13},"unlawful_entry":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["External"]},"hideNone":true,"order":14},"unlawful_entry_comments":{"type":"text","dependencies":{"unlawful_entry":["Yes"]},"order":15},"over_grown_landscape":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["External"]},"hideNone":true,"order":16},"over_grown_landscape_comments":{"type":"text","dependencies":{"over_grown_landscape":["Yes"]},"order":17},"exterior_light":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["External"]},"hideNone":true,"order":18},"exterior_light_comments":{"type":"text","dependencies":{"exterior_light":["No"]},"order":19},"exterior_cctv":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["External"]},"hideNone":true,"order":20},"exterior_cctv_comments":{"type":"text","dependencies":{"exterior_cctv":["No"]},"order":21},"public_hazards":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["External"]},"hideNone":true,"order":22},"public_hazards_comments":{"type":"text","dependencies":{"public_hazards":["Yes"]},"order":23},"external_photos":{"type":"repeatable","dependencies":{"location":["External"]},"toolbarSticky":true,"items":{"fields":{"subject":{"type":"text"},"photo":{"type":"camera"}}},"order":24},"windows_doors":{"type":"text","dependencies":{"location":["Internal"]},"order":25},"comments":{"type":"text","dependencies":{"location":["Internal"]},"order":26},"list_of_electrical":{"type":"repeatable","dependencies":{"location":["Internal"]},"toolbarSticky":true,"items":{"fields":{"electrical_device":{"type":"text"},"is_it_on_off":{"type":"radio","optionLabels":["On","Off"],"vertical":true,"sort":false,"hideNone":true},"electrical_device_comments":{"type":"text"},"electrical_device_photo":{"type":"camera"}}},"order":27},"electricals_comments":{"type":"text","dependencies":{"location":["Internal"]},"order":28},"health_and_safety":{"type":"text","dependencies":{"location":["Internal"]},"order":29},"issue_comment":{"type":"text","dependencies":{"location":["Internal"]},"order":30},"internal_inspection":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["Internal"]},"hideNone":true,"order":31},"internal_inspection_repeatable":{"type":"repeatable","dependencies":{"internal_inspection":["No"]},"toolbarSticky":true,"items":{"fields":{"heating_system":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"hideNone":true},"heating_system_comments":{"type":"text","dependencies":{"heating_system":["No"]}},"intruder_alarm":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"hideNone":true},"intruder_alarm_comments":{"type":"text","dependencies":{"intruder_alarm":["No"]}},"fire_detection":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"hideNone":true},"fire_detection_comments":{"type":"text","dependencies":{"fire_detection":["No"]}},"letter_box":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"hideNone":true},"letter_box_comments":{"type":"text","dependencies":{"letter_box":["No"]}},"inernal_photo":{"type":"camera"},"internal_subject":{"type":"text"}}},"order":32},"gas_isolated":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["Internal"]},"hideNone":true,"order":33},"gas_meter":{"type":"repeatable","dependencies":{"location":["Internal"]},"toolbarSticky":true,"items":{"fields":{"meter_number":{"type":"text"},"meter_reading":{"type":"number"},"meter_photo":{"type":"camera"}}},"order":34},"electricity_isolated":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["Internal"]},"hideNone":true,"order":35},"electric_meter":{"type":"repeatable","dependencies":{"location":["Internal"]},"toolbarSticky":true,"items":{"fields":{"electric_number":{"type":"text"},"electric_meter":{"type":"number"},"electric_photo":{"type":"camera"}}},"order":36},"water_isolated":{"type":"text","dependencies":{"location":["Internal"]},"order":37},"post_left":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["Internal"]},"hideNone":true,"order":38},"post_left_comment":{"type":"text","dependencies":{"post_left":["Yes"]},"order":39},"documentation_left":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["Internal"]},"hideNone":true,"order":40},"documentation_left_comments":{"type":"text","dependencies":{"documentation_left":["Yes"]},"order":41},"valuable_items":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"dependencies":{"location":["Internal"]},"hideNone":true,"order":42},"valuable_items_comments":{"type":"text","dependencies":{"valuable_items":["Yes"]},"order":43},"internal_photos_repeatable":{"type":"repeatable","dependencies":{"loation":["Internal"]},"toolbarSticky":true,"items":{"fields":{"internal_subject":{"type":"text"},"internal_photo":{"type":"camera"}}},"order":44},"valuable_items_photo":{"type":"camera","dependencies":{"valuable_items":["Yes"]},"order":45},"further_notes":{"type":"text","order":46},"change_in_occupation":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"hideNone":true,"order":47},"change_in_occupation_comments":{"type":"text","dependencies":{"change_in_occupation":["Yes"]},"order":48},"security_on_premises":{"type":"text","order":49},"other_comments":{"type":"text","order":50},"person_reporting":{"type":"text","order":51},"signature":{"type":"signature","order":52},"form_ref":{"type":"text","order":53}}}}`
	data := `{
		"building_intact": "Yes",
		"damaged_cladding": "No",
		"date_time": "2019-03-25T10:58",
		"exterior_cctv": "Yes",
		"exterior_light": "Yes",
		"external_photos": [{
			"photo": "[Image]"
		}],
		"form_ref": "SL Feb 17 Ref:C107",
		"graffiti_vandalism": "No",
		"gutters_downpipes": "No",
		"location": "External",
		"other_comments": "Test",
		"over_grown_landscape": "No",
		"person_reporting": "Test",
		"property_address": "Colliers - Fife Energy Park, High Street, Methil, KY8 3RA",
		"public_hazards": "No",
		"signature": "[Signature]",
		"site_secure": "Yes",
		"unlawful_entry": "No"
	}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestAnyField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"building_intact":"Yes","damaged_cladding":"No","date_time":"2019-03-25T10:58","exterior_cctv":"Yes","exterior_light":"Yes","external_photos":[{"photo":"[Image]"}],"form_ref":"SL Feb 17 Ref:C107","graffiti_vandalism":"No","gutters_downpipes":"No","location":"External","other_comments":"Test","over_grown_landscape":"No","person_reporting":"Test","property_address":"Colliers - Fife Energy Park, High Street, Methil, KY8 3RA","public_hazards":"No","signature":"[Signature]","site_secure":"Yes","unlawful_entry":"No"}` {
		t.Fatalf(`Should return data, instead returned %s`, result)
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
	if result == `"sandwich,cookie,drink"` {
		t.Fatalf(`Should return "sandwich,cookie,drink", instead returned %s`, result)
	}
}

func TestCheckboxArrayField(t *testing.T) {
	schema := `{
		"schema": {
			"type": "object",
			"properties": {
				"field": {
					"type": "array",
					"enum": [
						" Fire",
						"Flood",
						"Intruder",
						"Injury",
						"Drugs",
						"Weapon",
						"Suspicious Behaviour",
						"Physical Violence",
						"Threat of Violence",
						"Verbal Abuse",
						"Near-Miss",
						"Theft",
						"Other"
					],
					"required": true
				}
			}
		},
		"options": {
			"fields": {
				"field": {
					"type": "checkbox",
					"sort": false
				}
			}
		}
	}`
	data := `{
		"field": [
		  {
			"value": " Fire",
			"text": " Fire"
		  },
		  {
			"value": "Threat of Violence",
			"text": "Threat of Violence"
		  }
		]
	}`

	alpaca, err := New(AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		t.Fatalf("TestCheckboxArrayField error: %s", err)
	}

	result := alpaca.Parse()
	if result != `{"field":[{"text":" Fire","value":" Fire"},{"text":"Threat of Violence","value":"Threat of Violence"}]}` {
		t.Fatalf(`Should return {"field":[{"text":" Fire","value":" Fire"},{"text":"Threat of Violence","value":"Threat of Violence"}]}, instead returned %s`, result)
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
