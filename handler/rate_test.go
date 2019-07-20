package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kakhaUrigashvili/go-rest-api/model"
	"github.com/stretchr/testify/assert"
)


func TestSearchRateHandler(t *testing.T) {
	rates = model.RateCollection{Rates: []model.Rate{
		{
			Days: "mon,tues,thurs",
			Times: "0900-2100",
			TimeZone: "America/Chicago",
			Price: 1500,
		},
		{
			Days: "fri,sat,sun",
			Times: "0900-2100",
			TimeZone: "America/Chicago",
			Price: 2000,
		},
		{
			Days: "wed",
			Times: "0600-1800",
			TimeZone: "America/Chicago",
			Price: 1750,
		},
		{
			Days: "mon,wed,sat",
			Times: "0100-0500",
			TimeZone: "America/Chicago",
			Price: 1000,
		},
		{
			Days: "sun,tues",
			Times: "0100-0700",
			TimeZone: "America/Chicago",
			Price: 925,
		},
	}}
	req, err := http.NewRequest("GET", "/rates/search", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("start", "2015-07-01T07:00:00-05:00")
	q.Add("end", "2015-07-01T12:00:00-05:00")
	req.URL.RawQuery = q.Encode()
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchRateHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect.
	expected := `1750`
	assert.Equal(t, expected, strings.Trim(rr.Body.String(), "\n"))

}
func TestGetRatesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/rates", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRatesHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect.
	expected := `{"rates":[]}`
	assert.Equal(t, expected, strings.Trim(rr.Body.String(), "\n"))

}
func TestCreateRatesHandler(t *testing.T) {

	var rates model.RateCollection = model.RateCollection{Rates: []model.Rate{
		{
			Days:     "mon,tues,thurs",
			Times:    "0900-2100",
			TimeZone: "America/Chicago",
			Price:    1500,
		},
		{
			Days:     "fri,sat,sun",
			Times:    "0900-2100",
			TimeZone: "America/Chicago",
			Price:    2000,
		},
	}}

	jsonRates, _ := json.Marshal(rates)
	req, err := http.NewRequest("POST", "/rates", bytes.NewBuffer(jsonRates))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateRatesHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect.
	expected := `{"rates":[{"days":"mon,tues,thurs","times":"0900-2100","tz":"America/Chicago","price":1500},{"days":"fri,sat,sun","times":"0900-2100","tz":"America/Chicago","price":2000}]}`
	assert.Equal(t, expected, strings.Trim(rr.Body.String(), "\n"))
}
