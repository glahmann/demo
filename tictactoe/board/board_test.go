package board

import "testing"

// func TestCreate(t *testing.T) {
// 	got := Create()
// 	want := [3][3]string{
// 		{"-", "-", "-"},
// 		{"-", "-", "-"},
// 		{"-", "-", "-"},
// 	}
// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// func TestIsComplete(t *testing.T) {
// 	got := IsComplete()
// 	want := true
// }

func TestToString(t *testing.T) {
	testBoard := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
	got := ToString(testBoard)
	want := "- , - , -\n- , - , -\n- , - , -\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
