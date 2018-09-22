package lang

import "testing"

//
// Nil
//
func Test_Nil_String(t *testing.T) {
	want := "nil"
	x := Nil()

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Nil_Value(t *testing.T) {
	want := 0
	x := Nil()

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Nil_Equal(t *testing.T) {
	want := true
	a := Nil()
	b := Nil()

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Nil_Foldable(t *testing.T) {
	want := false
	x := Nil()

	got := x.Foldable()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Nill_Fold(t *testing.T) {
	want := Nil()
	x := Nil()

	got := x.Fold(Env())

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Nill_IsNil(t *testing.T) {
	want := true
	x := Nil()

	got := x.IsNil()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Num
//
func Test_Num_String(t *testing.T) {
	want := "5"
	x := Num(5)

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Num_Value(t *testing.T) {
	want := 5
	x := Num(5)

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Num_Equal(t *testing.T) {
	want := true
	a := Num(5)
	b := Num(5)

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Num_Foldable(t *testing.T) {
	want := false
	x := Num(5)

	got := x.Foldable()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Num_Fold(t *testing.T) {
	want := Num(5)
	x := Num(5)

	got := x.Fold(Env())

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Num_IsNil(t *testing.T) {
	want := false
	x := Num(5)

	got := x.IsNil()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Sum
//
func Test_Sum_String(t *testing.T) {
	want := "5 + 2"
	x := Sum(Num(5), Num(2))

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Value(t *testing.T) {
	want := 7
	x := Sum(Num(5), Num(2))

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Equal(t *testing.T) {
	want := true
	a := Sum(Num(5), Num(2))
	b := Sum(Num(5), Num(2))

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Foldable(t *testing.T) {
	want := true
	x := Sum(Num(5), Num(2))

	got := x.Foldable()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Nums(t *testing.T) {
	want := Num(7)
	x := Sum(Num(5), Num(2))

	got := x.Fold(Env())

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Left(t *testing.T) {
	want := "7 + 3"
	x := Sum(
		Sum(Num(5), Num(2)),
		Num(3),
	).Fold(Env())

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Left_Value(t *testing.T) {
	want := 10
	x := Sum(
		Sum(Num(5), Num(2)),
		Num(3),
	).Fold(Env())

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Right(t *testing.T) {
	want := "3 + 7"
	x := Sum(
		Num(3),
		Sum(Num(5), Num(2)),
	).Fold(Env())

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Right_Value(t *testing.T) {
	want := 10
	x := Sum(
		Num(3),
		Sum(Num(5), Num(2)),
	).Fold(Env())

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Left_And_Right(t *testing.T) {
	want := "7 + 6"
	env := Env()
	x := Sum(
		Sum(Num(5), Num(2)),
		Sum(Num(2), Num(4)),
	).Fold(env).Fold(env)

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Left_And_Right_Value(t *testing.T) {
	want := 13
	env := Env()
	x := Sum(
		Sum(Num(5), Num(2)),
		Sum(Num(2), Num(4)),
	).Fold(env).Fold(env)

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_IsNil(t *testing.T) {
	want := false
	x := Sum(Nil(), Nil())

	got := x.IsNil()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Mul
//
func Test_Mul_String(t *testing.T) {
	want := "5 * 2"
	x := Mul(Num(5), Num(2))

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Value(t *testing.T) {
	want := 10
	x := Mul(Num(5), Num(2))

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Equal(t *testing.T) {
	want := true
	a := Mul(Num(5), Num(2))
	b := Mul(Num(5), Num(2))

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Foldable(t *testing.T) {
	want := true
	x := Mul(Num(5), Num(2))

	got := x.Foldable()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Nums(t *testing.T) {
	want := Num(10)
	x := Mul(Num(5), Num(2))

	got := x.Fold(Env())

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Left(t *testing.T) {
	want := "10 * 2"
	x := Mul(
		Mul(Num(5), Num(2)),
		Num(2),
	).Fold(Env())

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Left_Value(t *testing.T) {
	want := 20
	x := Mul(
		Mul(Num(5), Num(2)),
		Num(2),
	).Fold(Env())

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Righ(t *testing.T) {
	want := Num(20)
	x := Mul(
		Num(2),
		Mul(Num(5), Num(2)),
	)

	got := x.Fold(Env())

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Righ_Value(t *testing.T) {
	want := 20
	x := Mul(
		Num(2),
		Mul(Num(5), Num(2)),
	).Fold(Env())

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Left_And_Righ(t *testing.T) {
	want := "12 * 10"
	env := Env()
	x := Mul(
		Mul(Num(4), Num(3)),
		Mul(Num(5), Num(2)),
	).Fold(env).Fold(env)

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Left_And_Righ_Value(t *testing.T) {
	want := 120
	env := Env()
	x := Mul(
		Mul(Num(4), Num(3)),
		Mul(Num(5), Num(2)),
	).Fold(env).Fold(env)

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_IsNil(t *testing.T) {
	want := false
	x := Mul(Nil(), Nil())

	got := x.IsNil()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Env
//
func Test_Env_Put_Get(t *testing.T) {
	want := Num(5)
	env := Env()
	env.Put("x", Num(5))

	got := env.Get("x")

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Env_Get_Nil(t *testing.T) {
	want := Nil()
	env := Env()

	got := env.Get("x")

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Variable
//
func Test_Var_String(t *testing.T) {
	want := "x"
	x := Var("x")

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Var_Value(t *testing.T) {
	want := 0
	x := Var("x")

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Var_Equal(t *testing.T) {
	want := false
	x := Var("x")
	y := Var("y")

	got := x.Equal(y)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Var_Foldable(t *testing.T) {
	want := true
	x := Var("x")

	got := x.Foldable()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Var_Fold(t *testing.T) {
	want := Num(5)
	env := Env()
	env.Put("x", Num(5))
	x := Var("x")

	got := x.Fold(env)

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

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
