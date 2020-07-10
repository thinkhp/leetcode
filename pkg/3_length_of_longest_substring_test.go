package pkg

import (
	"fmt"
	"testing"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func TestLengthOfLongestSubstring(t *testing.T) {
	var s string
	s = "abcdeafghaihkkkkkk"
	//s = "abcbba"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))
}

// 滑动窗口, 使用 byte 桶数组代替 hash 加快查询速度
// 给定一个字符串，请你找出其中不含有重复字符的最长子串的长度
func lengthOfLongestSubstring(ss string) int {
	bs := []byte(ss)
	m := [256]int{}
	for i, _ := range m {
		m[i] = -1
	}
	max := 0
	l := len(bs)
	if l == 0 {
		return max
	}

	s, e:= 0, 1
	for ;s < l && e < l ; {
		v := ss[e]
		if i := m[v]; i != -1 { //在 [s,e] 出现重复元素
			if max < e-s {
				max = e-s
			}
			s = m[v]+1
		}
		m[v] = e
		e++
	}
	if max < e-s {
		max = e-s
	}
	return max
}

// 思路:遍历,使用 map[byte]bool 管理目前子串的各元素在 []byte(str) 的索引
// 如果在 map 找到重复的元素,确定索引 j,子串变为 str[j+1:i+1],重新设置 map
func lengthOfLongestSubstring1(s string) int {
	bs := []byte(s)
	m := make(map[byte]int)
	max := 0
	l := 0
	for i, v := range bs {
		if j, ok := m[v]; ok {
			// save
			if max < l {
				max = l
			}
			// init
			m[v] = i
			l = i - j
			sc := bs[j+1 : i+1]
			//fmt.Println("left", j, i, i-j,string(sc))
			m = make(map[byte]int)
			for in, vn := range sc {
				m[vn] = in + j + 1
			}
			continue
		}
		m[v] = i
		l++
	}
	if max < l {
		max = l
	}

	return max
}

// 双指针,滑动窗口
func lengthOfLongestSubstring2(s string) int {
	bs := []byte(s)
	m := make(map[byte]int) // 管理对应元素的位置
	max := 0

	for i, j := 0, 0; j < len(bs); j++ {
		v := bs[j]
		if r, ok := m[v]; ok {
			fmt.Println("r", r)
			if i < r+1 {
				i = r + 1
			}
		}
		fmt.Println("[max]:", string(bs[i:j+1]))
		m[v] = j
		if max < j-i+1 {
			max = j - i + 1
		}
	}
	return max
}
