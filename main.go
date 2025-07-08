package main

import (
	"bufio"
	"fmt"
	//"math"
	"os"
	"slices"
	"strings"
)

func getName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	name := getName()
	name = strings.TrimSpace(name)
	if !isEmpty(name) {
		fmt.Println(greeting(name))
	} else {
		fmt.Println("Ok, no greeting for you")
	}
}

func shorten(name string) string {
	if len(name) > 20 {
		name = name[:20]
	}
	firstSpace := strings.Index(name, " ")
	if firstSpace == -1 {
		return name
	}
	return name[:firstSpace]
}

func isPalindrome(name string) bool {
	name = strings.ToLower(name)
	for i := 0; i < len(name)/2; i++ {
		if name[i] != name[len(name)-i-1] {
			return false
		}
	}
	return true
}

// isEmpty checks
func isEmpty(name string) bool {
	name = strings.TrimSpace(name)
	if name == "" {
		return true
	}
	return false
}

func greeting(name string) string {
	suffix := "."
	makers := []string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}
	if slices.Contains(makers, name) {
		suffix = ". Thanks for creating me!"
	}

	if len(name) > 20 {
		suffix += ".. Wow, that name’s too long for me!"
	}

	if isPalindrome(name) {
		suffix += " Cool, a palindromic name!"
	}

	name = shorten(name)

	return "Hello, " + name + suffix
}
