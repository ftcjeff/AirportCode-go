package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airportCodes); err != nil {
		panic(err)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airportCode := strings.Trim(vars["airportCode"], " ")
	airportInfo := RepoFindAirportCode(airportCode)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airportInfo); err != nil {
  	panic(err)
	}
}

func City(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := strings.Trim(vars["city"], " ")
	airportInfo := RepoFindAirportCodesByCity(city)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airportInfo); err != nil {
  	panic(err)
	}
}

func State(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	state := strings.Trim(vars["state"], " ")
	airportInfo := RepoFindAirportCodesByState(state)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airportInfo); err != nil {
  	panic(err)
	}
}

func Country(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := strings.Trim(vars["country"], " ")
	airportInfo := RepoFindAirportCodesByCountry(country)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airportInfo); err != nil {
  	panic(err)
	}
}

func Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airportCode := strings.Trim(vars["airportCode"], " ")
	RepoDestroyAirportCode(airportCode)
	fmt.Fprintln(w, "AirportCode Removed [", airportCode, "]")
}

func Create(w http.ResponseWriter, r *http.Request) {
	var airportCode AirportCode
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &airportCode); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateAirportCode(airportCode)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
