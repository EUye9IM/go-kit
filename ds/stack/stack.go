package stack

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

// 深拷贝
func (s *Stack[T]) Copy() *Stack[T] {
	ret := &Stack[T]{data: make([]T, 0)}
	if s != nil && len(s.data) != 0 {
		ret.data = make([]T, len(s.data))
		copy(ret.data, s.data)
	}
	return ret
}

// 返回栈顶
func (s *Stack[T]) Top() T {
	if s == nil || len(s.data) == 0 {
		return *new(T)
	}
	return s.data[len(s.data)-1]
}

// 插入栈顶
func (s *Stack[T]) Push(val T) {
	if s == nil {
		return
	}
	s.data = append(s.data, val)
}

// 弹出栈顶
func (s *Stack[T]) Pop() T {
	if s == nil || len(s.data) == 0 {
		return *new(T)
	}
	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ret
}

// 返回元素数量
func (s *Stack[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.data)
}

// 返回占用空间
func (s *Stack[T]) Cap() int {
	if s == nil {
		return 0
	}
	return cap(s.data)
}

// 增长占用空间
func (s *Stack[T]) Grow(newCap int) {
	if s == nil || cap(s.data) >= newCap {
		return
	}
	newData := make([]T, len(s.data), newCap)
	s.data = newData
}
