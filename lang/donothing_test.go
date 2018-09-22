package lang

import "testing"

//
// DoNothing
//
func Test_DoNothing(t *testing.T) {
	want := DoNothing()

	got := DoNothing()

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Name
//
func Test_DoNothing_Name(t *testing.T) {
	want := ""

	got := DoNothing().Name()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Env
//
func Test_DoNothing_Env(t *testing.T) {
	want := Env()

	got := DoNothing().Env()

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Exp
//
func Test_DoNothing_Exp(t *testing.T) {
	want := Nil()

	got := DoNothing().Exp()

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// IsDoNothing
//
func Test_DoNothing_IsDoNothing(t *testing.T) {
	want := true

	got := DoNothing().IsDoNothing()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// DoNothingWithEnv
//
func Test_DoNothingWithEnv(t *testing.T) {
	env := Env()
	want := DoNothingWithEnv(env)

	got := DoNothingWithEnv(env)

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// String
//
func Test_DoNothing_String(t *testing.T) {
	want := "do nothing"

	got := DoNothing().String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Foldable
//
func Test_DoNothing_Foldable(t *testing.T) {
	want := false

	got := DoNothing().Foldable()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Fold
//
func Test_DoNothing_Fold(t *testing.T) {
	want := true
	env := Env()
	env.Put("x", Num(1))
	do := DoNothing().Fold(env)

	got := do.Env()

	if !got.Equal(env) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}
