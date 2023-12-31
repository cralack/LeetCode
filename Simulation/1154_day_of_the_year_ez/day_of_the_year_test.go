package day_of_the_year_ez

import (
	"testing"
	"time"
)

func dayOfTheWeek(day, month, year int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return t.Weekday().String()
}

func dayOfYear(date string) int {
	t, _ := time.Parse(time.DateOnly, date)
	return t.YearDay()
}

func Test_day_of_the_year(t *testing.T) {
	tests := []struct {
		date string
	}{
		{"2019-01-09"},
		{"2019-02-10"},
	}
	for _, tt := range tests {
		t.Log(dayOfYear(tt.date))
	}
	t.Log(dayOfTheWeek(31, 8, 2019))
}
