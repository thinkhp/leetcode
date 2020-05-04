package leetcode

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
