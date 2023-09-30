package leet

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func twoSum(nums []int, target int) []int {
	// >> placing them in a set
	set := make(map[int]int, len(nums))
	for _, n := range nums {
		set[n] += 1
	}

	// >> going over the nums
	for i, n := range nums {
		if set[target-n] > 0 {
			j := sort.SearchInts(nums, target-n)
			return []int{i, j}
		}
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	nums1 := []int{2, 7, 11, 15}
	targ1 := 9
	exp1 := []int{0, 1}

	nums2 := []int{3, 2, 4}
	targ2 := 6
	exp2 := []int{1, 2}

	require.Equal(t, exp1, twoSum(nums1, targ1))
	require.Equal(t, exp2, twoSum(nums2, targ2))
}
