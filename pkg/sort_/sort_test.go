package sort_

import (
	"fmt"
	"math"
	"testing"
)

func TestSort(t *testing.T) {
	fmt.Println(swap(1, 2))
	return
	ss := [][]int{
		{13, 14, 94, 33, 82, 25, 59, 94, 65, 23, 45, 27, 73, 25, 39, 10},
		{1, 89, 16, 57, 5, 2, 5, 6, 1, 169},
		{5, 2, 4, 6, 1, 3},
		{1, 9, 3, 6, 3, 5, 2, 6, 4, 1, 4},
		{1, 2, 8, 6, 9, 0, 312321414, 23123},
	}
	for _, v := range ss {
		printSort(v, shellSort)
		//printSort(v, insertSort1)
		//printSort(v, insertSort)
		//printSort(v, selectSort)
		//printSort(v, mergeSort)
		//printSort(v, heapSort)
	}
}

// 桶排序
func BucketSort(list []int) {
	buckets := make([]bool, math.MaxUint32)
	for _, value := range list {
		buckets[value] = true
	}
}

// 冒泡排序
// 空间复杂度 O(n)
func BubbleSort(list []int) {
	need := len(list) - 1
	for {
		fmt.Println(list)
		if need == 1 {
			break
		}
		for i := 0; i < need; i++ {
			if list[i] > list[i+1] {
				exchange(list, i, i+1)
			}
		}
		need--
	}
}
