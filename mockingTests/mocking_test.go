package mockingtests

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMocking(t *testing.T) {
	t.Run("First test", func(t *testing.T) {
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

	})

	t.Run("Check testing order", func(t *testing.T) {
		spySleepPrinter := &SpySleeper2{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}

	})

}
