package lang

import "testing"

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

func Test_Env_Equal_True(t *testing.T) {
	want := true

	a := Env()
	a.Put("x", Num(1))

	b := Env()
	b.Put("x", Num(1))

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Env_Equal_False(t *testing.T) {
	want := false

	a := Env()
	a.Put("x", Num(1))

	b := Env()
	b.Put("x", Num(2))

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

// func Test_Env_Merge(t *testing.T) {
// 	want := Env()
// 	want.Put("x", Num(2))
// 	want.Put("y", Num(3))

// 	a := Env()
// 	want.Put("x", Num(1))

// 	got := env.Get("x")

// 	if !got.Equal(want) {
// 		t.Fatalf("\n got : %v \n want: %v", got, want)
// 	}
// }
