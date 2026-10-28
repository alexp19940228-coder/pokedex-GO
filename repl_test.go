package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello    world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "                hello            world        ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length of strings don't match")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Got: %v - Want: %v", word, expectedWord)
			}
		}
	}
}
