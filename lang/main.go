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
	IsNil() bool
	Fold(*env) Exp
}

//
// Nil
//
type nilexp struct{}

func Nil() *nilexp {
	return &nilexp{}
}

func (n *nilexp) String() string {
	return "nil"
}

func (n *nilexp) Value() int {
	return 0
}

func (n *nilexp) Equal(other Exp) bool {
	return n.Value() == other.Value()
}

func (n *nilexp) Foldable() bool {
	return false
}

func (n *nilexp) Fold(env *env) Exp {
	return Nil()
}

func (n *nilexp) IsNil() bool {
	return true
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

func (n *num) Fold(env *env) Exp {
	return Num(n.value)
}

func (n *num) IsNil() bool {
	return false
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

//
// Var
//
type variable struct {
	name string
}

func Var(name string) *variable {
	return &variable{name}
}

func (v *variable) String() string {
	return v.name
}

func (v *variable) Value() int {
	return 0
}

func (v *variable) Equal(other Exp) bool {
	return false
}

func (v *variable) Foldable() bool {
	return true
}

func (v *variable) Fold(env *env) Exp {
	return env.Get(v.name)
}

func (v *variable) IsNil() bool {
	return false
}

//
// Env
//
type env struct {
	state map[string]Exp
}

func Env() *env {
	e := new(env)
	e.state = map[string]Exp{}
	return e
}

func (e *env) Put(name string, exp Exp) {
	e.state[name] = exp
}

func (e *env) Get(name string) Exp {
	exp, ok := e.state[name]
	if !ok {
		return Nil()
	}
	return exp
}

//
// Machine
//
type machine struct {
	exp Exp
}

func Machine(exp Exp) *machine {
	return &machine{exp}
}

func (m *machine) Run(env *env) Exp {
	for m.exp.Foldable() {
		m.exp = m.exp.Fold(env)
	}
	return m.exp
}
