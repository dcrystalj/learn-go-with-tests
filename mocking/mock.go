package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdown = 3

func main() {
	output := os.Stdout
	Countdown(output, &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep})
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdown; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprint(w, "Go!")
}
