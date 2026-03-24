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
			input:    "   hello    world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello! ,world   ",
			expected: []string{"hello!", ",world"},
		},
		{
			input:    "  HeLLO   WoRlD",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual world not equal expected word")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Got: %v Expected: %v", word, expectedWord)
			}
		}
	}

}
