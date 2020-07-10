package pkg

import (
	"fmt"
	"testing"
)

//
//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
//
//比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//
//L   C   I   R
//E T O E S I I G
//E   D   H   N
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//
//请你实现这个将字符串进行指定行数变换的函数：
//
//string convert(string s, int numRows);
//示例 1:
//
//输入: s = "LEETCODEISHIRING", numRows = 3
//输出: "LCIRETOESIIGEDHN"
//示例 2:
//
//输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zigzag-conversion
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func TestConvert(t *testing.T) {
	ss := []struct {
		s string
		n int
	}{
		{"convert", 2},
		{"LEETCODEISHIRING", 3},
		{"LEETCODEISHIRING", 4},
	}

	for _, v := range ss {
		fmt.Println(convert(v.s, v.n))
	}
}
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	l := len(s)
	res := make([]byte, 0)
	cycle := numRows*2 - 2
	numField := l/numRows + 1
	fmt.Println(numField)
	for i := 0; i < numRows; i++ {
		for j := 0; j < numField; j++ {
			if n := cycle*j + i; n < l {
				fmt.Println("nmi", n, string(s[n]))
				res = append(res, s[n])
			}
			switch i {
			case 0, numRows - 1:
			default:
				if n := cycle*(j+1) - i; n < l {
					fmt.Println("mid", n, string(s[n]))
					res = append(res, s[n])
				}
			}
		}
	}

	return string(res)
}

// 二维数组
func convert1(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make([][]byte, 0)
	for i := 0; i < numRows; i++ {
		r := make([]byte, 0)
		res = append(res, r)
	}
	//
	for i, v := range s {
		ri := i % (numRows*2 - 2) // row index
		if ri >= numRows {
			ri = 2*numRows - ri - 2
		}
		res[ri] = append(res[ri], byte(v))
	}
	ss := ""
	for _, v := range res {
		fmt.Println(string(v))
		ss += string(v)
	}
	return ss
}
