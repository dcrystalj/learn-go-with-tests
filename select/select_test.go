package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	got := Racer(slowURL, fastURL)

	assert.Equal(t, fastURL, got)
}

func TestRacerWithMock(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRun(t *testing.T) {
	serverA := makeDelayedServer(1 * time.Second)
	serverB := makeDelayedServer(2 * time.Second)

	defer serverA.Close()
	defer serverB.Close()

	_, err := Racer2(serverA.URL, serverB.URL, 1*time.Second)

	if err == nil {
		t.Error("expected an error but didn't get one")
	}
}
