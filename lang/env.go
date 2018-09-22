package lang

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

func (e *env) Equal(other *env) bool {
	for k, v := range e.state {
		ov := other.Get(k)
		if !ov.Equal(v) {
			return false
		}
	}
	return true
}
