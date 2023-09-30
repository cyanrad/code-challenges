package leet

import "testing"

func search(nums []int, target int) int {
	max := 0
	for l, r := 0, 1; r < len(nums); {
		diff := nums[r] - nums[l]
		if diff > max && diff > 0 {
			max = diff
		}

		if nums[l] > nums[r] {
			l = r
			r = l + 1
		} else {
			r++
		}
	}
	return max
}

func TestSearch(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	out := search(nums, 9)
	t.Log(out)
}
