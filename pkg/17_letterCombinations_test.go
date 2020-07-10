package pkg

import (
	"fmt"
	"testing"
)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//示例:
//
//输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//说明:
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func TestLetterCombinations(t *testing.T) {
	ss := []string{
		"2",
		"23",
		"23456789",
	}

	for _, v := range ss {
		fmt.Println(v, letterCombinations(v))
	}
}

var kv = [][]byte{
	[]byte{'a', 'b', 'c'},
	[]byte{'d', 'e', 'f'},
	[]byte{'g', 'h', 'i'},
	[]byte{'j', 'k', 'l'},
	[]byte{'m', 'n', 'o'},
	[]byte{'p', 'q', 'r', 's'},
	[]byte{'t', 'u', 'v'},
	[]byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	for _, digit := range []byte(digits) {
		l := len(result)
		for _, letter := range kv[digit-'2'] {
			if l == 0 {
				result = append(result, string(letter))
			} else {
				for i := 0; i < l; i++ {
					v := append([]byte(result[i]), letter)
					result = append(result, string(v))
				}
			}
		}
		result = result[l:]
	}

	return result
}
