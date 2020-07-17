package pkg

import (
	"container/list"
	"fmt"
	"testing"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false
//示例 5:
//
//输入: "{[]}"
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func Test20IsValid(t *testing.T) {
	ss := []string{
		"(",
		")",
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}
	for _, v := range ss {
		t.Run("str:"+v, func(t *testing.T) {
			fmt.Println(isValid(v))
		})

	}
}

func isValid(s string) bool {
	iList := new(stack)
	iList.Init()
	for _, v := range []byte(s) {
		switch v {
		case '(', '{', '[':
			iList.Push(stackValue(v))
			fmt.Println("push", string(v))
			//fmt.Println(string(iList.Front().Value.(uint8)))
		case ')', '}', ']':
			t, err := iList.Pop()
			if err != nil {
				return false
			}
			x := byte(t)
			fmt.Println("pop", string(x), "===", string(v))
			//fmt.Println(string(iList.Front().Value.(uint8)), v)
			switch v {
			case ')':
				if '(' != x {
					return false
				}
			case '}':
				if '{' != x {
					return false
				}
			case ']':
				if '[' != x {
					return false
				}
			default:
				return false

			}
		}
	}
	return iList.Len() == 0
}

func isValid1(s string) bool {
	iList := list.New()
	iList.Init()
	for _, v := range []byte(s) {
		switch v {
		case '(', '{', '[':
			iList.PushFront(v)
			fmt.Println("push", string(v))
			//fmt.Println(string(iList.Front().Value.(uint8)))
		case ')', '}', ']':
			ele := iList.Front()
			iList.Remove(ele)
			x := ele.Value.(byte)
			fmt.Println("pop", string(x), "===", string(v))
			//fmt.Println(string(iList.Front().Value.(uint8)), v)
			switch v {
			case ')':
				if '(' != x {
					return false
				}
			case '}':
				if '{' != x {
					return false
				}
			case ']':
				if '[' != x {
					return false
				}

			}
		}
	}
	return true
}
