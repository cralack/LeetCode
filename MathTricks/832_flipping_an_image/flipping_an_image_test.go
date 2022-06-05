package flippinganimage

import "testing"

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)
	for i := range image {
		for j, k := 0, n-1; j <= k; j, k = j+1, k-1 {
			image[i][j], image[i][k] = 1^image[i][k], 1^image[i][j]
		}
	}
	return image
}
func Test_flipping_an_image(t *testing.T) {
	image := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	t.Log(flipAndInvertImage(image))
	image = [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	t.Log(flipAndInvertImage(image))
}
