package lang

import "fmt"

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

func (m *mul) Fold(env *env) Exp {
	if m.a.Foldable() {
		return Mul(m.a.Fold(env), m.b)
	}
	if m.b.Foldable() {
		return Mul(m.a, m.b.Fold(env))
	}
	return Num(m.Value())
}

func (m *mul) IsNil() bool {
	return false
}
