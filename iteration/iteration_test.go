package iteration

import (
	"testing"
)

func assertErrorMessage(t *testing.T, got, want string) {
	if got != want {
		t.Error("Failed to repeat char")
	}

}

func TestFor(t *testing.T) {

	t.Run("Repeate 3 times", func(t *testing.T) {
		got := repeat("a", 3)
		want := "aaa"

		assertErrorMessage(t, got, want)

	})

	t.Run("Repeate 4 times", func(t *testing.T) {
		got := repeat("a", 4)
		want := "aaaa"

		assertErrorMessage(t, got, want)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 3)
	}
}
