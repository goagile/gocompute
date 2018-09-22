package lang

import "testing"

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
