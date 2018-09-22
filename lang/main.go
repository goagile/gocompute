package lang

import "fmt"

//
// Exp
//
type Exp interface {
	fmt.Stringer
	Value() int
	Equal(other Exp) bool
	Foldable() bool
	Fold() Exp
}

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

func (n *num) Equal(other Exp) bool {
	return n.Value() == other.Value()
}

func (n *num) Foldable() bool {
	return false
}

func (n *num) Fold() Exp {
	return Num(n.value)
}

//
// Sum
//
type sum struct {
	a, b Exp
}

func Sum(a, b Exp) *sum {
	return &sum{a, b}
}

func (s *sum) String() string {
	return fmt.Sprintf(`%v + %v`, s.a, s.b)
}

func (s *sum) Value() int {
	return s.a.Value() + s.b.Value()
}

func (s *sum) Equal(other Exp) bool {
	return (s.Value() == other.Value())
}

func (s *sum) Foldable() bool {
	return true
}

func (s *sum) Fold() Exp {
	if s.a.Foldable() {
		return Sum(s.a.Fold(), s.b)
	}
	if s.b.Foldable() {
		return Sum(s.a, s.b.Fold())
	}
	return Num(s.Value())
}

//
// Mul
//
type mul struct {
	a, b Exp
}

func Mul(a, b Exp) *mul {
	return &mul{a, b}
}

func (m *mul) String() string {
	return fmt.Sprintf(`%v * %v`, m.a, m.b)
}

func (m *mul) Value() int {
	return m.a.Value() * m.b.Value()
}

func (m *mul) Equal(other Exp) bool {
	return m.Value() == other.Value()
}

func (m *mul) Foldable() bool {
	return true
}

func (m *mul) Fold() Exp {
	if m.a.Foldable() {
		return Mul(m.a.Fold(), m.b)
	}
	if m.b.Foldable() {
		return Mul(m.a, m.b.Fold())
	}
	return Num(m.Value())
}
