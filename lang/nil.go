package lang

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
