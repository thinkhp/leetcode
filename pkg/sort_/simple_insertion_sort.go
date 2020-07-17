package sort_

// 简单插入排序
// 维护一个有序集合,再不断地插入新的元素
// 空间复杂度 O(n)
// 时间复杂度 1+2+3+4+......+n = ½*n*(n+1) = O(n²)

func insertionSort(nums []int, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

// 在原切片上交换,空间复杂度为 1
func insertSort1(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
	return nums
}

// 维护一个新的有序的切片 ret ,不断抽取旧切片 nums 的值插入新切片 ret, 空间复杂度为 O(n)
func insertSort(nums []int) (ret []int) {
	for _, v := range nums {
		ret = insertIntoSorted(ret, v)
	}
	return
}

func insertIntoSorted(nums []int, num int) (ret []int) {
	if len(nums) == 0 {
		return []int{num}
	}
	if num <= nums[0] {
		return append([]int{num}, nums...)
	}
	if num >= nums[len(nums)-1] {
		return append(nums, num)
	}

	index := 0
	for i, v := range nums {
		index = i
		if v > num {
			break
		}
	}

	//ret = make([]int, index, len(nums)+1)
	//// 必须使用 copy, 在原切片上 append 会导致底层数组的改变
	//copy(ret, nums[:index])
	//ret = append(ret, num)
	//ret = append(ret, nums[index:]...)
	//fmt.Println(nums, "add", num, "=>", ret)
	//return ret

	ret = make([]int, 0, len(nums)+1)
	ret = append(append(append(ret, nums[:index]...), num), nums[index:]...)
	return
}
