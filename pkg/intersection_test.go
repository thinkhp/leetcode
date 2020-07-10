package pkg

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func HaveIntersection(n1, n2 []int) []int {
	sort.Ints(n1)
	sort.Ints(n2)

	// find
	i := 0
	j := 0
	result := make([]int, 0)
	for {
		if i >= len(n1) || j >= len(n2) {
			break
		}
		if n1[i] == n2[j] {
			result = append(result, n1[i])
			i++
			j++
		} else if n1[i] < n2[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

// 查找两个集合的交集
func TestHaveIntersection(t *testing.T) {
	n1 := []int{1, 3, 829312, 2131, 2}
	n2 := []int{232, 6, 0, 8, 2}
	//n2 := []int{232, 6, 0, 8, 9}
	//n2 := []int{232, 6, 0, 8, 829312}

	fmt.Println(HaveIntersection(n1, n2))
}

func BenchmarkIntersection(b *testing.B) {
	b.ResetTimer() //调用该函数停止压力测试的时间计数

	a1 := make([]int, 0)
	a2 := make([]int, 0)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000; i++ {
		a1 = append(a1, rand.Intn(300000))
		a2 = append(a2, rand.Intn(300000))
	}

	//b.StartTimer()
	//fmt.Println(len(inter1(a1, a2)))

	b.StartTimer()
	fmt.Println(len(inter2(a1, a2)))
}

func inter1(a1, a2 []int) []int {
	// sort
	sort.Ints(a1)
	sort.Ints(a2)

	// find
	i := 0
	j := 0
	result := make([]int, 0)
	for {
		if i >= len(a1) || j >= len(a2) {
			break
		}
		if a1[i] == a2[j] {
			result = append(result, a1[i])
			i++
			j++
		} else if a1[i] < a2[j] {
			i++
		} else {
			j++
		}
	}
	return result
}

func inter2(a1, a2 []int) []int {
	result := make([]int, 0)

	m := make(map[int]bool)
	for i := range a1 {
		m[a1[i]] = true
	}
	for i := range a2 {
		k := a2[i]
		if _, ok := m[k]; ok {
			result = append(result, k)
		}
	}

	return result
}
