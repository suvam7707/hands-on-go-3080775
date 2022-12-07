// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one", "#00", 0},
		{"two", "a2_32_^_&/)", 1},
		{"three", "812_%6ab//", 2},
	}

	lc := letterCounter{}

	// range over the tests and run them as subtests
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := lc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one", "#00", 2},
		{"two", "abc_1,?/", 1},
		{"three", "abc_12&8_^_", 3},
	}

	nc := numberCounter{}

	// range over the tests and run them as subtests
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := nc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one", "#00", 1},
		{"two", "abc_1,?/", 4},
		{"three", "abc_12&8_^_", 5},
	}

	sc := symbolCounter{}

	// range over the tests and run them as subtests
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := sc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
