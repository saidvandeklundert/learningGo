package main

import (
	"fmt"
	"os"
	"strings"
)

/*

The strings library contains all the common methods one needs.

Unlike in Python, strings do not inherit a lot of methods that you can apply to them.

With the strings library, you have a selection of functions that operate on a string and that return a new
 string.

*/

func main() {
	baseString := "The base string"
	lowerCaseBaseString := strings.ToLower(baseString) // the base string
	fmt.Println(lowerCaseBaseString)
	splitString := strings.Split("192.1.1.2", ".") //[192 1 1 2]
	fmt.Println(splitString)
	baseStringConcatenated := baseString + " plus something else"
	fmt.Println(baseStringConcatenated) // The base string plus something else
	formattedString := fmt.Sprintf("kind of like %s but not really", "f-string")
	fmt.Println(formattedString) // kind of like f-string but not really

	// Example
	configuration, _ := os.ReadFile("cfg.txt")

	lines := strings.Split(string(configuration), "\n")
	// print all lines:
	for i, line := range lines {
		fmt.Printf("line number %d:\t%s\n", i, line)
	}
	// hunt for lines with certain characteristics:
	for _, line := range lines {
		chunks := strings.Split(line, " ")
		fmt.Println(chunks) // [set protocols mstp interface ae3]
		if len(chunks) >= 3 {
			if chunks[2] == "mstp" {
				fmt.Println(line)
			}
		}
	}
}
