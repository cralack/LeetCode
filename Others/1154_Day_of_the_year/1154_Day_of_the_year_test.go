package dayoftheyear

import (
	"testing"
)

func TestDayoftheYear(t *testing.T) {
	dates := []string{"2019-01-09", "2019-02-10", "2003-03-01", "2004-03-01"}
	for _, date := range dates {
		t.Log(dayOfYear(date))
	}
}

func dayOfYear(date string) int {
	i := 0
	date_int := [3]int{}
	for _, v := range date {
		if v == '-' {
			i++
		} else {
			date_int[i] = date_int[i]*10 + int(v-'0')
		}
	}

	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if date_int[0]%400 == 0 || (date_int[0]%4 == 0 && date_int[0]%100 != 0) {
		days[2]++
	}
	//fmt.Printf("Year:%d,month:%d,day:%d", year, month, day)
	for i := 1; i < date_int[1]; i++ {
		date_int[2] += days[i]
	}
	return date_int[2]
}
