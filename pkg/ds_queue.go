package pkg

import (
	errs "errors"
	"github.com/pkg/errors"
)

// 队列
// 先进先出 first in,first out FIFO
type queue struct {
	size int
	head int
	tail int
	buf []data
}
func (q *queue)init(size int) {
	q.tail = 0
	q.head = 0
	q.size = size+1
	q.buf = make([]data, q.size)
}
func (q *queue)isFull() bool {
	if (q.tail+1)%q.size == q.head {
		return true
	}
	return false
}
func (q *queue)isEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}
// 入队
func (q *queue)enqueue(v data) error  {
	if q.isFull() {
		return errors.New("overflow")
	}
	if q.isEmpty() {
		q.head = 0
		q.tail = 1
		q.buf[0] = v
		return nil
	}
	q.tail = (q.tail+1) % q.size
	q.buf[q.tail] = v

	return nil
}
// 出队
func (q *queue)dequeue() (data, error) {
	if q.isEmpty() {
		return 0, errs.New("underflow 栈下溢")
	}
	q.head = (q.head+1)%q.size
	return q.buf[q.head-1],nil
}
