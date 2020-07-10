package pkg

import (
	"fmt"
	"strings"
	"testing"
)

// TODO
func TestIsMatch(t *testing.T) {
	ss := []struct {
		s string
		p string
	}{
		{s: "", p: ""},
		{s: "", p: "a*"},
		{s: "", p: "p."},
		{s: "", p: "a*p.a*"},
		{s: "aa", p: "a*"},
		{s: "p", p: "p."},
		{s: "", p: "a*p.a*"},
		{s: "abc", p: "abc*abc*"},
		{s: "abcabcp", p: "abc*p."},
		{s: "abcabc", p: "abc*abc"},
		{s: "acdf", p: "acd."},
	}

	for _, v := range ss {
		fmt.Println("`"+v.s+"`", "`"+v.p+"`", isMatch(v.s, v.p))
	}

}

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	switch {
	case pLen == 0 && sLen == 0:
		return true
	case pLen == 0 && sLen != 0:
		return false
	case pLen != 0 && sLen == 0:
		if p[0] == '*' && !strings.Contains(p, ".") {
			return true
		}
		return false
	case pLen != 0 && sLen != 0:
		switch {
		case p[0] == '*':
		case p[0] == '.':
		case p[0] == s[0]:
		default:
			return false
		}
		sNew := s[1:]
		pNew := p[1:]
		return isMatch(sNew, pNew)
	}
	return false
}
