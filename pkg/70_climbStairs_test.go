package pkg

import (
	"fmt"
	"testing"
)

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。

func TestClimbStairs(t *testing.T)  {
	fmt.Println(climbStairs(10))
}

// 解法 3
//
//● 全是一步，10个一步。1种。
//● 1个两步，8个一步。相当于从9个坑里放1个两步。C(9,1) = 9种。
//● 2个两步，6个一步。相当于从8个坑里放2个两步。C(8,2) = 28种。
//● 3个两步，4个一步。相当于从7个坑里放3个两步。C(7,3) = 35种。
//● 4个两步，2个一步。相当于从6个坑里放4个两步。C(6,4) = 15种。
//● 5个两步。C(5,5) = 1种。
//
//所以一共有：1 + 9 + 28 + 35 + 15 + 1 = 89种。

// f(n) = f(n-1)+f(n-2)
func climbStairs(n int) int {
	//1: 1
	//2: 2
	//3: f(2) + f(1) = 3
	//4: f(3) + f(2) = 3 + 2 = 5
	//5: f(4) + f(3) = 5 + 3 = 8
	//n: f(n-1) + f(n-2)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a, b := 1, 2
	for i := 3; i < n; i++ {
		a, b = b, a+b
	}
	return a + b
}

// 递归+备忘录
var mark = make(map[int]int)
func climbStairs1(n int) int {
	if v, ok := mark[n]; ok {
		return v
	}
	if n == 0 || n == 1 {
		return 1
	}
	a := climbStairs1(n-1)
	mark[n-1] = a
	b := climbStairs1(n-2)
	mark[n-2] = b
	return a+b
}

func ot(sum int, nums []int) [][]int {
	if sum == 0 || sum == 1 {
		nums = append(nums, sum)

		temp := make([]int, len(nums))
		copy(temp, nums)

		fmt.Println(temp)

		return append([][]int{}, temp)
	}
	return append(ot(sum-1, append(nums, 1)), ot(sum-2, append(nums, 2))...)

}
