package lang

import "testing"

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
