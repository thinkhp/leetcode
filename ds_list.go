package leetcode

import (
	"strconv"
	"strings"
)

type List struct {
	head *ListNode
	tail *ListNode
	len  int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	node := l
	res := make([]string, 0)
	for {
		if node == nil {
			res = append(res, "nil")
			break
		}
		res = append(res, strconv.Itoa(node.Val))
		node = node.Next
	}
	return strings.Join(res, " ->")
}

func MakeListNode(ss []int) *ListNode {
	root := new(ListNode)
	node := root
	for i, v := range ss {
		node.Val = v
		if i != len(ss)-1 {
			next := new(ListNode)
			node.Next = next
			node = next
		}
	}
	return root
}
