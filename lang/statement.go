package lang

import "fmt"

//
// Statement
//
type Statement interface {
	fmt.Stringer
	Name() string
	Exp() Exp
	Equal(other Statement) bool
	Foldable() bool
	Fold(*env) Statement
	Env() *env
	IsDoNothing() bool
}
