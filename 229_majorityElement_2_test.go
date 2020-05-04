package leetcode

import (
	"fmt"
	"testing"
)

//给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
//说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。
//
//示例 1:
//
//输入: [3,2,3]
//输出: [3]
//示例 2:
//
//输入: [1,1,1,3,3,2,2,2]
//输出: [1,2]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/majority-element-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func TestMajorityElement_2(t *testing.T) {
	ss := [][]int{
		{3, 2, 3},
		{2, 2, 1, 1, 1, 2, 2},
		{8, 8, 7, 7, 7},
	}

	for _, v := range ss {
		fmt.Println(v, majorityElement_2(v))
	}
}

//摩尔投票法
//|
//|		|
//|		|		|
//|		|		|
//|		|		|
//|		|		|		|
//|		|		|		|
//|		|		|		|
//|		|		|		|
//|_______|_______|_______|______
//		A		B		C
// 其中第一列全是A, 第二列全是B, 第三列为其他数, 且第三列即使所有数都一样, 也不会超过[n/3], 更何况可能存在不一样的数;
//
//如果我每次都拿走3个不一样的数, 那么最后剩下的, 一定是A, B.
func majorityElement_2(nums []int) []int {
	a := 0
	b := 0 //a,b不可以相等
	countA := 0
	countB := 0

	// 抵消
	for _, v := range nums {
		if countA == 0 {
			if v == b {
				countB++
				continue
			} else {
				a = v
				countA++
				continue
			}
		}
		if countB == 0 {
			if v == a {
				countA++
				continue
			} else {
				b = v
				countB++
				continue
			}

		}
		switch v {
		case a:
			countA++
		case b:
			countB++
		default:
			countB--
			countA--
		}

	}
	fmt.Println(a, b)

	// 遍历
	countA = 0
	countB = 0
	for _, v := range nums {
		switch v {
		case a:
			countA++
		case b:
			countB++
		default:
		}
	}
	res := make([]int, 0)
	if countA > len(nums)/3 {
		res = append(res, a)
	}
	if countB > len(nums)/3 {
		res = append(res, b)
	}

	return res
}

// hash
func majorityElement_2_1(nums []int) []int {
	l := len(nums)
	numCount := make(map[int]int, l)
	ret := make([]int, 0)

	for _, v := range nums {
		c, _ := numCount[v]
		c++
		numCount[v] = c
	}

	for num, count := range numCount {
		if count > l/3 {
			ret = append(ret, num)
		}
	}

	return ret
}
