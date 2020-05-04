package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	fmt.Println(reverseList2(MakeListNode([]int{1, 2, 3, 4, 5})))
}

// 空间复杂度为 2*N
// 创建一个新的链表,维持四个变量 node1,node1.next node2,node2.previous
func reverseList1(head *ListNode) *ListNode {
	node1 := head
	node2 := new(ListNode)
	for {
		node2.Val = node1.Val
		if node1.Next == nil {
			break
		}
		pre := new(ListNode)
		pre.Next = node2

		node1 = node1.Next
		node2 = pre
	}
	return node2
}

// 空间复杂度为 N
// 迭代,将元素依次置于链的表头;将之前的表头的 next 赋值 nil
func reverseList2(head *ListNode) *ListNode {
	node := head // 正在遍历的 node
	for node != nil {
		next := node.Next
		if node == head {
			node.Next = nil
		} else {
			node.Next = head
		}
		head = node
		node = next
	}
	return head

	// corner
	if head == nil || head.Next == nil {
		return head
	}
	// revise
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

// 思路同 2,递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oldTail := reverseList(head.Next)
	// head.Next 与 next 交换
	head.Next.Next = head
	head.Next = nil
	return oldTail
}
