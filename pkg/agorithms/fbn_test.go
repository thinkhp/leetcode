package agorithms

import (
	"fmt"
	"testing"
)

func TestFBN(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, fbn(i))
	}
}

func fbn(i int) int {
	switch i {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fbn(i-2) + fbn(i-1)
	}
}
