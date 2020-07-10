package concurrent

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 一个进程只有出现同时持有一把以上的锁才会出现死锁

func TestLocker(t *testing.T) {
	start()
}

// 筷子
type chopstick struct {
	sync.Mutex
	id int
}

// 哲学家
type philosopher struct {
	id        int
	left      *chopstick
	right     *chopstick
	eatTime   time.Duration
	thinkTime time.Duration
}

func (p *philosopher) initNot(id int, left, right *chopstick) {
	p.id = id

}

func (p *philosopher) init(id int, left, right *chopstick) {
	p.id = id
	p.left = left
	p.right = right
	rand.Seed(time.Now().Unix())
	p.eatTime = time.Duration(rand.Intn(10)) * time.Second
	p.thinkTime = time.Duration(rand.Intn(5)) * time.Second
	//fmt.Println(p)
}

func (p *philosopher) run() {
	log.Println("wait", p.id)
	time.Sleep(p.thinkTime)
	//fmt.Println(p)
	//fmt.Println(p.left)
	p.left.Lock()
	defer p.left.Unlock()
	p.right.Lock()
	defer p.right.Unlock()
	time.Sleep(p.eatTime)
	log.Println("ok", p.id)

}

func start() {
	ps := make([]*philosopher, 5)
	for k, _ := range ps {
		ps[k] = new(philosopher)
	}
	cs := make([]*chopstick, 5)
	for k, _ := range cs {
		cs[k] = &chopstick{id: k}
	}

	for i := 0; i < 5; i++ {
		left := cs[i]
		right := cs[i]
		if i == 4 {
			right = cs[0]
		} else {
			right = cs[i+1]
		}
		ps[i].init(i, right, left)
	}
	for i := 0; i < 5; i++ {
		go func(i int) {
			for {
				ps[i].run()
			}

		}(i)
	}

	select {}
}
