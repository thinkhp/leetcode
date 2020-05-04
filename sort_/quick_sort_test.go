package sort_

import (
	"fmt"
	"sort"
	"testing"
)
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}
func TestUnit(t *testing.T) {
	sort.Ints([]int{})
	fmt.Println(maxDepth(10))
	return
	//v1 := []int{1, 3, 5, 6, 6, 7, 8, 10}
	//v2 := []int{2, 3, 4, 5, 6, 7, 19, 20, 56}
	v1 := []int{1}
	v2 := []int{9, 3}
	v := sortTwoOrder(v1, v2)
	fmt.Println(len(v1)+len(v2), len(v), v)
}

func TestQuickSort(*testing.T) {
	ss := [][]int{
		{3,1,9,3,3,2,6},
		[]int{1, 9, 3, 6, 3, 5, 2, 6, 4, 1, 4},
	}

	for _, v := range ss {
		printSort(v, quickSort)
	}

}


// 快速排序
func sort2(nums []int) []int {
	return nums
}

func sort3(nums []int) []int {
	return nums
}
//从数列中挑出一个元素，称为 "基准"（pivot）;
//
//重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
//
//递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
// 3i,1,9,3,3,2,6j
// 3i,1,9,3,3,2j,6
// 2,1i,9,3,3,3j,6
// 2,1,3i,3,3j,9,6
func quickSort(nums []int) []int{
	l := len(nums)
	if l <= 1 {
		return nums
	}
	pivot := nums[0]
	spiltIndex := 0
	forwardLeft := true
	for i, j := 0, l-1;; {
		if i == j {
			spiltIndex = i
			break
		}
		if forwardLeft {
			if nums[j] >= pivot {
				j--
			} else {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				forwardLeft = false
			}
		} else {
			if nums[i] <= pivot {
				i++
			} else {
				nums[i], nums[j] = nums[j], nums[i]
				j--
				forwardLeft = true
			}
		}
	}
	if spiltIndex == 0 {
		spiltIndex++
	}
	//fmt.Println(nums[:spiltIndex], nums[spiltIndex:])
	return append(quickSort(nums[:spiltIndex]), quickSort(nums[spiltIndex:])...)
}
