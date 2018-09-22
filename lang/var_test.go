package lang

import "testing"

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

func Test_Var_IsNil(t *testing.T) {
	want := false
	x := Var("x")

	got := x.IsNil()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}
