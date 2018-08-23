package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper allows to to put delays
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is an implementation of Sleeper with a defined delay
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep will pause execution  for the defined duration
func (cs ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

// Countdown print count
func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		sleeper.Sleep()
		fmt.Fprintf(w, "%v\n", i)
	}
	sleeper.Sleep()
	fmt.Fprintf(w, finalWord)
}

func main() {
	Countdown(os.Stdout, ConfigurableSleeper{3 * time.Second, time.Sleep})
}
