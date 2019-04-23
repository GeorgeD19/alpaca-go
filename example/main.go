package main

import (
	"fmt"
	"net/http"

	alpaca "github.com/GeorgeD19/alpaca-go"
)

func main() {
	schema := `{"schema":{"type":"object","properties":{"complete_all_fields":{"type":"information","title":"Please complete all fields."},"report_completed_by":{"type":"string","title":"Report completed by:","required":true},"position":{"type":"string","title":"Position","required":true},"date_time":{"type":"string","title":"Date and Time of Incident","required":true},"site_venue_name":{"type":"string","title":"Site/Venue Name","required":true},"site_venue_address":{"type":"string","title":"Site/Venue Address"},"venue_area":{"type":"information","title":"Please select the venue's area? (If you are unsure please contact control)"},"securigroup_area":{"type":"string","title":"SecuriGroup Area","enum":["Central (Leisure)","Sierra West (Guarding)"],"required":true},"securigroup_manager":{"type":"string","title":"SecuriGroup Manager","required":true},"client_informed":{"type":"string","title":"Has the client been informed?","enum":["Yes","No"],"required":true},"client_informed_yes":{"type":"array","title":"Client Informed","items":{"type":"object","properties":{"client_manager_name":{"type":"string","title":"Client Manager Name","required":true},"client_manager_position":{"type":"string","title":"Client Manager Position","required":true},"date_time_informed":{"type":"string","title":"Date and Time Informed?","required":true}}}},"incident_type":{"type":"array","enum":["Fire","Flood","Intruder","Injury","Drugs","Weapon","Suspicious Behaviour","Physical Violence","Threat of Violence","Verbal Abuse","Near-Miss","Theft","Other"],"required":true},"incident_type_other":{"type":"string","title":"Please specify","required":true},"emergency_services_contacted":{"type":"string","title":"Emergency Services Contacted","enum":["Yes","No"],"required":true},"emergency_services_contacted_yes":{"type":"array","title":"Emergency Services","items":{"type":"object","properties":{"escalated_to":{"type":"string","title":"Escalated to:","required":true},"ambulance_service":{"type":"string","title":"Ambulance Service","enum":["Yes","No"],"required":true},"ambulance_service_yes":{"type":"array","title":"Ambulance Called","items":{"type":"object","properties":{"ambulance_time_called":{"title":"Time Called","required":true},"ambulance_arrived":{"title":"Arrived on site at:","required":true},"ambulance_left":{"title":"Left site at:","required":true},"ambulance_callsign":{"type":"string","title":"Callsign/s"}}}},"fire_service":{"type":"string","title":"Fire Service","enum":["Yes","No"],"required":true},"fire_service_yes":{"type":"array","title":"Fire Service Called","items":{"type":"object","properties":{"fire_time_called":{"title":"Time Called","required":true},"fire_arrived":{"title":"Arrived on site at:","required":true},"fire_left":{"title":"Left site at:","required":true},"fire_callsign":{"type":"string","title":"Callsign/s"}}}},"police":{"type":"string","title":"Police","enum":["Yes","No"],"required":true},"police_yes":{"type":"array","title":"Police Called","items":{"type":"object","properties":{"police_called":{"title":"Time Called","required":true},"police_arrived":{"title":"Arrived on site at:","required":true},"police_left":{"title":"Left site at:","required":true},"police_number":{"type":"string","title":"Officer No/s"}}}}},"dependencies":{"ambulance_service_yes":["ambulance_service"],"fire_service_yes":["fire_service"],"police_yes":["police"]}}},"incident_details":{"type":"information","title":"Incident Details. Complete the below boxes, providing as much detail as you can. "},"who":{"type":"string","title":"Who?","required":true},"where":{"type":"string","title":"Where?","required":true},"what":{"type":"string","title":"What?","required":true},"officer_signature":{"type":"string","title":"Officer Signature","required":true},"form_ref":{"type":"string","title":"Form Reference","readonly":true,"default":"SL Feb 17 Ref:C040"}},"dependencies":{"client_informed_yes":["client_informed"],"incident_type_other":["incident_type"],"emergency_services_contacted_yes":["emergency_services_contacted"]}},"options":{"fields":{"complete_all_fields":{"order":0},"report_completed_by":{"type":"text","order":1},"position":{"type":"text","order":2},"date_time":{"type":"datetime","order":3},"site_venue_name":{"type":"text","order":4},"site_venue_address":{"type":"text","order":5},"venue_area":{"order":6},"securigroup_area":{"type":"radio","optionLabels":["Central (Leisure)","Sierra West (Guarding)"],"vertical":true,"sort":false,"hideNone":true,"order":7},"securigroup_manager":{"type":"text","order":8},"client_informed":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"hideNone":true,"order":9},"client_informed_yes":{"type":"array","dependencies":{"client_informed":["Yes"]},"toolbarSticky":true,"items":{"fields":{"client_manager_name":{"type":"text"},"client_manager_position":{"type":"text"},"date_time_informed":{"type":"text"}}},"order":10},"incident_type":{"type":"checkbox","rightLabel":"Incident Type (check all that apply)","sort":false,"order":11},"incident_type_other":{"type":"text","dependencies":{"incident_type":["Other"]},"order":12},"emergency_services_contacted":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"hideNone":true,"order":13},"emergency_services_contacted_yes":{"type":"array","dependencies":{"emergency_services_contacted":["Yes"]},"toolbarSticky":true,"items":{"fields":{"escalated_to":{"type":"text"},"ambulance_service":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"hideNone":true},"ambulance_service_yes":{"type":"array","dependencies":{"ambulance_service":["Yes"]},"toolbarSticky":true,"items":{"fields":{"ambulance_time_called":{"type":"time"},"ambulance_arrived":{"type":"time"},"ambulance_left":{"type":"time"},"ambulance_callsign":{"type":"text"}}}},"fire_service":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"hideNone":true},"fire_service_yes":{"type":"array","dependencies":{"fire_service":["Yes"]},"toolbarSticky":true,"items":{"fields":{"fire_time_called":{"type":"time"},"fire_arrived":{"type":"time"},"fire_left":{"type":"time"},"fire_callsign":{"type":"text"}}}},"police":{"type":"radio","optionLabels":["Yes","No"],"vertical":true,"sort":false,"hideNone":true},"police_yes":{"type":"array","dependencies":{"police":["Yes"]},"toolbarSticky":true,"items":{"fields":{"police_called":{"type":"time"},"police_arrived":{"type":"time"},"police_left":{"type":"time"},"police_number":{"type":"text"}}}}}},"order":14},"incident_details":{"order":15},"who":{"type":"text","helper":"Who was involved/witnessed the incident (Provide descriptions if names not taken)","helpersPosition":"above","order":16},"where":{"type":"text","helper":"Where? (Exactly where on site did the incident occur?)","helpersPosition":"above","order":17},"what":{"type":"text","helper":"What? (What happened? Start from the beginning and work through events as they occurred)","helpersPosition":"above","order":18},"officer_signature":{"type":"signature","order":19},"form_ref":{"type":"text","order":20}}}}`
	data := `{
		"client_informed": "Yes",
		"client_informed_yes": [
		  {
			"client_manager_name": "Bshsiddbff",
			"client_manager_position": "Hdhdjdjfbfbf ",
			"date_time_informed": "20/10/12 2pm"
		  }
		],
		"incident_type": "Fire,Flood,Intruder",
		"date_time": "2019-04-02T12:47",
		"emergency_services_contacted": "No",
		"form_ref": "SL Feb 17 Ref:C040",
		"officer_signature": "[Signature]",
		"position": "Test",
		"report_completed_by": "Testing ",
		"securigroup_area": "Central (Leisure)",
		"securigroup_manager": "Sgsudjdofb",
		"site_venue_name": "Test",
		"what": "Hdodoedifhs ",
		"where": "Yshdidkforbed ",
		"who": "Gsjdgdude b"
	}`

	r := http.Request{}
	alpaca, err := alpaca.New(alpaca.AlpacaOptions{Schema: schema, Data: data, Request: &r})
	fmt.Println(alpaca.Parse())
	fmt.Println(err)
}
