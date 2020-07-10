package pkg

import (
	"fmt"
	"testing"
)

func TestDP1(t *testing.T) {
	sum := 10
	fmt.Println(len(ot(sum, []int{})))
}

// 10 个桶,每个桶可放 0-10 条鱼,一共 180 条鱼,一共有几种方法
// x1+x2+x3+ ..... + xN = m (0 <= x <= 10)