package iteration

import "testing"

func TestRepeat(t *testing.T) {
	expected := "aaaaaaaaaa"
	actual := Repeat("a", 10)

	if expected != actual {
		t.Errorf("Expected %q got %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
