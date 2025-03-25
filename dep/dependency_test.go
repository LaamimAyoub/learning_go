package dep

import (
	"bytes"
	"testing"
)

func TestDependecy(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chriss")

	got := buffer.String()
	want := "Hello, Chriss"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
