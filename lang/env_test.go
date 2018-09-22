package lang

import (
	"reflect"
	"testing"
)

//
// Env
//

//
// Put Get
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
// String
//
func Test_Env_String(t *testing.T) {
	want := `Env:{x:1; }`
	env := Env()
	env.Put("x", Num(1))

	got := env.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Len
//
func Test_Env_Len_Zero(t *testing.T) {
	want := 0
	env := Env()

	got := env.Len()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Env_Len(t *testing.T) {
	want := 3
	env := Env()
	env.Put("x", Num(1))
	env.Put("y", Num(2))
	env.Put("z", Num(2))

	got := env.Len()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Keys
//
func Test_Env_Keys(t *testing.T) {
	want := []string{"a", "b", "c"}
	env := Env()
	env.Put("a", Num(1))
	env.Put("c", Num(3))
	env.Put("b", Num(2))

	got := env.Keys()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Equal
//
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

func Test_Env_Equal_False_With_Other_Key(t *testing.T) {
	want := false

	a := Env()
	a.Put("x", Num(1))

	b := Env()
	b.Put("y", Num(2))

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Env_Equal_False_By_Len(t *testing.T) {
	want := false

	a := Env()
	a.Put("x", Num(1))

	b := Env()

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Merge
//
func Test_Env_Merge_Add(t *testing.T) {
	want := Env()
	want.Put("x", Num(1))
	want.Put("y", Num(2))

	a := Env()
	a.Put("x", Num(1))

	b := Env()
	b.Put("y", Num(2))

	got := a.Merge(b)

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}
