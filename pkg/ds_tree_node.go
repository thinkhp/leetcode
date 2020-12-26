package pkg

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

///**
// * Definition for a binary tree node.
// * type TreeNode struct {
// *     Val int
// *     Left *TreeNode
// *     Right *TreeNode
// * }
// */
//func levelOrder(root *TreeNode) [][]int {
//	var result [][]int
//
//	if root == nil {
//		return result
//	}
//
//	//初始化一个队列
//	list := list.New()
//	//从头部插入root
//	list.PushFront(root)
//
//	//开始层次遍历，即广度优先遍历
//	for list.Len() > 0 {
//		var currentLevel []int
//		//取本层的节点数
//		curentLenth := list.Len()
//		for i := 0; i < curentLenth; i++ {
//			//从尾部移除，Remove返回值为接口类型，需指定为TreeNode
//			node := list.Remove(list.Back()).(*TreeNode)
//			currentLevel = append(currentLevel, node.Val)
//			if node.Left != nil {
//				list.PushFront(node.Left)
//			}
//			if node.Right != nil {
//				list.PushFront(node.Right)
//			}
//		}
//		//当前层结束
//		result = append(result, currentLevel)
//	}
//
//	return result
//}
//
//                              E1
//              D1                              D1
//      C1              C2              C1              C2
//  B1      B2      B3      B4      B1      B2      B3      B4
//A1  A2  A3  A4  A5  A6  A7  A8  A1  A2  A3  A4  A5  A6  A7  A8
func (t *TreeNode) String() string {
	res := ""
	// 深度遍历,当做完全二叉树打印
	// 获取深度以及字符字段宽度
	treeDepth := 0
	charLen := 0

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		// 变更最大字符长度
		l := len(strconv.Itoa(node.Val))
		if charLen < l {
			charLen = l
		}
		// 修改深度
		depth++
		if treeDepth < depth {
			treeDepth = depth
		}
		dfs(node.Left, depth)
		dfs(node.Right, depth)
	}

	dfs(t, 0)
	fmt.Printf("tree depth: %v, charLen: %v\n", treeDepth, charLen)
	table := make([]string, treeDepth)

	// printTree
	var pt func(node *TreeNode, depthIndex int)
	// 深度遍历
	pt = func(node *TreeNode, depthIndex int) {
		// 2 ^ (l - d - 1) - 1
		gap := (int(math.Exp2(float64(treeDepth-depthIndex-1))) - 1) * charLen
		if node == nil {
			x := gap*2 + 2*charLen
			//if depthIndex < treeDepth {
			//	table[depthIndex] += sprintSpace(x)
			//}

			for i := depthIndex; i < treeDepth; i++ {
				table[i] += sprintSpace(x)
			}
			return
		}
		val := strconv.Itoa(node.Val)
		v := sprintSpace(gap) + val + sprintSpace(charLen-len(val)) + sprintSpace(gap) + sprintSpace(charLen)
		table[depthIndex] += v

		pt(node.Left, depthIndex+1)
		pt(node.Right, depthIndex+1)
	}

	pt(t, 0)
	for _, v := range table {
		res += v + "\n"
	}
	return res
}

func sprintSpace(n int) string {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = ' '
	}
	return string(bs)
}

// BFS
func GetTreeNodeFromSlice(s []string) *TreeNode {
	var getValue = func(str string) *TreeNode {
		if str == "null" {
			return nil
		} else {
			num, _ := strconv.Atoi(str)
			return &TreeNode{num, nil, nil}
		}
	}
	root := getValue(s[0])
	useList := []*TreeNode{root}
	nextUseList := make([]*TreeNode, 0)
	useIndex := 0
	use := useList[0]

	for i := 1; i < len(s); i += 2 {
		//fmt.Println(s[i])
		use.Left = getValue(s[i])
		if use.Left != nil {
			nextUseList = append(nextUseList, use.Left)
		}
		if i+1 >= len(s) {
			break
		} else {
			use.Right = getValue(s[i+1])
		}

		if use.Right != nil {
			nextUseList = append(nextUseList, use.Right)
		}
		useIndex++
		if useIndex >= len(useList) {
			useList = nextUseList
			nextUseList = make([]*TreeNode, 0)
			useIndex = 0
			use = useList[0]
			continue
		}
		use = useList[useIndex]

	}
	return root
}

//
//1. 首先将根节点放入stack中。
//2. 从stack中取出第一个节点，并检验它是否为目标。
//	 如果找到目标，则结束搜寻并回传结果。
//	 否则将它某一个尚未检验过的直接子节点加入stack中。
//3. 重复步骤2。
//4. 如果不存在未检测过的直接子节点。 将上一级节点加入stack中。 重复步骤2。
//5. 重复步骤4。
//6. 若stack为空，表示整张图都检查过了——亦即图中没有欲搜寻的目标。结束搜寻并回传“找不到目标”。
func DFS(n *TreeNode, stack []int) [][]int {
	if n == nil {
		return [][]int{}
	}
	if n.Left == nil && n.Right == nil {
		// must copy
		stack = append(stack, n.Val)
		pathN := make([]int, len(stack))
		copy(pathN, stack)
		return append([][]int{}, pathN)
	}
	return append(DFS(n.Left, append(stack, n.Val)), DFS(n.Right, append(stack, n.Val))...)
}
