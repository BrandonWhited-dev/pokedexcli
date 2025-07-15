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
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "charmander bulbasaur",
			expected: []string{"charmander", "bulbasaur"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("The len of strings do not match. Expected: %v | Actual: %v", len(c.expected), len(actual))
			t.FailNow()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("The strings do not match. Expected: %s | Actual: %s", expectedWord, word)
				t.FailNow()
			}
		}
	}
}
