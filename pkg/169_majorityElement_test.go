package pkg

import (
	"fmt"
	"testing"
)

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//示例 1:
//
//输入: [3,2,3]
//输出: 3
//示例 2:
//
//输入: [2,2,1,1,1,2,2]
//输出: 2
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/majority-element
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func TestMajorityElement(t *testing.T) {
	ss := [][]int{
		{3, 2, 3},
		{2, 2, 1, 1, 1, 2, 2},
	}

	for _, v := range ss {
		fmt.Println(v, majorityElement(v))
	}
}

// 摩尔投票法
//Boyer-Moore 投票算法
// ???将任意不等的数字两两抵消,留下来的数字必定为众数???
// 证明:
//设真正的众数为 maj，且在遍历数组 nums 的某一任意时刻的候选众数为 c。
//此时我们遍历数组 nums 中的下一个元素 x。
//如果 x == maj，那么无论 c 是否等于 maj，计数器的值都会加 1（仔细思考一下，为什么？）；
//如果 x != maj，那么计数器在 x == c 时会加 1，在 x != c 时会减 1。
//由于 maj 是数组 nums 的众数，出现的次数超过一半，因此在最后计数器 count 的值一定大于 0（因为在超过一半的遍历中，计数器的值都加 1）。
//同理，如果我们以非众数的任一整数作为基准，可以用同样的方法得到，在最后计数器 count 的值一定小于 0，但 count 在我们的算法步骤中显然不会小于 0，因此最终的 candidate 一定就是 maj。
// 应用原理:如果数 a 是数组 nums 的众数，如果我们将 nums 分成两部分，那么 a 必定是至少一部分的众数。
func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

// hash
func majorityElement1(nums []int) int {
	l := len(nums)
	numCount := make(map[int]int, l)

	for _, v := range nums {
		c, _ := numCount[v]
		c++
		numCount[v] = c
		if c > l/2 {
			return v
		}
	}

	return 0
}
