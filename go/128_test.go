package leet

import (
	"log"
	"testing"
)

type links struct {
	before nil_int
	after  nil_int
}
type nil_int struct {
	Int   int
	valid bool
}

func longestConsecutive(nums []int) int {
	chains := map[int]links{}
	for _, n := range nums {
		tempLinks := links{}
		if _, ok := chains[n]; ok {
			continue
		}

		if v, ok := chains[n-1]; ok {
			tempLinks.before = nil_int{n - 1, true}
			v.after.valid = true
			v.after.Int = n
			chains[n-1] = v
		}
		if v, ok := chains[n+1]; ok {
			tempLinks.after = nil_int{n + 1, true}
			v.before.valid = true
			v.before.Int = n
			chains[n+1] = v
		}

		chains[n] = tempLinks
	}

	log.Println(chains)

	longest := 1
	marked := map[int]struct{}{}
	for k, v := range chains {
		if _, ok := marked[k]; ok {
			continue
		}

		marked[k] = struct{}{}
		chainLength := 1

		iterator := v.after
		for iterator.valid {
			chainLength++
			marked[iterator.Int+1] = struct{}{}
			iterator.valid = chains[iterator.Int].after.valid
			iterator.Int++
		}

		iterator = v.before
		for iterator.valid {
			marked[iterator.Int-1] = struct{}{}
			chainLength++
			iterator.valid = chains[iterator.Int].before.valid
			iterator.Int--
		}

		if chainLength > longest {
			longest = chainLength
		}
	}

	return longest
}

func longestConsecutive_2(nums []int) int {
	chains := map[int]bool{}
	for _, n := range nums {
		chains[n] = true
	}

	longest := 0
	for n := range chains {
		if _, ok := chains[n-1]; ok {
			continue
		}

		curr := 1
		for {
			if _, ok := chains[n+curr]; !ok {
				break
			}
			curr++
		}

		if curr > longest {
			longest = curr
		}
	}

	return longest
}

func TestLongestConsecutive(t *testing.T) {
	val := longestConsecutive_2([]int{100, 4, 200, 1, 3, 2})
	log.Println(val)
}
