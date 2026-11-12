package pointer

import "fmt"

// Inc 把 p 指向的整数加 1。p 为 nil 时什么也不做。
func Inc(p *int) {
	// TODO: 通过 *p 读写 p 指向的那个变量；先判断 p 是不是 nil。
	if p == nil {
		return
	}
	*p++
}

// Swap 交换 a、b 指向的两个整数。
func Swap(a, b *int) {
	if a == nil || b == nil {
		fmt.Println("输入为nil指针，请重新输入")
		return
	}
	*a, *b = *b, *a
}

// Student 是一个普通的结构体。
type Student struct {
	Name string
	Age  int
}

// Birthday 让 s 指向的学生长一岁。
func Birthday(s *Student) {
	// TODO: 结构体指针可以直接写 s.Age，Go 会自动解引用。
	if s == nil {
		return
	}
	s.Age++
}

// Renamed 返回一个改了名字的学生副本；原来的 s 不能被修改。
func Renamed(s Student, name string) Student {
	// TODO: s 是值传递进来的副本，改它不会影响调用方。
	s.Name = name
	return s
}

// Scaled 返回 a 中每个元素乘以 k 之后的新数组；原数组不能被修改。
func Scaled(a [3]int, k int) [3]int {
	// TODO: 数组是值类型，参数 a 已经是一份完整副本。
	for i, v := range a {
		a[i] = k * v
	}
	return a
}

// ScaleInPlace 把 a 指向的数组中每个元素原地乘以 k。
func ScaleInPlace(a *[3]int, k int) {
	// TODO: 想改调用方的数组，就要拿到它的地址。
	if a == nil {
		return
	}
	for i, v := range a {
		a[i] = k * v
	}
}
