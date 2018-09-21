package lang

import "fmt"

//
// Number
//
type number struct {
	value int
}

func Number(value int) *number {
	return &number{value}
}

func (n *number) String() string {
	return fmt.Sprint(n.value)
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
