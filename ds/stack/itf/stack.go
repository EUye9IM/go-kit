package itf

type StackItf[T any] interface {
	// 返回栈顶
	Top() T
	// 插入栈顶
	Push(T)
	// 弹出栈顶
	Pop() T
	// 返回元素数量
	Len() int
	// 返回占用空间
	Cap() int
	// 增长占用空间
	Grow(int)
}
