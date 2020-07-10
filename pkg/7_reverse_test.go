package pkg

import (
	"fmt"
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println(reverse(-2147483648))
	fmt.Println(reverse(math.MaxInt32))
	fmt.Println(reverse(math.MaxInt64))
	fmt.Println(reverse(123))

}

func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}
	num := x
	numR := 0
	for {
		if num == 0 {
			break
		}
		if numR > math.MaxInt32/10 || (numR == math.MaxInt32/10 && num%10 > 7) {
			return 0
		}
		numR = numR*10 + num%10
		num = num / 10
	}
	return numR
}
