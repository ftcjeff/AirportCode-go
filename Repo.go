package main

import (
	"io/ioutil"
	"strings"
)

var airportCodes AirportCodes

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	tableName := "airport"
	tableFields := []string{"Id", "Name", "City", "Country", "IATA", "ICAO", "Lat", "Lon", "Altitude", "TimezoneOffset", "DST", "Timezone"}

	mysql := GetServiceURI("mysql")
	CreateTable(mysql, "picasso", "picasso", "picasso", tableName)

	dat, err := ioutil.ReadFile("airport_codes.csv")
	check(err)

	ports := strings.Split(string(dat), "\n")

	for _, port := range ports {
		if strings.Contains(port, ",") {
			tokens := strings.Split(port, ",")

			if len(tokens[4]) == 0 {
				continue
			}

			AddRow(tableName, tableFields, tokens)
		}
	}
}
