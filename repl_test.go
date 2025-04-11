package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello ",
			expected: []string{"hello"},
		},
		{
			input:    "hello  world", // Two spaces
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "\thello  world\n",
			expected: []string{"hello", "world"},
		},
		{
			input:    "one two three four",
			expected: []string{"one", "two", "three", "four"},
		},
		{
			input:    "HEllO ",
			expected: []string{"hello"},
		},
		{
			input:    "HELLO worlD ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(c.expected) != len(actual) {
			t.Errorf(`cleanInput(%v) length %v != %v, error`, c.input, len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf(`cleanInput(%v) != %v, error`, word, expectedWord)
			}
		}
	}
}
