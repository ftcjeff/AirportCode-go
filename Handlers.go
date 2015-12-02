package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	allAirports := GetAllRows("airport")

	if err := json.NewEncoder(w).Encode(allAirports); err != nil {
		panic(err)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airportCode := strings.Trim(vars["airportCode"], " ")
	airportInfo := GetRows("airport", "IATA", airportCode)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airportInfo); err != nil {
		panic(err)
	}
}

func ShowById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := strings.Trim(vars["Id"], " ")
	airportInfo := GetRows("airport", "Id", id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airportInfo); err != nil {
		panic(err)
	}
}

func City(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := strings.Trim(vars["city"], " ")
	airportInfo := GetRows("airport", "City", city)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airportInfo); err != nil {
		panic(err)
	}
}

func Country(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := strings.Trim(vars["country"], " ")
	airportInfo := GetRows("airport", "Country", country)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airportInfo); err != nil {
		panic(err)
	}
}
