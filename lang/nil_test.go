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
