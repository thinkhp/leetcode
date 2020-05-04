package sort_

import "testing"

// 堆排序
func TestHeapSort(t *testing.T) {
	ss := [][]int{
		{1, 89, 16, 57, 5, 2, 5, 6, 1, 169},
		{5, 2, 4, 6, 1, 3},
		{1, 9, 3, 6, 3, 5, 2, 6, 4, 1, 4},
		{1, 2, 8, 6, 9, 0, 312321414, 23123},
	}
	for _, v := range ss {
		printSort(v, heapSort)
		//printSort(v, mergeSort)
	}
}

func maxHeapify(nums []int, parent, length int) {
	left := 2*parent + 1
	right := 2*parent + 2
	largest := parent
	if left < length && nums[left] > nums[parent] {
		largest = left
	}
	if right < length && nums[right] > nums[largest] {
		largest = right
	}
	if largest == parent {
		return
	}
	// 需要 parent 和子节点交换
	nums[largest], nums[parent] = nums[parent], nums[largest]
	maxHeapify(nums, largest, length)
}

func maxHeapify1(nums []int, parent, length int) {
	left := 2*parent + 1
	if left >= length { //没有叶节点不需要heap
		return
	}
	// 保证大根堆
	right := 2*parent + 2
	if right >= length { //单叶节点
		if nums[left] > nums[parent] {
			nums[left], nums[parent] = nums[parent], nums[left]
		}
		return
	}
	largest := parent
	// 双叶节点:保证 parent 大
	if nums[right] > nums[left] {
		if nums[right] > nums[parent] {
			largest = right
		}
	} else {
		if nums[left] > nums[parent] {
			largest = left
		}
	}
	if largest == parent {
		return
	}
	// 需要 parent 和子节点交换
	nums[largest], nums[parent] = nums[parent], nums[largest]
	maxHeapify1(nums, largest, length)
}

func heapSort(nums []int) []int {
	l := len(nums)
	// 1.最大堆化
	for i := l / 2; i >= 0; i-- {
		maxHeapify(nums, i, l)
	}
	for i := l - 1; i > 0; i-- {
		// 2.分离最大元素,重新最大堆化
		nums[i], nums[0] = nums[0], nums[i]
		maxHeapify(nums, 0, i)
	}
	return nums
}
