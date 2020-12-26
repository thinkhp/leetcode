package pkg

import "testing"

//给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
//
//
//
//示例 1：
//
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出：true
//示例 2：
//
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出：false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/interleaving-string
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func TestIsInterleave(t *testing.T) {
	ss := []struct {
		s1 string
		s2 string
		s3 string
	}{
		{"", "", ""},
	}

	for _, v := range ss {
		t.Logf("%+v result: %v\n", v, isInterleave(v.s1, v.s2, v.s3))
	}
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	return true
}
