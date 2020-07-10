package pkg

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestGetTreeNodeFromSlice(t *testing.T) {
	s := "-260,-202,-903,-980,-570,-858,218,764,-300,205,null,-35,null,null,-204,950,-769,258,-652,614,-584,76,817,-192,null,null,-114,880,null,-200,71,671,344,801,193,-18,876,-920,-730,222,679,null,-680,null,null,null,-859,744,-261,692,null,-341,-163,null,null,482,-979,205,null,146,165,801,100,-656,714,-629,995,474,307,-581,-150,-941,null,null,null,-937,-69,-23,82,null,-139,-591,null,-453,-861,-370,null,null,null,216,233,null,430,null,5,-110,null,null,-660,624,-510,-588,null,null,381,null,368,559,null,521,-301,null,522,379,null,null,null,null,456,519,null,null,482,349,null,null,19,null,null,288,-811,null,-372,null,null,-536,null,-404,-457,-740,860,null,null,-636,null,null,342,-874,-462,-504,781,855,-392,null,null,null,406,null,-758,541,null,-947,null,null,null,null,null,-964,null,600,-45,null,null,null,null,null,null,null,null,null,-194,null,null,null,-802,null,null,null,-3,null,-792,672,643,null,14,null,null,489,457,null,null,null,null,412,null,558,null,null,null,null,-846,158,-146,null,null,-76,-650,null,-782,null,-127,null,-678,null,null,null,null,null,null,-464,-426,null,-366,null,null,null,null,null,81,-607,716,null,null,-213,null,379,null,null,null,null,644,445,null,null,-419,-845,-720,null,null,-915,null,null,null,null,null,null,-686,594,-243,null,496,null,907,null,null,null,null,null,null,579,873,702,null,null,null,-834,null,null,null,null,null,-300,-214,-466,null,null,972,null,null,null,814,null,-940,null,763,null,null,null,null,-449,-844,null,null,null,null,-47"
	root := GetTreeNodeFromSlice(strings.Split(s, ","))
	fmt.Println(DFS(root, []int{}))
}

func TestTreeNodeHasPathSum(t *testing.T) {
	fmt.Println(strings.Compare("bb", "abc"))
	fmt.Println(strings.Compare("ab", "abc"))
	//s := "-10,9,20,null,null,15,7"
	//s := "-1,-2"
	//s := "5,4,8,11,null,13,4,7,2,null,null,null,1"
	//s := "2,2,1,null,1,0,null,0"
	s := "3,9,20,null,null,15,7"
	s = strings.ReplaceAll(s, " ", "")
	root := GetTreeNodeFromSlice(strings.Split(s, ","))
	fmt.Println(treePathAndSum(root, []int{}))
	//fmt.Println(GetAllPathFromRoot(root, []int{}))
	//fmt.Println(pathSumAndMax(root))
	//fmt.Println(allTreeNodePath(root, []int{}, 0))
	//fmt.Println(treeNodePathSum(root, []int{}, -243))
}

func treePathAndSum(node *TreeNode, path []int) string {
	if node == nil {
		return ""
	}
	path = append(path, node.Val)
	if node.Right == nil && node.Left == nil {
		str := getStr(path)
		return str
	}
	strL := treePathAndSum(node.Left, path)
	strR := treePathAndSum(node.Right, path)
	if strL == "" {
		return strR
	}
	if strR == "" {
		return strL
	}
	if strings.Compare(strL, strR) == -1 {
		return strL
	}
	return strR
}

func getStr(path []int) string {
	str := ""
	for i := len(path) - 1; i >= 0; i-- {
		str += string(path[i] + 'a')
	}
	return str
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func getMax(node, left, right int) {

}

// 计算节点的最大贡献值以及整个 tree 下最大的值
// p 父节点;l 左节点;r 右节点
func pathSumAndMax(p *TreeNode) (int, int) {
	// 易错:在节点为空的情况下,返回的最小值不应该为 0
	// 用什么样的代码可以替代 math.MinInt32,保证空节点的最大值不被计算
	if p == nil {
		return 0, math.MinInt32
	}
	v := p.Val
	if p.Right == nil && p.Left == nil {
		//fmt.Println("v:",v)
		return v, v
	}
	l, ml := pathSumAndMax(p.Left)
	f, mr := pathSumAndMax(p.Right)
	mp := maxInt(maxInt(l+v, f+v), v)
	// 易错点
	// 节点的贡献值只需要考虑 (父节点+某个子节点)/父节点
	// tree的最大值要考虑 (父节点+某个子节点)/父节点/(父节点+全部子节点)
	return mp, maxInt(maxInt(maxInt(ml, mr), mp), v+l+f)
}

type treeNodePath struct {
	list []int
	sum  int
}

// 输出所有的
var target = -243

func allTreeNodePath(n *TreeNode, path []int, sum int) []treeNodePath {
	if n == nil {
		return []treeNodePath{}
	}
	sum += n.Val
	if n.Left == nil && n.Right == nil {
		if sum == target {
			path = append(path, n.Val)
			pathN := make([]int, len(path))
			copy(pathN, path)
			fmt.Println(sum, pathN)
			return append([]treeNodePath{}, treeNodePath{list: pathN, sum: sum})
		}
		return []treeNodePath{}
	}
	return append(allTreeNodePath(n.Left, append(path, n.Val), sum), allTreeNodePath(n.Right, append(path, n.Val), sum)...)
}

func treeNodePathSum(n *TreeNode, path []int, sum int) [][]int {
	if n == nil {
		return [][]int{}
	}
	sum -= n.Val
	if n.Left == nil && n.Right == nil {
		if sum == 0 {
			return append([][]int{}, append(path, n.Val))
		}
		return [][]int{}
	}
	return append(treeNodePathSum(n.Left, append(path, n.Val), sum), treeNodePathSum(n.Right, append(path, n.Val), sum)...)
}

//func pathSum(root *TreeNode, sum int) [][]int {
//
//}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
