package main

import (
	"testing"
)

func TestCleaninput(t *testing.T) {
	cases := []struct {
	input    string
	expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "This  is my String",
			expected: []string{"this", "is", "my", "string"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]	
			if word != expectedWord {
				t.Errorf("word: %s is not %s", word, expectedWord)
			}
		}
	}
}