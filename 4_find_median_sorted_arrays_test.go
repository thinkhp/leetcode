package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
//
//示例 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//示例 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5

func TestSort(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{3, 4, 5}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{2, 3}))
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 第 k 小数:二分法
func findK(nums1, nums2 []int, k int) int {
	fmt.Println(nums1, nums2, k)
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 == 0 {
		return nums2[k-1]
	}
	if l2 == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return minInt(nums1[0], nums2[0])
	}

	k1 := k / 2
	k2 := k - k1
	k1 = minInt(k1, l1)
	k2 = minInt(k2, l2)
	x1 := 0
	x2 := 0

	x1 = nums1[k1-1]
	x2 = nums2[k2-1]

	if x1 <= x2 {
		return findK(nums1[k1:], nums2, k-k1)
	} else {
		return findK(nums1, nums2[k2:k2], k-k2)
	}

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sort.Ints(nums1)
	sort.Ints(nums2)

	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	if l%2 == 0 {
		x1 := findK(nums1, nums2, l/2)
		x2 := findK(nums1, nums2, l/2+1)
		return (float64(x1) + float64(x2)) / 2
	} else {
		return float64(findK(nums1, nums2, l/2+1))
	}
}

// 双指针遍历
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	sort.Ints(nums1)
	sort.Ints(nums2)

	//ans := make([]int, 0)
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	t1, x1 := l/2-1, 0
	t2, x2 := l/2, 0
	n := 0
	i1, i2 := 0, 0
T:
	for {
		fmt.Println(i1, i2)
		if i1 >= l1 {
			n = nums2[i2]
			i2++
		} else if i2 >= l2 {
			n = nums1[i1]
			i1++
		} else if i1 < l1 && i2 < l2 {
			n1 := nums1[i1]
			n2 := nums2[i2]
			if n1 < n2 {
				i1++
				n = n1
			} else {
				i2++
				n = n2
			}
		}

		switch i1 + i2 - 1 {
		case t1:
			x1 = n
		case t2:
			x2 = n
			break T
		}
	}
	if l%2 == 0 {
		return (float64(x1) + float64(x2)) / 2
	} else {
		return float64(x2)
	}
}
