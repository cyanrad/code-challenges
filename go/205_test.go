package leet

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func isIsomorphic(ts *testing.T, s string, t string) bool {
	runemap := make(map[rune]int)
	order := make([]int, 0, len(s))
	counter := 0

	for _, c := range s {
		ts.Log(c)
		if id, ok := runemap[c]; !ok {
			runemap[c] = counter
			order = append(order, counter)
			counter++
		} else {
			order = append(order, id)
		}
	}
	ts.Log(order)

	runemap = make(map[rune]int)
	counter = 0
	for i, c := range t {
		current := 0
		if id, ok := runemap[c]; !ok {
			runemap[c] = counter
			current = counter

			counter++
		} else {
			current = id
		}
		if order[i] != current {
			return false
		}
	}
	return true
}

func betterIsIsomorphic(s string, t string) bool {

	var arrS, arrT [128]byte

	for i := range s {

		vS, vT := s[i], t[i]

		if (arrS[vS] != 0 && arrS[vS] != vT) || (arrT[vT] != 0 && arrT[vT] != vS) {
			return false
		} else {
			arrS[vS], arrT[vT] = vT, vS
		}

	}

	return true
}

func TestIsIsomorphic(t *testing.T) {
	string1 := "egg"
	string2 := "add"
	expected1 := true

	string3 := "paper"
	string4 := "title"
	expected2 := true

	string5 := "foo"
	string6 := "bar"
	expected3 := false

	require.Equal(t, expected1, isIsomorphic(t, string1, string2))
	require.Equal(t, expected2, isIsomorphic(t, string3, string4))
	require.Equal(t, expected3, isIsomorphic(t, string5, string6))

}

func TestBetterIsIsomorphic(t *testing.T) {
	string1 := "egg"
	string2 := "add"
	expected1 := true

	string3 := "paper"
	string4 := "title"
	expected2 := true

	string5 := "foo"
	string6 := "bar"
	expected3 := false

	require.Equal(t, expected1, betterIsIsomorphic(string1, string2))
	require.Equal(t, expected2, betterIsIsomorphic(string3, string4))
	require.Equal(t, expected3, betterIsIsomorphic(string5, string6))
}
