package iteration

import "testing"

func BenchmarkRepeat(t *testing.T) {
	repeaded := Repeat("a")
	expected := "aaaaa"

	if repeaded != expected {
		t.Errorf("expected %q, but got %q", expected, repeaded)
	}
}
