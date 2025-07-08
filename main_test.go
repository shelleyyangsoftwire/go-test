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
	{"Robert Griesemer", "Hello, Robert. Thanks for creating me!"},
	{"Ken Thompson", "Hello, Ken. Thanks for creating me!"},
	{"Ken of Barbie", "Hello, Ken"},
	{"Far More Than Twenty Characters", "Hello, Far... Wow, that name’s too long for me!"},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("incorrect greeting for %s --  got: %s, expected: %s", test.name, result, test.expected)
		}
	}

}
