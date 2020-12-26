package pkg

import (
	"math"
	"testing"
)

//给定一个非空二叉树，返回其最大路径和。
//
//本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
//
//
//示例 1：
//
//输入：[1,2,3]
//
//1
/// \
//2   3
//
//输出：6
//示例 2：
//
//输入：[-10,9,20,null,null,15,7]
//
//   -10
//   / \
//  9  20
//    /  \
//   15   7
//
//输出：42
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func TestMaxPathSum(t *testing.T) {
	ss := [][]string{
		[]string{"10", "9", "-8"},
		[]string{"-3", "-7", "-10"},
		[]string{"-7", "-3", "-10"},
		[]string{"-3"},
		[]string{"1", "2", "3"},
		[]string{"-10", "9", "20", "null", "null", "15", "7"},
	}
	for _, s := range ss {
		tn := GetTreeNodeFromSlice(s)
		t.Log(s, "\n"+tn.String()+"\n", maxPathSum(tn))
	}

}

// 每个节点存在两个值:最大贡献值, 以该节点为中心的最大路径和
// mp 最大路径和
func maxPathSum(root *TreeNode) int {
	var mp = root.Val

	// 最大贡献值
	var maxContributionValue func(n *TreeNode) int
	maxContributionValue = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		{
			v, left, right := n.Val, maxContributionValue(n.Left), maxContributionValue(n.Right)
			mp = maxInt(v+right, v+left, v+left+right, mp)
			return maxInt(v, v+left, v+right)
		}
		{
			// 保证贡献值不小于 0, 可以减少比较个数
			v := n.Val
			left := maxInt(0, maxContributionValue(n.Left))
			right := maxInt(0, maxContributionValue(n.Right))

			mp = maxInt(v+left+right, mp)
			return v + maxInt(left, right)
		}
	}

	maxContributionValue(root)
	return mp
}

var max int

// const INT_MAX = int(^uint(0) >> 1)
// const INT_MIN = ^INT_MAX

func maxPathSum2(root *TreeNode) int {
	max = math.MinInt32
	dfs(root)
	return max
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftmax := dfs(root.Left)
	rightmax := dfs(root.Right)

	sum := leftmax + rightmax + root.Val
	if sum > max {
		max = sum
	}

	contribute := root.Val + maxval(leftmax, rightmax)
	if contribute > 0 {
		return contribute
	}
	return 0
}

func maxval(a, b int) int {
	if a > b {
		return a
	}
	return b
}
