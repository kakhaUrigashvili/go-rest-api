package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/schema"
	"github.com/kakhaUrigashvili/go-rest-api/model"
)

var decoder = schema.NewDecoder()

var rates model.RateCollection = model.RateCollection{Rates: []model.Rate{}}

func dateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
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

func calculateRate(dateRange model.Range) string {

	unavailable := "unavailable"
	var res string

	// input cannot span multiple days
	if !dateEqual(dateRange.StartTime, dateRange.EndTime) {
		return unavailable
	}

	weekday := dateRange.WeekdayAbbreviation()
	start := dateRange.HourMinuteStart()
	end := dateRange.HourMinuteEnd()

	found := false
	
	// searching for rate
	for _, rate := range rates.Rates {
		if strings.Contains(rate.Days, weekday) &&
			start >= rate.HourMinuteStart() && 
			end <= rate.HourMinuteEnd() {
			// checking if we already have rate that we found before
		    if found {
				return unavailable // returning unavailable since found more than 1 rate
			}		
			found = true
			res = strconv.Itoa(rate.Price)
		}
	}
	
	if !found {
		return unavailable
	}

	return res
}

// SearchRateHandler searches rate based on start and end time
func SearchRateHandler(w http.ResponseWriter, r *http.Request) {
	err := validateDateParameters(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dateRange model.Range

	err = decoder.Decode(&dateRange, r.URL.Query())

	if err != nil {
		http.Error(w, "Unable process request", http.StatusBadRequest)
		return
	}

	res := calculateRate(dateRange)
	w.Write([]byte(res))
}

// GetRatesHandler gets all rates
func GetRatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rates)
}

// CreateRatesHandler updates all rates
func CreateRatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRates model.RateCollection
	_ = json.NewDecoder(r.Body).Decode(&newRates)
	rates = newRates
	json.NewEncoder(w).Encode(rates)
}

// func SetRates (r model.RateCollection) {
//   rates = r
// }