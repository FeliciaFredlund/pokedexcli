package main

import "testing"

func TestCleanInput(t *testing.T) {
	testcases := map[string]struct {
		input    string
		expected []string
	}{
		"simple":            {input: "hello world", expected: []string{"hello", "world"}},
		"extra white space": {input: "  hello  world  ", expected: []string{"hello", "world"}},
		"capitals":          {input: "Hello WoRlD", expected: []string{"hello", "world"}},
		"empty":             {input: "", expected: []string{}},
		"white space only":  {input: "	 ", expected: []string{}},
	}

	for name, test := range testcases {
		actual := cleanInput(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("%s: length of actual, %#v, and expected, %#v, don't match", name, actual, test.expected)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := test.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%s: actual %#v and expected %#v does not match", name, actual, test.expected)
				break
			}
		}
	}
}
