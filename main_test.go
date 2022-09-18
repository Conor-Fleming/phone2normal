package main

import "testing"

// creating struct with input value and expected result to drive test
var numbers = []struct {
	input    string
	expected string
}{
	{
		input:    "1234567890",
		expected: "1234567890",
	},
	{
		input:    "123 456 7890",
		expected: "1234567890",
	},
	{
		input:    "(123)4567890",
		expected: "1234567890",
	},
	{
		input:    "(123) 456 7890",
		expected: "1234567890",
	},
	{
		input:    "(123)456-7890",
		expected: "1234567890",
	},
	{
		input:    "123-456-7890",
		expected: "1234567890",
	},
	{
		input:    "(123)-456-7890",
		expected: "1234567890",
	},
	{
		input:    "1234t67890",
		expected: "123467890",
	},
}

func TestNormalize(t *testing.T) {
	for _, test := range numbers {
		if result := normalize(test.input); result != test.expected {
			t.Errorf("\n Input ---> %s\n Test result ----> %s\n Expected ---> %s\n", test.input, result, test.expected)
		}
	}
}
