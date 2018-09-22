package lang

import (
	"fmt"
	"sort"
)

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

func (e *env) String() string {
	s := ""
	s += fmt.Sprintf("Env:{")
	for k, v := range e.state {
		s += fmt.Sprintf("%v:%v; ", k, v.String())
	}
	s += "}"
	return s
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

func (e *env) Equal(other *env) bool {
	if e.Len() != other.Len() {
		return false
	}
	for k, v := range e.state {
		ov := other.Get(k)
		if !ov.Equal(v) {
			return false
		}
	}
	return true
}

func (e *env) Len() int {
	return len(e.state)
}

func (e *env) Keys() []string {
	s := []string{}
	for k := range e.state {
		s = append(s, k)
	}

	sort.Strings(s)
	return s
}

func (e *env) Merge(other *env) *env {
	result := Env()
	for _, k := range e.Keys() {
		result.Put(k, e.Get(k))
	}
	for _, k := range other.Keys() {
		result.Put(k, other.Get(k))
	}
	return result
}
