package model

import (
	"strconv"
	"strings"
)

// Rate struct
type Rate struct {
	Days     string `json:"days" example:"mon,tues,thurs"`
	Times    string `json:"times" example:"0900-2100"`
	TimeZone string `json:"tz" example:"America/Chicago"`
	Price    int    `json:"price" example:"1500"`
}

// HourMinuteStart returns integer respresenting hour and minute of start time
func (r *Rate) HourMinuteStart() int {
	s := strings.Split(r.Times, "-")[0]
	i, _ := strconv.Atoi(s)
	return i
}

// HourMinuteEnd returns integer respresenting hour and minute of end time
func (r *Rate) HourMinuteEnd() int {
	s := strings.Split(r.Times, "-")[1]
	i, _ := strconv.Atoi(s)
	return i
}

// RateCollection struct
type RateCollection struct {
	Rates []Rate `json:"rates"`
}
