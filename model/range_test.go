package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRangeWeekdayAbbreviation(t *testing.T) {

	type test struct {
		dateRange Range
		expected  string
	}

	// table driven tests
	tests := []test{
		{
			dateRange: 
				Range{
				StartTime: time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC),
				}, 
			expected: "mon",
		},
		{
			dateRange: 
				Range{
				StartTime: time.Date(2016, time.August, 16, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2016, time.August, 16, 0, 0, 0, 0, time.UTC),
				}, 
			expected: "tues",
		},
		{
			dateRange: 
				Range{
				StartTime: time.Date(2016, time.August, 17, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2016, time.August, 17, 0, 0, 0, 0, time.UTC),
				}, 
			expected: "wed",
		},
		{
			dateRange: 
				Range{
				StartTime: time.Date(2016, time.August, 18, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2016, time.August, 18, 0, 0, 0, 0, time.UTC),
				}, 
			expected: "thurs",
		},
		{
			dateRange: 
				Range{
				StartTime: time.Date(2016, time.August, 19, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2016, time.August, 19, 0, 0, 0, 0, time.UTC),
				}, 
			expected: "fri",
		},
		{
			dateRange: 
				Range{
				StartTime: time.Date(2016, time.August, 20, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2016, time.August, 20, 0, 0, 0, 0, time.UTC),
				}, 
			expected: "sat",
		},
		{
			dateRange: 
				Range{
				StartTime: time.Date(2016, time.August, 21, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2016, time.August, 21, 0, 0, 0, 0, time.UTC),
				}, 
			expected: "sun",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.dateRange.WeekdayAbbreviation())
	}
}


func TestRangeHourMinuteStart(t *testing.T) {
	dateRange := Range{
		StartTime: time.Date(2016, time.August, 15, 13, 5, 0, 0, time.UTC),
		}
	assert.Equal(t, 1305, dateRange.HourMinuteStart())
}

func TestRangeHourMinuteEnd(t *testing.T) {
	dateRange := Range{
		EndTime: time.Date(2016, time.August, 15, 20, 22, 0, 0, time.UTC),
		}
	assert.Equal(t, 2022, dateRange.HourMinuteEnd())
}