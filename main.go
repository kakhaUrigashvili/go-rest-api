package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// Rate struct
type Rate struct {
	Days string `json:"days"`
	Times  string `json:"times"`
	TimeZone  string `json:"tz"`
	Price int `json:"price"`
}

// RateCollection struct
type RateCollection struct {
	Rates []Rate `json:"rates"`
}

var rates RateCollection

// RateResult struct
type RateResult struct {
	Rate int `json:"rate"`
}

// RateRequest struct
type RateRequest struct {
	StartTime time.Time `schema:"start"`
	EndTime   time.Time `schema:"end"`
}

func validateDateParameters(w http.ResponseWriter, r *http.Request) error {
	_, err := time.Parse(time.RFC3339, r.URL.Query().Get("start"))
	errorMessage := "%v query parameter is required to be valid ISO-8601 datetime"
	if err != nil {
		return fmt.Errorf(errorMessage, "start")
	}
	_, err = time.Parse(time.RFC3339, r.URL.Query().Get("end"))
	if err != nil {
		return fmt.Errorf(errorMessage, "end")
	}
	return nil
}

// Get rate
func doRateSearch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params)
	err := validateDateParameters(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var rateRequest RateRequest

	err = decoder.Decode(&rateRequest, r.URL.Query())

	if err != nil {
		http.Error(w, "Unable process request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	rate := &RateResult{Rate: 1000}
	json.NewEncoder(w).Encode(rate)
}

// Get all rates
func getRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rates)
}

// Update all rates
func createRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRates RateCollection
	_ = json.NewDecoder(r.Body).Decode(&newRates)
	rates = newRates
	json.NewEncoder(w).Encode(rates) 
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	rates.Rates = append(rates.Rates, Rate{Days: "mon,tues,thurs", Times: "0900-2100", TimeZone: "America/Chicago", Price: 1500})
	rates.Rates = append(rates.Rates, Rate{Days: "fri,sat,sun", Times: "0900-2100", TimeZone: "America/Chicago", Price: 2000})

	routePrefix := "/api/v1"

	// Route handles & endpoints
	r.HandleFunc(routePrefix + "/rates/search", doRateSearch).Methods("GET")
	r.HandleFunc(routePrefix + "/rates", getRates).Methods("GET")
	r.HandleFunc(routePrefix + "/rates", createRates).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
