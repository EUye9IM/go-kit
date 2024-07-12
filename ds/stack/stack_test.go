package stack_test

import (
	"testing"

	"github.com/EUye9IM/go-kit/ds/stack"
	"github.com/EUye9IM/go-kit/testtool"
)

func TestNil(t *testing.T) {
	testtool.TestCase(t, func(c testtool.Case) {
		st := (*stack.Stack[int])(nil)
		c.Assert(st.Copy(), stack.NewStack[int]())
		c.Assert(st.Cap(), 0)
		c.Assert(st.Len(), 0)
		st.Grow(10)
		c.Assert(st.Len(), 0)
		c.Assert(st.Pop(), 0)
		st.Push(0)
		c.Assert(st.Top(), 0)
	})
}
func TestCopy(t *testing.T) {
	testtool.TestCase(t, func(c testtool.Case) {
		st1 := stack.NewStack[int]()
		st1.Push(0)
		st1.Push(1)
		st1.Push(2)
		st2 := st1.Copy()
		c.Assert(st2.Top(), 2)
		st2.Push(3)
		c.Assert(st1.Top(), 2)
		c.Assert(st2.Top(), 3)
	})
}
func TestStack(t *testing.T) {
	testtool.TestCase(t, func(c testtool.Case) {
		st := stack.NewStack[int]()
		c.Assert(st.Len(), 0)
		st.Push(0)
		c.Assert(st.Len(), 1)
		st.Push(1)
		c.Assert(st.Len(), 2)
		st.Push(2)
		c.Assert(st.Len(), 3)
		c.Assert(st.Top(), 2)
		c.Assert(st.Pop(), 2)
		c.Assert(st.Len(), 2)
		c.Assert(st.Top(), 1)
		c.Assert(st.Pop(), 1)
		c.Assert(st.Len(), 1)
		c.Assert(st.Top(), 0)
		c.Assert(st.Pop(), 0)
		c.Assert(st.Len(), 0)
		c.Assert(st.Top(), 0)
		c.Assert(st.Pop(), 0)
		st.Grow(100)
		c.Assert(st.Cap(), 100)
		st.Grow(0)
		c.Assert(st.Cap(), 100)
		c.Assert(st.Len(), 0)
	})
}
