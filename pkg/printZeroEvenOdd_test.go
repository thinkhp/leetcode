package pkg

import (
	"fmt"
	"testing"
)

type msg struct {
	close  bool
	num    int
	maxNum int
}

var zeroChannel = make(chan msg)
var evenChannel = make(chan msg)
var oddChannel = make(chan msg)
var mainChannel = make(chan struct{})

func TestPrintZeroEvenOdd(t *testing.T) {
	m := msg{close: false, num: 1, maxNum: 1}
	go PrintEven()
	go PrintOdd()
	go PrintZero()
	zeroChannel <- m
	<-mainChannel
}

func PrintZero() {
	for m := range zeroChannel {
		if m.close {
			close(zeroChannel)
			evenChannel <- m
			break
		}
		fmt.Print("0")
		evenChannel <- m
		//fmt.Println(m)
	}
	fmt.Println()
	fmt.Println("PrintZero", "over")
}

func PrintEven() {
	for m := range evenChannel {
		if m.close {
			close(evenChannel)
			oddChannel <- m
			break
		}
		if m.num%2 != 0 && m.num != 0 {
			fmt.Print(m.num)
		}
		oddChannel <- m
		//fmt.Println(m)
	}
	fmt.Println("PrintEven", "over")
}
func PrintOdd() {
	for m := range oddChannel {
		if m.close {
			close(oddChannel)
			mainChannel <- struct{}{}
			break
		}
		if m.num%2 == 0 && m.num != 0 {
			fmt.Print(m.num)
		}
		//fmt.Println(m)

		m.num = m.num + 1
		if m.maxNum < m.num {
			m.close = true
		}
		zeroChannel <- m
	}
	fmt.Println("PrintOdd", "over")
}
