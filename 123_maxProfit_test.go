package leetcode

import (
	"fmt"
	"testing"
)
// TODO
//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
//注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 数组
func TestMaxProfit_3(t *testing.T) {
	ss := [][]int{
		{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{7, 6, 4, 3, 1},
		{3, 3, 5, 0, 0, 3, 1, 4},
		{1, 2, 3, 4, 5},
		{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
	}
	for _, v := range ss {
		fmt.Println(v, maxProfit_3(v))
	}
}
func maxProfit_3(prices []int) int {
	// 先求出 Am-An 的二维数组
	dp := make([][]int, len(prices))
	for i, _ := range prices {
		dp[i] = make([]int, len(prices))
	}
	for i := 1; i < len(prices); i++ {
		pI := prices[i]
		dp[i-1][i] = pI - prices[i-1]
		for j := i - 2; j >= 0; j-- {
			dp[j][i] = dp[i-1][i] + dp[j][i-1]
		}
	}

	// 交易发生在第 j 天到第 i 天的情况下,单笔收益的最大值
	dpMax := make([][]int, len(prices))
	for i, _ := range prices {
		dpMax[i] = make([]int, len(prices))
	}
	//for i := 0; i < len(prices); i++ {
	//	for j := i-2; j >= 0; j-- {
	//		dpMax[j][i] =
	//	}
	//}

	//var max = func(start, end int) int {
	//	for i := 0; i < len(); i++ {
	//
	//	}
	//}

	// 分别考虑 1----k, k+1----n下
	for k := 1; k < len(prices); k++ {
		// max(1-k)
	}
	if len(prices) <= 1 {
		return 0
	}
	maxProfit := 0

	return maxProfit
}
