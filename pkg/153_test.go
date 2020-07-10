package pkg

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	nums := []int{3,4,5,0,1,2,3}


	last := nums[0]
	min := last
	for i := 1; i < len(nums); i++ {
		if nums[i] < last {
			min = nums[i]
			break
		}
	}

	fmt.Println(min)
}
