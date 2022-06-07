package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalk(t *testing.T) {

	got := make([]string, 0)

	x := struct {
		Name string
	}{"Chris"}

	walk(x, func(input string) {
		got = append(got, input)
	})

	assert.Equal(t, []string{"Chris"}, got)
}

type Person struct {
	City string
	Age  int
	Name string
}

type PersonA struct {
	City string
	Age  int
	Name string
	Address
}

type Address struct {
	Village string
	Num     int
}

func TestWalkCases(t *testing.T) {
	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{"A", struct {
			Name string
			Age  int
		}{"Chris", 30}, []string{"Chris"}},
		{"B", struct {
			Age  int
			Name string
		}{30, "Chris"}, []string{"Chris"}},
		{"C", Person{"NY", 30, "Chris"}, []string{"NY", "Chris"}},
		{"Nested", PersonA{"NY", 30, "Chris", Address{"Vrhp", 3}}, []string{"NY", "Chris", "Vrhp"}},
		{"Nested pointer", &PersonA{"NY", 30, "Chris", Address{"Vrhp", 3}}, []string{"NY", "Chris", "Vrhp"}},
	}

	for _, c := range cases {
		got := make([]string, 0)

		walk(c.Input, func(input string) {
			got = append(got, input)
		})

		assert.Equal(t, c.Expected, got)
	}
}
