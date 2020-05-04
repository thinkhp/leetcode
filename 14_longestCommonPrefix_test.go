package leetcode

import (
	"fmt"
	"testing"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。
func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{
		"asdfgfdsada",
		"asdfg",
		"asdfg3",
		"",
	}
	fmt.Println(longestCommonPrefix(strs))
}
//  i
//  longest
//j longest
func longestCommonPrefix(strs []string) string {
	str := strs[0]
	i := 0
	T:
	for ; i < len(str); {
		for j := 1; j < len(strs); j++ {
			l := len(strs[j])
			if i > l-1 || strs[j][i] != str[i] {
				break T
			}
		}
		i++
	}
	return str[:i]

}