package iteration

import "testing"

func TestRepeat(t *testing.T) {
	n := 10
	char := "a"
	repeated := Repeat(char, n)
	expected := ""
	for i := 0; i < n; i++ {
		expected += char
	}

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
