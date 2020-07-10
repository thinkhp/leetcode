package pkg

import (
	"fmt"
	"testing"
)

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//进阶:
//
//如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func TestMaxSubArray(t *testing.T) {
	//ss := []int{1,2,3}
	ss := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(ss))
}

// 贪心
func maxSubArray(nums []int) int {
	currSum := nums[0] //当前元素位置的最大和
	maxSum := nums[0]  //迄今为止的最大和
	for _, v := range nums[1:] {
		currSum = maxInt(v, currSum+v)
		maxSum = maxInt(maxSum, currSum)
	}

	return maxSum
}

// DP
func maxSubArray1(nums []int) int {
	// 1.状态定义
	// dp[i] 表示前 i 个元素的最大连续子数组的和
	dp := make([]int, len(nums))

	// 2. 状态初始化，数组中第一个元素的最大和就是第一个元素值
	dp[0] = nums[0]
	max := nums[0]

	// 3.状态转移
	// 转移方程：dp[i] = max(dp[i - 1]+nums[i], nums[i])
	// dp 当前元素的值等于前一个元素值和 0 的最大值再加上 nums[i]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxInt(dp[i-1], 0)+nums[i]
		max = maxInt(max, dp[i])
	}

	return max
}


// 0:nums[0]
// 1:max(nums[1], nums[0]+nums[1])
// 2:max()