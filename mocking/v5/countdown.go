package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}


const finalWord = "Go!"
const countDownStart = 3

/*implementa um sleeper real*/
type DefaultSleeper struct{}
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}


func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}