package mockingtests

import (
	"fmt"
	"io"
	"time"
)

type SpySleeper struct {
	Calls int
}

type Sleeper interface {
	Sleep()
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSpleeper struct{}

func (s *DefaultSpleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func CountDown(out io.Writer, Sle Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		Sle.Sleep()
	}
	fmt.Fprint(out, "Go!")
}
