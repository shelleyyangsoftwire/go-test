package main

import (
	"testing"
)

var tests = []struct {
	name     string
	expected string
}{
	{"terrible", "Hello, terrible"},
	{"green", "Hello, green"},
	{"cola", "Hello, cola"},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("incorrect greeting for %s --  got: %s, expected: %s", test.name, result, test.expected)
		}
	}

}
