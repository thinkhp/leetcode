package leetcode

import (
	"fmt"
	"testing"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//示例:
//给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]

func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 3, 5}
	target := 8
	//nums := []int{-3,4,3,90}
	//target := 0
	//nums := []int{3,3}
	//target := 6

	fmt.Println(twoSum(nums, target))
}

func SliceToMap(s []int) map[int][]int {
	m := make(map[int][]int)
	for index, num := range s {
		m[num] = append(m[num], index)
	}
	return m
}

// array 转 hash,空间换时间
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return []int{}
}

func twoSum1(nums []int, target int) []int {
	var fast = func() []int {
		m := SliceToMap(nums)
		for n1, i1 := range m {
			tmp := target - n1
			if _, ok := m[tmp]; ok {
				for _, vv := range m[tmp] {
					if i1[0] != vv {
						res := []int{vv, i1[0]}
						//sort.Ints(ret)
						return res
					}
				}
			}
		}
		return nil
	}
	var small = func() []int {
		for i, v1 := range nums {
			for j := i + 1; j < len(nums); j++ {
				if v1+nums[j] == target {
					return []int{i, j}
				}
			}
		}
		return nil
	}
	_ = small()
	return fast()
}
