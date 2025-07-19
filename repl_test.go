package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"simple":            {input: "hello world", expected: []string{"hello", "world"}},
		"extra white space": {input: "  hello  world  ", expected: []string{"hello", "world"}},
		"capitals":          {input: "Hello WoRlD", expected: []string{"hello", "world"}},
	}

	for name, case := range cases {
		actual := cleanInput(case.input)
		if len(actual) != len(case.expected) {
			t.Errorf("%s: length of actual, %#v, and expected, %#v, don't match", name, actual, case.expected)
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := case.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%s: actual %#v and expected %#v does not match", name, actual, case.expected)
			}
		}
	}
}
