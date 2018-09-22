package lang

//
// DoNothing
//
type donothing struct {
	env *env
}

func DoNothingWithEnv(env *env) *donothing {
	return &donothing{env}
}

func DoNothing() *donothing {
	return &donothing{Env()}
}

func (d *donothing) Name() string {
	return ""
}

func (d *donothing) Exp() Exp {
	return Nil()
}

func (d *donothing) Env() *env {
	return d.env
}

func (d *donothing) IsDoNothing() bool {
	return true
}

func (d *donothing) Equal(other Statement) bool {
	return other.IsDoNothing()
}

func (d *donothing) String() string {
	return "do nothing"
}

func (d *donothing) Foldable() bool {
	return false
}

func (d *donothing) Fold(env *env) Statement {
	return DoNothingWithEnv(env)
}
