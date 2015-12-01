package main

type AirportCode struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	City           string `json:"city"`
	Country        string `json:"country"`
	IATA           string `json:"iata"`
	ICAO           string `json:"icao"`
	Lat            string `json:"latitude"`
	Lon            string `json:"longitude"`
	Altitude       string `json:"altitude"`
	TimezoneOffset string `json:"timezone_offset"`
	DST            string `json:"dst"`
	Timezone       string `json:"timezone"`
}

type AirportCodes []AirportCode
