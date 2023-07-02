package minimumhoursoftrainingtowin

import "testing"

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) (ans int) {
	s := 0
	for _, x := range energy {
		s += x
	}
	if dif := s - initialEnergy + 1; dif > 0 {
		ans += dif
	}
	for _, x := range experience {
		if initialExperience <= x {
			ans += x - initialExperience + 1
			initialExperience = x + 1
		}
		initialExperience += x
	}
	return
}

func Test_minimum_hours_of_training_to_win_a_competition(t *testing.T) {
	// test table
	tests := []struct {
		iiEne, iiExp int
		ene, exp     []int
	}{
		{5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1}},
		{2, 4, []int{1}, []int{3}},
	}

	// table drive test
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			ans := minNumberOfHours(tt.iiEne, tt.iiExp, tt.ene, tt.exp)
			t.Log(ans)
		})
	}
}
