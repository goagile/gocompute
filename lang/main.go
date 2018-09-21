package lang

import "fmt"

//
// Num
//
type num struct {
	value int
}

func Num(value int) *num {
	return &num{value}
}

func (n *num) String() string {
	return fmt.Sprint(n.value)
}

func (n *num) Value() int {
	return n.value
}

func (n *num) Equal(other *num) bool {
	return n.Value() == other.Value()
}

func (n *num) Foldable() bool {
	return false
}

func (n *num) Fold() *num {
	return Num(n.value)
}

//
// Sum
//
type sum struct {
	a, b fmt.Stringer
}

func Sum(a, b fmt.Stringer) *sum {
	return &sum{a, b}
}

func (s *sum) String() string {
	return fmt.Sprintf(`%v + %v`, s.a, s.b)
}

func (n *sum) Foldable() bool {
	return true
}

//
// Mul
//
type mul struct {
	a, b fmt.Stringer
}

func Mul(a, b fmt.Stringer) *mul {
	return &mul{a, b}
}

func (m *mul) String() string {
	return fmt.Sprintf(`%v * %v`, m.a, m.b)
}

func (n *mul) Foldable() bool {
	return true
}
