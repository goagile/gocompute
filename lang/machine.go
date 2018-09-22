package lang

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
