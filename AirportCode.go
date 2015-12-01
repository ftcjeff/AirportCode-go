package main

type AirportCode struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IATA 			string `json:"iata"`
	City 	 		string `json:"city"`
	State 		string `json:"state"`
	Country  	string `json:"country"`
}

type AirportCodes []AirportCode
