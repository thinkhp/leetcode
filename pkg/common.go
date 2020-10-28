package pkg

// todo 使用二分的思想减少时间复杂度
func maxInt(compare ...int) int {
	max := compare[0]
	for i := 1; i < len(compare); i++ {
		v := compare[i]
		if max < v {
			max = v
		}
	}

	return max
}
