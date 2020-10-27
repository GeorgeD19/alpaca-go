package main

import (
	"fmt"

	"github.com/GeorgeD19/alpaca-go"
)

func main() {
	schema := `{
		"schema": {
			"type": "object",
			"properties": {
				"field": {
					"type": "array",
					"enum": [
						"Fire",
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
		"field": [{
			"value": "Fire",
			"text": "Fire"
		},
		{
			"value": "Threat of Violence",
			"text": "Threat of Violence"
		}]
	}`

	parser, err := alpaca.New(alpaca.AlpacaOptions{Schema: schema, Data: data})
	if err != nil {
		fmt.Println(err)
	}

	result := parser.Parse()
	fmt.Println(result)
}
