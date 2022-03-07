package findmedianfromdatastream

import "testing"

type MedianFinder struct {
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (p *MedianFinder) AddNum(num int) {

}

func (p *MedianFinder) FindMedian() float64 {
	return 0.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
func Test_find_median_from_data_stream(t *testing.T) {
	finder := Constructor()
	finder.AddNum(1)
	finder.AddNum(2)
	t.Log(finder.FindMedian())
	finder.AddNum(3)
	t.Log(finder.FindMedian())
}
