package lang

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
