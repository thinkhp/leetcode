package sort_

// 希尔排序

//希尔排序通过将比较的全部元素分为几个区域来提升插入排序的性能。
//这样可以让一个元素可以一次性地朝最终位置前进一大步。
//然后算法再取越来越小的步长gap进行排序，算法的最后一步就是普通的插入排序 ，但是到了这步，需排序的数据几乎是已排好的了（此时插入排序较快）。

//例如，假设有这样一组数[ 13 14 94 33 82 25 59 94 65 23 45 27 73 25 39 10 ]，如果我们以步长为5开始进行排序，我们可以通过将这列表放在有5列的表中来更好地描述算法，这样他们就应该看起来是这样：
//
//13 14 94 33 82
//25 59 94 65 23
//45 27 73 25 39
//10
//然后我们对每列进行排序：
//
//10 14 73 25 23
//13 27 94 33 39
//25 59 94 65 82
//45
//将上述四行数字，依序接在一起时我们得到：[ 10 14 73 25 23 13 27 94 33 39 25 59 94 65 82 45 ].这时10已经移至正确位置了，然后再以3为步长进行排序：
//
//10 14 73
//25 23 13
//27 94 33
//39 25 59
//94 65 82
//45
//排序之后变为：
//
//10 14 13
//25 23 33
//27 25 59
//39 65 73
//45 94 82
//94
//最后以1步长进行排序（此时就是简单的插入排序了）。

func shellSort(nums []int) []int {
	gap := len(nums) / 2
	//gap := 3
	for ; gap > 0; gap /= 2 {
		for i := gap; i < len(nums); i++ {
			for j := i; j >= gap && nums[j-gap] > nums[j] ; j -= gap {
				nums[j-gap], nums[j] = nums[j], nums[j-gap]
			}
		}
	}

	return nums

	//return insertSort(nums)
}

func shellSortStd(nums []int) []int {
	gap := len(nums) / 2
	//gap := 3
	for ; gap > 0; gap-- {
		for i := gap; i < len(nums); i++ {
			if nums[i-gap] > nums[i] {
				nums[i-gap], nums[i] = nums[i], nums[i-gap]
			}
		}
	}

	return nums

	//return insertSort(nums)
}
