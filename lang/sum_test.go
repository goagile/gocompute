package lang

import "testing"

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
