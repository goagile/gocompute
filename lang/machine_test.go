package lang

import "testing"

//
// Machine
//
func Test_Machine_Fold(t *testing.T) {
	want := 25
	env := Env()
	x := Machine(
		Sum(
			Sum(Num(5), Num(2)),
			Mul(Num(3), Num(6)),
		),
	).Run(env)

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Machine_Fold_With_Vars(t *testing.T) {
	want := 25
	env := Env()
	env.Put("x", Num(5))
	env.Put("y", Num(2))
	env.Put("p", Num(3))
	env.Put("q", Num(6))
	x := Machine(
		Sum(
			Sum(Var("x"), Var("y")),
			Mul(Var("p"), Var("q")),
		),
	).Run(env)

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}
