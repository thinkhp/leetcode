package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
//例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
//
//与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum-closest
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func TestThreeSumClosest(t *testing.T)  {
	ss := [][]int{
		[]int{1,2,3,4,5},
		[]int{-1,2,1,-4},
	}
	for _, v := range ss {
		fmt.Println(threeSumClosest(v, 1))
	}
}

// 双指针
func threeSumClosest(nums []int, target int) int {
	result := 0
	if len(nums) <= 3 {
		for _, v := range nums {
			result+= v
		}
		return result
	}
	sort.Ints(nums)
	fmt.Println(nums)

	result = nums[0]+nums[1]+nums[len(nums)-1]
	min := absInt(result-target)
	x := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			x = nums[i]+nums[j]+nums[k]-target
			switch {
			case x < 0:
				j++
			case x == 0:return target
			case x > 0:
				k--
			}
			if absInt(x) < min {
				min = absInt(x)
				result = target+x
			}
		}

	}

	return result
}

func absInt(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}