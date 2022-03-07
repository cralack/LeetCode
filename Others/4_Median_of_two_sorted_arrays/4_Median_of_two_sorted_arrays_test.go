package medianoftwosortedarrays

import (
	"sort"
	"testing"
)

func findMeditmportedArrays(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if n < m {
		return findMeditmportedArrays(nums2, nums1)
	}

	var midM = (m - 1) / 2
	var midN = (n - 1) / 2

	if m == 0 {
		if n%2 == 1 {
			return float64(nums2[midN])
		}
		return float64(nums2[midN]+nums2[midN+1]) / 2
	}
	if m == 1 || m == 2 {
		if n < 3 {
			for i := 0; i < n; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else if n%2 == 1 {
			for i := midN - 1; i < midN+2; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else {
			for i := midN - 1; i < midN+3; i++ {
				nums1 = append(nums1, nums2[i])
			}
		}
	}
	sort.Ints(nums1)
	m = len(nums1)
	midM = (m - 1) / 2

	if len(nums1)%2 == 1 {
		return float64(nums1[midM])
	} else {
		return float64(nums1[midM]+nums1[midM+1]) / 2
	}
}

func TestMedian(t *testing.T) {
	nums2 := []int{1, 2}
	nums1 := []int{3, 4}
	ans := findMeditmportedArrays(nums1, nums2)
	t.Log(ans)
}
