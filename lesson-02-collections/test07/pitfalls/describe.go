package pitfalls

import (
	"fmt"
	"reflect"
)

// Point 是平面上的一个点。
type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// TODO: 给 Point 实现 fmt.Stringer，让 fmt.Sprint(Point{1, 2}) 输出 "(1, 2)"。

// Describe 用 type switch 描述任意值 v。按下面的顺序匹配，先匹配到的优先：
//
//	v == nil            -> "nil"
//	error               -> "error: " + 错误信息
//	fmt.Stringer        -> "stringer: " + String() 的结果
//	int                 -> "int: 42"
//	string              -> "string: \"hi\""（用 %q）
//	[]int               -> "[]int of len 3"
//	其他                -> "other: " + 类型名（用 %T，例如 "other: float64"）
func Describe(v any) string {
	// TODO: switch x := v.(type) { ... }
	switch value := v.(type) {
	case nil:
		return "nil"
	case error:
		return "error: " + value.Error()
	case fmt.Stringer:
		return "stringer: " + value.String()
	case int:
		return fmt.Sprintf("int: %d", value)
	case string:
		return fmt.Sprintf("string: %q", value)
	case []int:
		return fmt.Sprintf("[]int of len %d", len(value))
	default:
		return fmt.Sprintf("other: %T", value)

	}
}

// SafeEqual 比较两个 any 是否相等。
// 当两边的动态类型不可比较（例如 []int、map）时，返回 false, false，而不是 panic；
// 可比较时返回 a == b, true。
//
// BUG: 直接写 a == b，两边都是切片时会 panic: runtime error: comparing uncomparable type。
func SafeEqual(a, b any) (equal bool, comparable bool) {
	if a == nil || b == nil {
		return a == b, true
	}
	t := reflect.TypeOf(a)
	z := reflect.TypeOf(b)

	if t != z {
		return false, true
	}
	if !t.Comparable() {
		return false, false
	}

	return a == b, true
}
