package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	in := [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}}
	out := longestLine(in)
	expected := 3

	if out != expected {
		t.Errorf("got %d, want %d", out, expected)
	}

}
