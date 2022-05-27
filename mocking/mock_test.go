package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	spy := &SpySleeper{Calls: 0}
	Countdown(buffer, spy)

	assert.Equal(t, "3\n2\n1\nGo!", buffer.String())
	assert.Equal(t, 4, spy.Calls)
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdownOrder(t *testing.T) {
	spy := &SpyCountdownOperations{}
	Countdown(spy, spy)

	assert.Equal(t, []string{
		sleep, write, sleep, write, sleep, write, sleep, write,
	}, spy.Calls)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
