package main

import (
	"fmt"
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
	dat, err := ioutil.ReadFile("airport_codes.csv")
	check(err)

	ports := strings.Split(string(dat), "\n")

	for _, port := range ports {
		if strings.Contains(port, ",") {
			tokens := strings.Split(port, ",")

			go RepoCreateAirportCode(
			     AirportCode{Id: tokens[0],
					             Name: tokens[1],
											 IATA: tokens[0],
											 City: tokens[2],
											 State: tokens[3],
											 Country: tokens[4]})
		}
	}
}

func RepoFindAirportCodesByCity(city string) AirportCodes {
  var rv AirportCodes

	for _, t := range airportCodes {
		if strings.ToLower(t.City) == strings.ToLower(city) {
  		rv = append(rv, t)
		}
	}

	return rv
}

func RepoFindAirportCodesByState(state string) AirportCodes {
  var rv AirportCodes

	for _, t := range airportCodes {
		if strings.ToLower(t.State) == strings.ToLower(state) {
  		rv = append(rv, t)
		}
	}

	return rv
}

func RepoFindAirportCodesByCountry(country string) AirportCodes {
  var rv AirportCodes

	for _, t := range airportCodes {
		if strings.ToLower(t.Country) == strings.ToLower(country) {
  		rv = append(rv, t)
		}
	}

	return rv
}

func RepoFindAirportCode(id string) AirportCode {
	for _, t := range airportCodes {
		if strings.ToLower(t.Id) == strings.ToLower(id) {
			return t
		}
	}

	return AirportCode{}
}

func RepoCreateAirportCode(t AirportCode) AirportCode {
	fmt.Println("Creating ", t)
	airportCodes = append(airportCodes, t)
	return t
}

func RepoDestroyAirportCode(id string) error {
	for i, t := range airportCodes {
		if t.Id == id {
			airportCodes = append(airportCodes[:i], airportCodes[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find AirportCode with id of %s to delete", id)
}