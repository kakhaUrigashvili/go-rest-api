package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRateHourMinuteStart(t *testing.T) {
	rate := Rate{Times: "0715-2100"}
	assert.Equal(t, 715, rate.HourMinuteStart())
}

func TestRateHourMinuteEnd(t *testing.T) {
	rate := Rate{Times: "0715-2101"}
	assert.Equal(t, 2101, rate.HourMinuteEnd())
}
