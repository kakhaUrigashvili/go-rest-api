package model

import (
	"strconv"
	"strings"
)

// Rate struct
type Rate struct {
	Days     string `json:"days"`
	Times    string `json:"times"`
	TimeZone string `json:"tz"`
	Price    int    `json:"price"`
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
