package main

import (
	"testing"
)

func TestForTesting(t *testing.T) {
	cases := []struct {
		test1, test2, want string
	}{
		{"Markus", "Smith", "Markus Smith"},
		{"Alex", "Springer", "Alex Springer"},
		{"", "", " "},
		{"Adam", "", "Adam"},
	}

	for _, item := range cases {
		res := forTesting(item.test1, item.test2)
		if res != item.want {
			t.Errorf("forTesting(%v) == %v, want %v", item.test1, item.test2, item.want)
		}
	}

}
