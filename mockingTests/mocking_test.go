package mockingtests

import (
	"bytes"
	"testing"
)

func TestMocking(t *testing.T) {
	buffer := &bytes.Buffer{}
	spsleeper := SpySleeper{}
	CountDown(buffer, &spsleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	if spsleeper.Calls != 3 {
		t.Errorf("didnt call spsleeper 3 times")
	}

}
