package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Charmander Bulbasaur    PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "I am a TEApot",
			expected: []string{"i", "am", "a", "teapot"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length do not match: actual=%d; expected=%d", len(actual), len(c.expected))
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words do not match: actual=%s; expected=%s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
