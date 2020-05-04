package leetcode

import (
	"fmt"
	"testing"
)

func TestShortestPalindrome(t *testing.T) {
	ss := []string{
		"",
		"abcgfg",
		"aacecaaa",
		"abcd",
	}
	for _, v := range ss {
		fmt.Println(shortestPalindrome(v))
	}

}

func shortestPalindrome(s string) string {
	rs := reverseStr(s)
	size := len(s)
	for i := 0; i < size; i++ {
		if s[:size-i] == rs[i:] {
			return rs[:i] + s
		}
	}

	return rs + s
}
func reverseStr(s string) string {
	var str string
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}

func shortestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}
	rs := make([]byte, 0)
	sNew := []byte{'#'}
	for _, v := range s {
		sNew = append(sNew, byte(v), '#')
	}
	fmt.Println(string(sNew))

	length := len(sNew)
	split := 1
	ok := false
	for i := length / 2; i > 0; i-- {
		fmt.Println(string(sNew[:i]), string(sNew[i]), string(sNew[i+1:]))
		l := i - 1
		r := i + 1
		for sNew[l] == sNew[r] {
			fmt.Println("need move")
			if l == 0 { // || r == length-1
				split = i
				ok = true
				break
			}
			l--
			r++
		}
		if ok {
			break
		}
	}
	fmt.Println(string(s[split/2]))
	left := make([]byte, 0)
	splitBytes := sNew[2*split:]
	for i := len(splitBytes) - 1; i >= 0; i-- {
		left = append(left, splitBytes[i])
	}
	fmt.Println(string(left), string(sNew))
	for _, v := range append(left, sNew...) {
		if v != '#' {
			rs = append(rs, v)
		}
	}

	return string(rs)
}
