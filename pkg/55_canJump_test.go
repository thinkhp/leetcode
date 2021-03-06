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
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		if j := i + nums[i]; j > max {
			max = j
		}
	}

	return true
}

// DP
// dp[i] = dp[i+1] || dp[i+2] || ...... || dp[i+nums[i]]
func canJump2(nums []int) bool {
	var dp = make(map[int]bool)
	l := len(nums)
	dp[l-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		n := nums[i]
		for j := 1; j <= n && i+j < l; j++ {
			dp[i] = dp[i] || dp[i+j]
		}
	}
	fmt.Println(dp)
	return dp[0]
}

// 递归
func canJump1(nums []int) bool {
	return jump(make(map[int]bool), nums, 0)
}

func jump(marks map[int]bool, nums []int, i int) (ok bool) {
	if ok, exist := marks[i]; exist {
		return ok
	}
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
		ok = ok || jump(marks, nums, i+j)
		if ok {
			return
		}
	}
	marks[i] = ok

	return
}
