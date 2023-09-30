package leet

import (
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	numbersSet := map[int]int{}
	for _, n := range nums {
		numbersSet[n]++
	}

	result := [][]int{}
	exsists := map[[2]int]bool{}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			z := -(nums[i] + nums[j])

			if _, ok := numbersSet[z]; !ok {
				continue
			}

			count := numbersSet[z]
			if z == nums[i] || z == nums[j] {
				if count < 2 {
					continue
				}
				if nums[i] == nums[j] && count < 3 {
					continue
				}
			}

			temp := []int{nums[i], nums[j], z}
			sort.Ints(temp)
			key := [2]int{temp[0], temp[2]}

			if exsists[key] {
				continue
			}

			exsists[key] = true
			result = append(result, temp)
		}
	}

	return result
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	marked := map[[3]int]bool{}
	result := [][]int{}

	for i, j := 0, len(nums)-1; j > i; {
		z := -(nums[i] + nums[j])
		if z > nums[j] || marked[[3]int{nums[i], z, nums[j]}] {
			i++
			continue
		}

		if !rangedBinarySearch(nums, i+1, j-1, z) {
			i++
			continue
		}

		result = append(result, []int{nums[i], z, nums[j]})
		marked[[3]int{nums[i], z, nums[j]}] = true
		j--
	}

	return result
}

func rangedBinarySearch(arr []int, start int, end int, target int) bool {
	for end >= start {
		curr := (start + end) / 2
		if arr[curr] == target {
			return true
		} else if arr[curr] > target {
			end = curr - 1
		} else {
			start = curr + 1
		}
	}
	return false
}

func threeSum3(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		start := i + 1
		end := len(nums) - 1
		for end > start {
			sum := nums[i] + nums[start] + nums[end]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[start], nums[end]})

				start++
				end--
				for end > start && nums[end] == nums[end+1] {
					end--
				}
				for end > start && nums[start] == nums[start-1] {
					start++
				}
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

func TestThreeSum(t *testing.T) {
	val := threeSum3([]int{-2, 0, 1, 1, 2})
	t.Log(val)
}
