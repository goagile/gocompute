package lang

import "fmt"

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

func (s *sum) Fold(env *env) Exp {
	if s.a.Foldable() {
		return Sum(s.a.Fold(env), s.b)
	}
	if s.b.Foldable() {
		return Sum(s.a, s.b.Fold(env))
	}
	return Num(s.Value())
}

func (s *sum) IsNil() bool {
	return false
}
