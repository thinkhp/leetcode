package pkg

import (
	"fmt"
	"testing"
)

//给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//说明：你不能倾斜容器，且 n 的值至少为 2。
//
//示例:
//
//输入: [1,8,6,2,5,4,8,3,7]
//输出: 49
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/container-with-most-water
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func TestMaxArea(t *testing.T) {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(h))
}

// 双指针
// 移动较短的线段short(因为移动较长的线段long,无论下一个线段next 比 long 长还是短,都会导致面积减少或者不变)
// 一共 C(n,2)=n*(n-1)/2 情况,不会缺少任意一种情况
//
//双指针正确性证明
//需要证明:
//1.会不会重现重复或者缺少的情况
//2.会不会遗漏最大值
//
//设：共 m 条线段，
//故 一共存在 C(m,2)=m*(m-1)/2 种情况
//证明 2
//当移动较长线段long时,无论下一个线段next 比 long 长还是短,都会导致面积减少或者不变
//故以较短线段 short 为一端,另外一段不为排除线段(在(long,short)之外的线段)的情况下, (long,short)的面积最大
//证明 1
//
func maxArea(height []int) int {
	max := 0
	l := len(height)
	for i, j := 0, l-1; i < j; {
		yi := height[i]
		yj := height[j]
		s := minInt(yj, yi) * (j - i)
		if s > max {
			max = s
		}
		if yi < yj {
			j--
		} else {
			i++
		}
	}
	return max
}

// 暴力法 n*n
func maxArea1(height []int) int {
	max := 0
	l := len(height)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			x := j - i
			y := minInt(height[i], height[j])
			if x*y > max {
				max = x * y
			}
		}
	}
	return max
}
