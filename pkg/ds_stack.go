package pkg

import errs "errors"

// 栈 stack
// 先进后出 last in, first out LIFO
type stackValue byte

type stack struct {
	buf []stackValue
}

func (s *stack)Init()  {
	s.buf = make([]stackValue, 0)
}

func (s *stack)Push(v stackValue)  {
	s.buf = append(s.buf, v)
}

func (s *stack)Pop() (stackValue, error) {
	l := len(s.buf)
	if l == 0 {
		return 0, errs.New("underflow 栈下溢")
	}
	v := s.buf[l-1]
	s.buf = s.buf[:l-1]

	return v, nil
}

func (s *stack)IsEmpty() bool {
	return len(s.buf) == 0
}

func (s *stack)Len() int {
	return len(s.buf)
}
