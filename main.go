package main

import (
	mockingtests "mymodule/mockingTests"
	"os"
)

func main() {
	dsleeper := mockingtests.DefaultSpleeper{}
	mockingtests.CountDown(os.Stdout, &dsleeper)
}
