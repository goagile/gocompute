package lang

import "testing"

//
// Assign
//

//
// String
//
func Test_Assign_String(t *testing.T) {
	want := "x = 1"
	stmnt := Assign("x", Num(1))

	got := stmnt.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Assign_String_Sum_Var(t *testing.T) {
	want := "x = x + 1"
	stmnt := Assign("x", Sum(Var("x"), Num(1)))

	got := stmnt.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Exp
//
func Test_Assign_Exp(t *testing.T) {
	want := Num(1)
	stmnt := Assign("x", Num(1))

	got := stmnt.Exp()

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Name
//
func Test_Assign_Name(t *testing.T) {
	want := "x"
	stmnt := Assign("x", Num(1))

	got := stmnt.Name()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Env
//
func Test_Assign_Env(t *testing.T) {
	want := Env()
	stmnt := Assign("x", Num(1))

	got := stmnt.Env()

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// IsDoNothing
//
func Test_Assign_IsDoNothing(t *testing.T) {
	want := false
	stmnt := Assign("x", Num(1))

	got := stmnt.IsDoNothing()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Equal
//
func Test_Assign_Equal_True(t *testing.T) {
	want := true
	a := Assign("x", Nil())
	b := Assign("x", Nil())

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Assign_Equal_False_By_Names(t *testing.T) {
	want := false
	a := Assign("x", Nil())
	b := Assign("y", Nil())

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Assign_Equal_False_By_Exps(t *testing.T) {
	want := false
	a := Assign("x", Nil())
	b := Assign("x", Num(1))

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Foldable
//
func Test_Assign_Foldable(t *testing.T) {
	want := true
	stmnt := Assign("x", Nil())

	got := stmnt.Foldable()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Fold
//
func Test_Assign_Fold_With_Not_Foldable_Exp_Returns_DoNothing(t *testing.T) {
	want := DoNothing()
	stmnt := Assign("x", Num(1))

	got := stmnt.Fold(Env())

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Assign_Fold_With_Not_Foldable_Exp_Returns_DoNothing_With_Merged_Env(t *testing.T) {
	want := Env()
	want.Put("x", Num(232))
	want.Put("y", Num(2))

	got := Assign("x", Num(1)).Fold(want).Env()

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Assign_Fold_With_Foldable_Exp_Returns_Folded_Assign(t *testing.T) {
	want := Assign("x", Num(3))
	stmnt := Assign("x", Sum(Num(1), Num(2)))

	got := stmnt.Fold(Env())

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Assign_Fold_With_Vars(t *testing.T) {
	want := Assign("x", Sum(Num(1), Num(2)))
	stmnt := Assign("x", Sum(Var("x"), Num(2)))
	env := Env()
	env.Put("x", Num(1))

	got := stmnt.Fold(env)

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}
