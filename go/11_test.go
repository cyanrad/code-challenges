package leet

import "testing"

func maxArea(height []int) int {
	max := 0
	for i, j := 0, len(height)-1; j > i; {
		sum := (j - i) * min(height[i], height[j])
		if sum > max {
			max = sum
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func TestMaxArea(t *testing.T) {
	val := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	t.Log(val)
}
