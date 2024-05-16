package maximum_number_of_weeks_for_which_you_can_work_mid

import (
	"slices"
	"testing"
)

func numberOfWeeks(milestones []int) int64 {
	sum := 0
	for _, x := range milestones {
		sum += x
	}
	mx := slices.Max(milestones)
	if 2*mx > sum+1 {
		return int64(sum-mx)*2 + 1
	}
	return int64(sum)
}

func Test_maximum_number_of_weeks_for_which_you_can_work(t *testing.T) {
	tests := []struct{ milestones []int }{
		{[]int{1, 2, 3}},
		{[]int{5, 2, 1}},
	}
	for _, tt := range tests {
		t.Log(numberOfWeeks(tt.milestones))
	}
}
