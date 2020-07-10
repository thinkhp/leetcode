package pkg

import (
	"fmt"
	"testing"
)

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

func TestAddTwoNumbers(t *testing.T) {
	rootNode := genNode(18)
	fmt.Println(rootNode.Val, rootNode.Next.Val)
	fmt.Println(genNum(rootNode))
	//fmt.Println(math.MaxInt64)
	fmt.Println(genNum(addTwoNumbers(genNode(123), genNode(897))))

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	l := &ListNode{}
	node := l

	var get = func(l *ListNode) int {
		if l == nil {
			return 0
		} else {
			return l.Val
		}
	}
	var set = func(l *ListNode) *ListNode {
		if l == nil {
			return nil
		} else {
			return l.Next
		}
	}
	for {
		n1 := get(l1)
		n2 := get(l2)
		s := n1 + n2 + add
		add = s / 10
		node.Val = s % 10

		l1 = set(l1)
		l2 = set(l2)
		if (l1 == nil || l2 == nil) && add == 0 {
			if l1 != nil {
				node.Next = l1
				break
			}
			if l2 != nil {
				node.Next = l2
				break
			}
			break
		}
		node.Next = &ListNode{}
		node = node.Next
	}
	return l
}

func genNum(l *ListNode) int {
	sum := 0
	i := 1
	for {
		sum += l.Val * i
		if l.Next == nil {
			break
		}
		l = l.Next
		i *= 10
	}
	return sum
}

// 12345 转 1->2->3->4->5
func genNode(sum int) *ListNode {
	root := new(ListNode)
	node := root
	for {
		s := sum / 10
		y := sum % 10
		node.Val = y
		if s == 0 {
			return root
		}
		next := new(ListNode)
		node.Next = next

		node = next
		sum = s
	}
}
