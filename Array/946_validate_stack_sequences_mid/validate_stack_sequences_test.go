package validatestacksequencesmid

import "testing"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}
	return len(stack) == 0
}
func Test_validate_stack_sequences(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	t.Log(validateStackSequences(pushed, popped))
	pushed = []int{1, 2, 3, 4, 5}
	popped = []int{4, 2, 5, 1, 2}
	t.Log(validateStackSequences(pushed, popped))
}
