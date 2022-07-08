package movingaveragefromdatastreamez

import "testing"

type MovingAverage struct {
	Size int
	Sum  int
	Arr  []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{Size: size, Sum: 0, Arr: make([]int, 0, size+1)}
}

func (this *MovingAverage) Next(val int) float64 {
	this.Arr = append(this.Arr, val)
	this.Sum += val
	if len(this.Arr) > this.Size {
		this.Sum -= this.Arr[0]
		this.Arr = this.Arr[1:]
	}
	return float64(this.Sum) / float64(len(this.Arr))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
func Test_moving_average_from_data_stream(t *testing.T) {
	var sol MovingAverage
	c1 := []string{"MovingAverage", "next", "next", "next", "next"}
	c2 := []int{3, 1, 10, 3, 5}
	for i, c := range c1 {
		switch c {
		case "MovingAverage":
			sol = Constructor(c2[i])
		case "next":
			t.Log(sol.Next(c2[i]))
		}
	}
}
