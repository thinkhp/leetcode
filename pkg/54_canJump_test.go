package pkg

import (
	"fmt"
	"testing"
)

//给定一个非负整数数组，你最初位于数组的第一个位置。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//判断你是否能够到达最后一个位置。
//
//示例 1:
//
//输入: [2,3,1,1,4]
//输出: true
//解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
//示例 2:
//
//输入: [3,2,1,0,4]
//输出: false
//解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。

func Test_canJump(t *testing.T) {
	ss := [][]int{
		{2, 3, 1, 1, 4},
		{2, 3, 1, 1, 0},
		{3, 2, 1, 0, 4},
	}

	for _, v := range ss {
		fmt.Println(v, canJump(v))
	}
}

func canJump(nums []int) bool {
	max := 0
	for i := max; i < len(nums); i++ {
		if i > max {
			return false
		}
		if j := i + nums[i]; j > max {
			max = j
		}
	}

	return true
}

// 遍历法
func canJump1(nums []int) bool {
	return jump(nums, 0)
}

func jump(nums []int, i int) (ok bool) {
	v := nums[i]
	fmt.Printf("nums:%v in:%v canJump:%v\n", nums, i, v)
	// 结束
	if i+v >= len(nums)-1 {
		//fmt.Println("ok")
		return true
	}
	if v == 0 {
		return false
	}
	// 约束
	for j := 1; j <= v; j++ {
		ok = ok || jump(nums, i+j)
		if ok {
			return
		}
	}

	return
}
