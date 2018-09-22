package lang

import "fmt"

//
// Assing
//

type assign struct {
	name string
	exp  Exp
	env  *env
}

func Assign(name string, exp Exp) *assign {
	return &assign{name, exp, Env()}
}

func (a *assign) Exp() Exp {
	return a.exp
}

func (a *assign) Name() string {
	return a.name
}

func (a *assign) Env() *env {
	return a.env
}

func (a *assign) IsDoNothing() bool {
	return false
}

func (a *assign) String() string {
	return fmt.Sprintf("%v = %v", a.name, a.exp)
}

func (a *assign) Equal(other Statement) bool {
	return (a.name == other.Name() &&
		a.exp.Equal(other.Exp()))
}

func (a *assign) Foldable() bool {
	return true
}

func (a *assign) Fold(env *env) Statement {
	if a.exp.Foldable() {
		return &assign{a.name, a.exp.Fold(env), env}
	}
	return DoNothingWithEnv(a.Env().Merge(env))
}
