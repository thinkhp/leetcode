package pkg

import (
	"fmt"
	"testing"
)

func TestStdLibBytes(t *testing.T) {
	fmt.Println()
}

func TestKMP(t *testing.T) {
	ss := []struct {
		s string
		p string
	}{
		{"ABCDABCABB", "ABCABCDBB"},
		{"aabaaabaaac", "abdabad"},
		{"aabaaabaaac", "aabaaac"},
		{s: "", p: ""},
		{s: "mississippi", p: "pi"},
		{s: "substring searching", p: "search"},
		{s: "ABC ABCDAB ABCDABCDABDE", p: "ABCDABD"},
		{s: "ABC ABCDAB ABCDABCDABDE", p: "PARTICIPATE IN PARACHUTE"},
		{s: "AABAACAAAB", p: "AAA"},
		{s: "AAAB", p: "AAAB"},
		{s: "AAAAAAAAB", p: "AAAB"},
	}

	format := "s: %s\np: %s\n%d %s\n"
	for _, v := range ss {
		if len(v.p) == 0 {
			continue
		}
		fmt.Println(v.p)
		fmt.Println(getNext(v.p))
		fmt.Println(getNext1(v.p))
		fmt.Println()
	}
	return
	for _, v := range ss {
		start := KMP(v.s, v.p)
		if start == -1 {
			fmt.Printf(format, v.s, v.p, start, "not")
		} else {
			fmt.Printf(format, v.s, v.p, start, v.s[start:start+len(v.p)])
		}
		fmt.Println("**************************")
	}
}

// 在不错过任何潜在匹配的情况下，我们"预搜索"这个模式串本身并将其译成一个包含所有可能失配的位置对应可以绕过最多无效字符的列表。
//i		0	1	2	3	4	5	6
//W[i]	A	B	C	D	A	B	D
//T[i] -1	0	0	0	0	1	2

//i		0	1	2	3	4	5	6
//W[i]	A	A	B	A	A	A	D
//T[i] -1	0	1	0	1	2	2

//i		0	1	2	3	4	5	6	7	8	9	10	11	12	13	14	15	16	17	18	19	20	21	22	23
//W[i]	P	A	R	T	I	C	I	P	A	T	E		I	N		P	A	R	A	C	H	U	T	E
//T[i]	-1	0	0	0	0	0	0	0	1	2	0	0	0	0	0	0	1	2	3	0	0	0	0	0
func getNext1(str string) []int {
	next := make([]int, len(str)) //0,0,0....0
	// str[k]... 与 str[0]... 的匹配程度
	for i, k := 1, 0; i < len(str); {
		//初始版本
		//if k == 0 {
		//	if str[i] == str[0] {
		//		k = 1
		//		next[i] = k
		//	}
		//} else {
		//	if str[i] == str[k] {
		//		k++
		//		next[i] = k
		//	} else {
		//		k = next[k-1]
		//		i--
		//	}
		//}
		//i++
		fmt.Println(i, k)
		if str[i] == str[k] {
			k++
			next[i] = k
		} else {
			if k != 0 {
				k = next[k-1]
				i--
			}
		}
		i++
	}
	//fmt.Println(next)
	return append([]int{0}, next[:len(next)-1]...)
}

func getNext(str string) []int {
	next := make([]int, 0)
	next = append(next, -1)
	for j, k := 0, -1; j < len(str)-1; {
		if k == -1 || str[k] == str[j] {
			j++
			k++
			next = append(next, k)
		} else {
			k = next[k]
		}
	}
	next[0] = 0
	return next
}

func IndexByte(s string, b byte) int {
	for i, v := range []byte(s) {
		if v == b {
			return i
		}
	}
	return -1
}

//让我们用一个实例来演示这个算法。在任意给定时间，本算法被两个整数m和i所决定：
//
//m代表主文字符串S内匹配字符串W的当前查找位置，
//i代表匹配字符串W当前做比较的字符位置。
//图示如下：
//
//			   1         2
//m: 01234567890123456789012
//S: ABC ABCDAB ABCDABCDABDE
//W: ABCDABD
//i: 0123456
//我们从W与S的开头比较起。我们比对到S[3](=' ')时，发现W[3](='D')与其不符。接着并不是从S[1]比较下去。我们已经知道S[1]~S[3]不与W[0]相合。因此，略过这些字符，令m = 4以及i = 0。
//
//			   1         2
//m: 01234567890123456789012
//S: ABC ABCDAB ABCDABCDABDE
//W:     ABCDABD
//i:     0123456
//如上所示，我们检核了"ABCDAB"这个字符串。然而，这与目标仍有些差异。我们可以注意到，"AB"在字符串头尾处出现了两次。这意味着尾端的"AB"可以作为下次比较的起始点。因此，我们令m = 8, i = 2，继续比较。图标如下：
//
//			   1         2
//m: 01234567890123456789012
//S: ABC ABCDAB ABCDABCDABDE
//W:         ABCDABD
//i:         0123456
//于m = 10的地方，又出现不相符的情况。类似地，令m = 11, i = 0继续比较：
//
//			   1         2
//m: 01234567890123456789012
//S: ABC ABCDAB ABCDABCDABDE
//W:            ABCDABD
//i:            0123456
//这时，S[17](='C')不与W[6]相同，但是亦出现了两次"AB"，我们采取一贯的作法，令m = 15和i = 2，继续搜索。
//
//			   1         2
//m: 01234567890123456789012
//S: ABC ABCDAB ABCDABCDABDE
//W:                ABCDABD
//i:                0123456
//我们找到完全匹配的字符串了，其起始位置于S[15]的地方。
func KMP(s string, p string) int {
	l := len(p)
	switch {
	case l == 0:
		return 0
	case l == 1:
		return IndexByte(s, p[0])
	case l == len(s):
		if s == p {
			return 0
		}
		return -1
	case l > len(s):
		return -1
	default:
	}
	// 部分匹配表
	next := getNext(p)
	fmt.Println("next", next)
	// 开始匹配
	i, j := 0, 0 //主串索引,模式串索引
	for {
		if j >= len(p) {
			return i - len(p)
		}
		if i >= len(s) {
			return -1
		}
		if s[i] != p[j] {
			if j == 0 {
				i++
			}
			j = next[j]
		} else {
			i++
			j++
		}
		fmt.Println(i, j)
	}
}
