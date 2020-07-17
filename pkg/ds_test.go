package pkg

import (
	"fmt"
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	fmt.Println((math.Sqrt(5) - 1) / 2)
	fmt.Println(math.Mod(math.Sqrt(5), 1))
	fmt.Println(math.Floor(math.Sqrt(5)))
	fmt.Println(math.Trunc(math.Sqrt(5)))
	fmt.Println(math.Pow(5, 0.5))
	fmt.Println(math.Sqrt(5))
	fmt.Println(math.Sqrt(4))
}

func TestBit(t *testing.T) {
	fmt.Println(1 | 2)
	fmt.Println(300 & 13)
	fmt.Println(13 & 300)
	fmt.Println(300 % 13)
	return
}

func TestData(t *testing.T) {
	s := new(stack)
	s.Init()
	s.Push(1)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push(4)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func TestQueue(t *testing.T) {
	q := new(queue)
	q.init(4)
	fmt.Println(q.enqueue(1))
	fmt.Println(q.enqueue(2))
	fmt.Println(q.enqueue(3))
	fmt.Println(q.enqueue(4))
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.enqueue(5))
	fmt.Println(q.dequeue())

}

// 堆
type heap []interface{}
