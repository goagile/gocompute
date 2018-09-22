package lang

import "testing"

//
// Simple Expression
//
func Test_Simple_Expression_String(t *testing.T) {
	want := "5 * 2 + 3 * 6"
	x := Sum(
		Mul(Num(5), Num(2)),
		Mul(Num(3), Num(6)),
	)

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Simple_Expression_Fold(t *testing.T) {
	strigns := []string{
		"5 + 2 + 3 * 6",
		"7 + 3 * 6",
		"7 + 18",
		"25",
		"25",
	}
	env := Env()

	x := Sum(
		Sum(Num(5), Num(2)),
		Mul(Num(3), Num(6)),
	)

	expstrings := []string{
		x.String(),
		x.Fold(env).String(),
		x.Fold(env).Fold(env).String(),
		x.Fold(env).Fold(env).Fold(env).String(),
		x.Fold(env).Fold(env).Fold(env).Fold(env).String(), // Num
	}

	for i, want := range strigns {
		got := expstrings[i]
		if got != want {
			t.Fatalf("\n got : %v \n want: %v", got, want)
		}
	}
}

func Test_Simple_Expression_With_Vars(t *testing.T) {
	want := "x * y + p * q"
	x := Sum(
		Mul(Var("x"), Var("y")),
		Mul(Var("p"), Var("q")),
	)

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Simple_Expression_With_Vars_Fold(t *testing.T) {
	strigns := []string{
		"x + y + p * q",
		"5 + y + p * q",
		"5 + 2 + p * q",
		"7 + p * q",
		"7 + 3 * q",
		"7 + 3 * 6",
		"7 + 18",
		"25",
		"25",
	}
	env := Env()
	env.Put("x", Num(5))
	env.Put("y", Num(2))
	env.Put("p", Num(3))
	env.Put("q", Num(6))

	x := Sum(
		Sum(Var("x"), Var("y")),
		Mul(Var("p"), Var("q")),
	)

	expstrings := []Exp{
		x,
		x.Fold(env),
		x.Fold(env).Fold(env),
		x.Fold(env).Fold(env).Fold(env),
		x.Fold(env).Fold(env).Fold(env).Fold(env),
		x.Fold(env).Fold(env).Fold(env).Fold(env).Fold(env),
		x.Fold(env).Fold(env).Fold(env).Fold(env).Fold(env).Fold(env),
		x.Fold(env).Fold(env).Fold(env).Fold(env).Fold(env).Fold(env).Fold(env),
		x.Fold(env).Fold(env).Fold(env).Fold(env).Fold(env).Fold(env).Fold(env).Fold(env),
	}

	for i, want := range strigns {
		got := expstrings[i].String()
		if got != want {
			t.Fatalf("\n got : %v \n want: %v", got, want)
		}
	}
}
