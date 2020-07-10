package sort_

// 归并排序
func mergeSort(nums []int) []int {
	all := make([][]int, 0)
	for _, v := range nums {
		all = append(all, []int{v})
	}
	ret := make([][]int, 0)
	for {
		//fmt.Println("all", all)
		if len(all)%2 != 0 {
			ret = append(ret, all[0])
			all = all[1:]
		}
		for i := 0; i < len(all); i += 2 {
			//fmt.Println("1,2:", all[i], all[i+1])
			ret = append(ret, sortTwoOrder(all[i], all[i+1]))
		}
		if len(ret) == 1 {
			break
		}
		all = ret
		ret = make([][]int, 0)
	}
	return ret[0]
}

// 对两个有序数组进行合并升序
func sortTwoOrder(nums1, nums2 []int) []int {
	ret := make([]int, 0)
	i, j := 0, 0
	for {
		if i >= len(nums1) {
			ret = append(ret, nums2[j:]...)
			break
		}
		if j >= len(nums2) {
			ret = append(ret, nums1[i:]...)
			break
		}
		v1 := nums1[i]
		v2 := nums2[j]
		if v1 < v2 {
			ret = append(ret, v1)
			i++
		} else {
			ret = append(ret, v2)
			j++
		}
	}

	return ret
}

