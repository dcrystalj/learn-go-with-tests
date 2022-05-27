package hello

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "Hello world"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
