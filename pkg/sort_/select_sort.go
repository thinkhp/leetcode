package sort_

func exchange(list []int, x, y int) {
	temp := list[x]
	list[x] = list[y]
	list[y] = temp
}

// 简单选择排序
// 不断找出最值,交换
func selectSort(nums []int) (ret []int) {
	return ret
}

// 二元选择排序
// 找出最大值+最小值,交换
// 空间复杂度 O(n)
// 时间复杂度 n + n-2 + n-4 + ... + 1 =  ¼ * n * (n-1)
func selectTwoSort(list []int) []int {
	i := 0
	l := len(list)
	for {
		max := list[i]
		maxIndex := i
		min := list[i]
		minIndex := i
		//fmt.Println(i, l, list)
		if i >= l {
			break
		}
		for j := i; j < l; j++ {
			if list[j] > max {
				max = list[j]
				maxIndex = j
			}
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		exchange(list, i, minIndex)
		exchange(list, l-1, maxIndex)

		i++
		l--
	}

	return list
}
