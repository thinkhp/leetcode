package pkg

import (
	"fmt"
	"testing"
)

func TestStringMultiply(t *testing.T) {
	fmt.Println(reverseBytes([]byte{1, 2, 3}))
	fmt.Println(multiplyString("222", "999"))
	fmt.Println(multiplyString("0", "0"))
}

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//示例 1:
//输入: num1 = "2", num2 = "3"
//输出: "6"
//
//示例 2:
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//
//说明：
//num1 和 num2 的长度小于110。
//num1 和 num2 只包含数字 0-9。
//num1 和 num2 均不以零开头，除非是数字 0 本身。
//不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
func multiplyString(num1 string, num2 string) string {
	return string(multiplyByte([]byte(num1), []byte(num2)))
}

func reverseBytes(bs []byte) []byte {
	l := len(bs)
	res := make([]byte, l)
	for i, v := range bs {
		res[l-i-1] = v
	}
	return res
}
func multiplyByteFast(bs1 []byte, bs2 []byte) []byte {
	bs := make([]byte, 0)
	//for i1, v1 := range bs1 {
	//	for i2, v2 := range bs2 {
	//		sum := (v1 - '0') * (v2 - '0')
	//		//a, b := sum/10, sum%10
	//		for j := 0; sum != 0; j++ {
	//			old := bs[i1+i2+j]
	//			bs[i1+i2+j] = (old+sum)%10
	//			sum = (old+sum)/10
	//		}
	//	}
	//}
	//
	//
	return bs
}
func multiplyByte(bs1 []byte, bs2 []byte) []byte {
	if len(bs1) == 1 && bs1[0] == '0' {
		return []byte{'0'}
	}
	if len(bs2) == 1 && bs2[0] == '0' {
		return []byte{'0'}
	}
	if len(bs1) == 1 && bs1[0] == '1' {
		return bs2
	}
	if len(bs2) == 1 && bs2[0] == '1' {
		return bs1
	}

	bs1 = reverseBytes(bs1)
	bs2 = reverseBytes(bs2)
	l1 := len(bs1)
	l2 := len(bs2)
	bs := make([]byte, l1+l2+1)

	for i1, v1 := range bs1 {
		for i2, v2 := range bs2 {
			sum := (v1 - '0') * (v2 - '0')
			//a, b := sum/10, sum%10
			for j := 0; sum != 0; j++ {
				old := bs[i1+i2+j]
				bs[i1+i2+j] = (old + sum) % 10
				sum = (old + sum) / 10
			}
		}
	}

	fmt.Println("un reverse", bs)
	bs = reverseBytes(bs)
	fmt.Println("do reverse", bs)
	for i, v := range bs {
		if v != 0 {
			bs = bs[i:]
			break
		}
	}
	if bs[0] == 0 {
		return []byte{'0'}
	}
	fmt.Println("out zero", bs)
	for i, v := range bs {
		bs[i] = v + '0'
	}

	fmt.Println("trans string", bs)
	return bs
}

func StringToSliceBit(s string) []byte {
	bs := []byte(s)
	fmt.Println(bs, bs[0]*bs[1])
	return bs
}
