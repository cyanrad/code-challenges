package leet

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func pivotIndex(t *testing.T, nums []int) int {
	left := 0
	right := nums[len(nums)-1]
	pivot := 0
	j := len(nums) - 2

	for pivot < j {
		t.Log(left, right)
		if left >= right {
			right += nums[j]
			j--
		} else if left < right {
			left += nums[pivot]
			pivot++
		}
	}

	if left == right {
		return pivot
	}
	return -1
}

func TestPivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	nums2 := []int{-1, -1, -1, -1, -1, 0}
	expected := 3
	expected2 := 2
	require.Equal(t, expected, pivotIndex(t, nums))
	require.Equal(t, expected2, pivotIndex(t, nums2))
}
