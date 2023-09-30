package leet

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func getLetters(s string) []rune {
	r := make([]rune, 0)
	for _, c := range s {
		if unicode.IsLetter(c) {
			r = append(r, unicode.ToLower(c))
		} else if unicode.IsNumber(c) {
			r = append(r, c)
		}
	}
	return r
}

func isPalindrome(s string) bool {
	letters := getLetters(s)
	for i, j := 0, len(letters)-1; j > i; {
		if letters[j] != letters[i] {
			return false
		}
		i++
		j--
	}
	return true
}

func TestIsPalindrom(t *testing.T) {
	s1 := "A man, a plan, a canal: Panama"
	exp := "amanaplanacanalpanama"
	exp1 := true

	require.Equal(t, exp, string(getLetters(exp)))
	require.Equal(t, exp1, isPalindrome(s1))
}
