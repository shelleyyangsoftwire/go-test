package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func getName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	name := getName()
	fmt.Println(greeting(name))
}

func greeting(name string) string {
	suffix := ""
	makers := []string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}
	if slices.Contains(makers, name) {
		suffix = ". Thanks for creating me!"
	} else if len(name) > 20 {
		name = name[:20]
		suffix = "... Wow, that name’s too long for me!"
	}

	return "Hello, " + name + suffix
}
