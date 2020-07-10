package pkg

import (
	"fmt"
	"testing"
)

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
//
//注意你不能在买入股票前卖出股票。
//
//示例 1:
//
//输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
//示例 2:
//
//输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func TestMaxProfit_1(t *testing.T) {
	ss := [][]int{
		{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
	}
	for _, v := range ss {
		fmt.Println(v, maxProfit_1_1(v), maxProfit_1_5(v))
	}
}

// 官方谷底/谷峰比较法
func maxProfit_1(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}

	}
	return maxProfit
}

// 我们需要找出给定数组中两个数字之间的最大差值.此外，第二个数字的索引必须大于第一个数字的索引
// 对每一个谷底之后的谷峰进行比较
func maxProfit_1_5(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxProfit := 0
	minPrice := prices[0]
	maxPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if maxPrice < price {
			maxPrice = price
		}
		if minPrice > price {
			// 新的谷底
			profit := maxPrice - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
			minPrice = price
			maxPrice = 0
		}
	}
	profit := maxPrice - minPrice
	if profit > maxProfit {
		maxProfit = profit
	}
	//minPrice := prices[0]
	//
	//for i := 0; i < len(prices); i++ {
	//	if prices[i] < minPrice {
	//		minPrice = prices[i]
	//	} else if prices[i] - minPrice > maxProfit {
	//		maxProfit = prices[i] - minPrice
	//	}
	//
	//}
	return maxProfit
}

// 区间和可以转换成求差的问题;求差问题，也可以转换成区间和的问题
// Am-An = Am-A(m-1)+A(m-1)-A(m-2)+......+A(n+1)-An
func maxProfit_1_3(prices []int) int {
	maxProfit := 0
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 0; i < len(prices); i++ {
		dp[i] = prices[i] - prices[i-1]
	}
	for i, j := 0, len(prices)-1; i < j; {

	}
	return maxProfit
}

// DP
// 暴力法中 An-A1= An-A(n-1) + A(n-1)-A1
//| 5    | X    | X     | X     | X     | X     |
//| 4    | X    | X     | X     | X     | A5-A4 |
//| 3    | X    | X     | X     | A4-A3 | 左+上 |
//| 2    | X    | X     | A3-A2 | 左+上  | 左+上 |
//| 1    | X    | A2-A1 | 左+上  | 左+上  | 左+上 |
//| 0    | 1    | 2     | 3     | 4     | 5     |
func maxProfit_1_2(prices []int) int {
	dp := make([][]int, len(prices))
	for i, _ := range prices {
		dp[i] = make([]int, len(prices))
	}
	maxProfit := -1
	for i := 1; i < len(prices); i++ {
		pI := prices[i]
		dp[i-1][i] = pI - prices[i-1]
		for j := i - 2; j >= 0; j-- {
			dp[j][i] = dp[i-1][i] + dp[j][i-1]
		}
	}
	for i, vv := range dp {
		for j, v := range vv {
			fmt.Printf("A%d-A%d=%d\n", j, i, v)
		}
	}
	return maxProfit
}

// 暴力法
// 假设股票价格为 A1,A2,A3,A4,A5......An
// 求解如果在第 X 天卖出,res 的最大值
// 1:无法交易
// 2:max(A1-A0, 0)
// 3:max(A1-A0, 0, A2-A1)
// 4:max(A1-A0, 0, A2-A1, A3-A2, A3-A1)
// 5:max(A1-A0, 0, A2-A1, A3-A2, A3-A1, A4-A3, A4-A2, A4-A1)
// ......第 x 天的所有可能为 C(x, 2)+1,相较于 第 x-1 天增加 x-1 总结果
// x:max(A1-A0, 0, A2-A1, A3-A2, A3-A1, A4-A3, A4-A2, A4-A1......An-A1)
// 所以时间复杂度为 O(n*n)
func maxProfit_1_1(prices []int) int {
	maxProfit := -1
	for i, pI := range prices {
		for j := 0; j < i; j++ {
			pJ := prices[j]
			profit := pI - pJ
			if maxProfit < profit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
