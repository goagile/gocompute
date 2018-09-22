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
