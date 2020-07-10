package pkg

import (
	"fmt"
	"testing"
)

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//说明：
//
//给定的 n 保证是有效的。
//
//进阶：
//
//你能尝试使用一趟扫描实现吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func TestRemoveNthFromEnd(t *testing.T) {
	type mock struct {
		nums int
		nTH int
	}

	ss := []mock{
		{123456, 1},
		{123456, 6},
		{123456, 3},
		{12, 1},
		{12, 2},
		{123, 1},
		{123, 3},
	}
	for _, v := range ss {
		head := reverseList(genNode(v.nums))
		//fmt.Println(head)
		node := removeNthFromEnd(head, v.nTH)
		fmt.Println(v, node.String())
	}
}



func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeQuick := head
	nodeSlow := head

	// 使得 (nodeQuick,head) 开区间的节点个数为 n
	for i := 0; i < n+1; i++ {
		nodeQuick = nodeQuick.Next
		if nodeQuick == nil && i==n-1 { // 因为在 n == len(list) 的情况下,会越界
			// 要删除的节点为第一个节点
			return head.Next
		}
	}
	//fmt.Println(head, nodeQuick)
	for nodeQuick != nil {
		nodeQuick = nodeQuick.Next
		nodeSlow = nodeSlow.Next
	}

	// 要删除的节点为最后一个节点, 或者要删除的节点为中间的节点
	nodeSlow.Next = nodeSlow.Next.Next
	return head
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	i := 0 //nodeQuick已经遍历过的节点
	nodeQuick := head
	nodeSlow := head
	for {
		nodeQuick = nodeQuick.Next
		i++
		if i > n+1 {
			nodeSlow = nodeSlow.Next
		}
		if nodeQuick == nil {
			break
		}
	}
	if i-n==1 || nodeSlow != head {
		nodeSlow.Next = nodeSlow.Next.Next
		return head
	} else {
		return head.Next
	}
}