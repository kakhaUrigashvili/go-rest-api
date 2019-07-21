package handler

import (
	"encoding/json"
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

func timeIn(origTime time.Time, name string) time.Time {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return origTime.In(loc)
}

func toWeekdayAbbreviation(time time.Time) string {
	index := int(time.Weekday())
	weekdays := [7]string{"sun", "mon", "tues", "wed", "thurs", "fri", "sat"}
	return weekdays[index]
}

func toHourMinute(time time.Time) int {
	return time.Hour()*100 + time.Minute()
}

func calculateRate(startDate, endDate time.Time) string {
	unavailable := "unavailable"
	var res string
	found := false

	// searching for rate
	for _, rate := range rates.Rates {
		// convert to timezone matching the rates data
		startDateLoc := timeIn(startDate, rate.TimeZone)
		endDateLoc := timeIn(endDate, rate.TimeZone)

		// input cannot span multiple days
		if !dateEqual(startDateLoc, endDateLoc) {
			return unavailable
		}

		weekday := toWeekdayAbbreviation(startDateLoc)
		start := toHourMinute(startDateLoc)
		end := toHourMinute(endDateLoc)

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

// SearchRateHandler godoc
// @Summary Search rate
// @Description search rate by start and end date.
// @Tags rates
// @Produce plain
// @Param start query string true "date/time as ISO-8601 with timezone"
// @Param end query string true "date/time as ISO-8601 with timezone"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /rates/search [get]
func SearchRateHandler(w http.ResponseWriter, r *http.Request) {

	errorMessage := "query parameter is required to be valid ISO-8601 datetime"
	start, err := time.Parse(time.RFC3339, r.URL.Query().Get("start"))
	if err != nil {
		http.Error(w, "start "+errorMessage, http.StatusBadRequest)
		return
	}
	end, err := time.Parse(time.RFC3339, r.URL.Query().Get("end"))
	if err != nil {
		http.Error(w, "end "+errorMessage, http.StatusBadRequest)
		return
	}

	res := calculateRate(start, end)
	w.Write([]byte(res))
}

// GetRatesHandler godoc
// @Summary Get all rates
// @Description get all rates available in the system
// @Tags rates
// @Success 200 {object} model.RateCollection
// @Router /rates [get]
func GetRatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rates)
}

// CreateRatesHandler godoc
// @Summary Upload rates
// @Description upload rates using json payload
// @Tags rates
// @Accept  json
// @Param account body model.RateCollection true "Add rates"
// @Success 201
// @Router /rates [post]
func CreateRatesHandler(w http.ResponseWriter, r *http.Request) {
	var newRates model.RateCollection
	_ = json.NewDecoder(r.Body).Decode(&newRates)
	rates = newRates
	w.WriteHeader(http.StatusCreated)
}
