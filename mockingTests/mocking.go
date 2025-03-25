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

type SpySleeper2 struct {
	Calls []string
}

func (s *SpySleeper2) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpySleeper2) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

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
