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
