package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chriss"

	var got []string

	x := struct {
		name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("Wrong number of function calls, got %d, want %d", len(got), 1)

	}

}
