package lang

import "testing"

//
// Number
//
func Test_Number_String(t *testing.T) {
	want := "5"
	x := Number(5)

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Sum
//
func Test_Sum_String(t *testing.T) {
	want := "5 + 2"
	x := Sum(Number(5), Number(2))

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Mul
//
func Test_Mul_String(t *testing.T) {
	want := "5 * 2"
	x := Mul(Number(5), Number(2))

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Simple_Expression_String(t *testing.T) {
	want := "5 * 2 + 3 * 6"
	x := Sum(
		Mul(Number(5), Number(2)),
		Mul(Number(3), Number(6)),
	)

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}
