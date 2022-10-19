package numberofstudentsunabletoeatlunch

import "testing"

func countStudents(students []int, sandwiches []int) (ans int) {
	j, cnt := 0, 0
	for 0 < len(students) && cnt < len(students) {
		cur := students[0]
		students = students[1:]
		if cur == sandwiches[j] {
			j++
			cnt = 0
			continue
		} else {
			students = append(students, cur)
			cnt++
		}
	}
	return len(students)
}
func Test_number_of_students_unable_to_eat_lunch(t *testing.T) {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	t.Log(countStudents(students, sandwiches))
	students = []int{1, 1, 1, 0, 0, 1}
	sandwiches = []int{1, 0, 0, 0, 1, 1}
	t.Log(countStudents(students, sandwiches))
}
