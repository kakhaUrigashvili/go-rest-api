package model

import "time"

// Range struct
type Range struct {
	StartTime time.Time `schema:"start"`
	EndTime   time.Time `schema:"end"`
}

// WeekdayAbbreviation returns weekday string
func (r *Range) WeekdayAbbreviation() string {
	index := int(r.StartTime.Weekday())
	weekdays := [7]string{"sun", "mon", "tues", "wed", "thurs", "fri", "sat"}
	return weekdays[index]
}

func getHourMinute(time time.Time) int {
	return time.Hour()*100 + time.Minute()
}

// HourMinuteStart returns integer respresenting hour and minute of start time
func (r *Range) HourMinuteStart() int {
	return getHourMinute(r.StartTime)
}

// HourMinuteEnd returns integer respresenting hour and minute of end time
func (r *Range) HourMinuteEnd() int {
	return getHourMinute(r.EndTime)
}
