package pkg

import (
	"github.com/thinkhp/logsim"
	"testing"
)

func BenchmarkLongestPalindrome(b *testing.B) {

	ps := []string{
		"babadada",
		"babad",
		"abb",
		"ccd",
		"ac",
		"babad",
		"cbbd",
		"",
		"ccccccccccccccccccccccccccc",
		"adbgsdasdiadjabhashdhsdhbkdjasd",
	}
	logsim.SetLevelNotPrint(logsim.TraceLevel)
	//thinkLog.SetLevelNotPrint(thinkLog.DebugLevel)
	for _, v := range ps {
		logsim.DebugLog.Println(longestPalindrome(v))
	}
	//return
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range ps {
			longestPalindrome(v)
		}
	}
	b.StopTimer()
}

//BenchmarkLongestPalindrome-4   	    1437	    879787 ns/op
// Manacher's Algorithm 马拉车算法 n
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	start := 2
	maxC, maxR := start, 0
	C, R := start, 0 //center, center回文的半径

	sNew := []byte{'^', '#'}
	for _, v := range s {
		sNew = append(sNew, byte(v), '#')
	}
	sNew = append(sNew, '$')
	logsim.TraceLog.Println(string(sNew))
	P := make([]int, len(sNew))
	for i := start; i < len(sNew)-1; i++ {
		iMirror := 2*C - i
		length := 0
		// 是同一个 center 回文的子集,i.length >= iMirror.length
		if P[iMirror]+i <= C+R {
			length = P[iMirror]
		}
		logsim.TraceLog.Println(i, length, C, R, i+length < C+R, string(sNew[i-length:i+length+1]))

		l := i - length - 1
		r := i + length + 1
		for sNew[l] == sNew[r] {
			logsim.TraceLog.Println("need right")
			length++
			C, R = i, length
			l--
			r++
		}
		P[i] = length
		logsim.TraceLog.Println(i, length, string(sNew[i-length:i+length+1]))
		if length > maxR {
			maxR = length
			maxC = i
		}
		if i+R >= len(sNew)-1 {
			break
		}
	}

	logsim.DebugLog.Println(string(sNew), string(sNew[maxC-maxR:maxC+maxR+1]))
	ss := (maxC - maxR) / 2
	// strings.ReplaceAll(string(sNew[maxC-maxR:maxC+maxR+1]), "#", "")
	return s[ss : ss+maxR]
}

// 动态规划

// BenchmarkLongestPalindrome-4   	  661288	      1709 ns/op
// 从中间开始,左右拓展,确定最长的回文 n*n
func longestPalindrome2(s string) string {
	longest := s[0:1]
	var findL = func(l, r int) {
		for l >= 0 && r <= len(s)-1 {
			if s[l] == s[r] {
				v := s[l : r+1]
				if len(v) > len(longest) {
					longest = v
				}
			} else {
				break
			}
			l--
			r++
		}
	}
	for i, _ := range s {
		l, r := 0, 0
		// 回文串中间为字母
		l = i - 1
		r = i + 1
		findL(l, r)
		// 回文串中间为空
		l, r = i, i+1
		findL(l, r)
	}
	return longest
}

//BenchmarkLongestPalindrome-4   	  135452	      8982 ns/op
// 所有的子串,验证是否回文 n*n*n
func longestPalindrome1(s string) string {
	longest := s[0:1]
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			v := s[i : j+1]
			if isPalindrome(v) && len(v) > len(longest) {
				longest = v
			}

		}
	}
	return longest
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
