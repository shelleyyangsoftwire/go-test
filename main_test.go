package main

import (
	"testing"
)

var tests = []struct {
	name     string
	expected string
}{
	{"terrible", "Hello, terrible."},
	{"Robert Griesemer", "Hello, Robert. Thanks for creating me!"},
	{"Ken Thompson", "Hello, Ken. Thanks for creating me!"},
	{"Ken of Barbie", "Hello, Ken."},
	{"Far More Than Twenty Characters", "Hello, Far... Wow, that name’s too long for me!"},
	{"Bob", "Hello, Bob. Cool, a palindromic name!"},
	{"Abba abba", "Hello, Abba. Cool, a palindromic name!"},
	{"bob bob bob bob bob bob", "Hello, bob... Wow, that name’s too long for me! Cool, a palindromic name!"},
	{"bobbobbobbobbobbobbobbobbob", "Hello, bobbobbobbobbobbobbo... Wow, that name’s too long for me! Cool, a palindromic name!"},
}

var emptyTests = []struct {
	name     string
	expected bool
}{
	{"terrible", false},
	{"", true},
	{"     ", true},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("incorrect greeting for %s --  got: %s, expected: %s", test.name, result, test.expected)
		}
	}
}

func TestEmpty(t *testing.T) {
	for _, test := range emptyTests {
		result := isEmpty(test.name)
		if result != test.expected {
			t.Errorf("Input %s empty test failed -- Expected: %t, got: %t", test.name, test.expected, result)
		}
	}
}
