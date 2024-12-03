package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// get the input for the puzzle
func GetPuzzleInput(fileName string) []string {

	inputPath := filepath.Join(fileName)
	content, err := os.ReadFile(inputPath)
	fmt.Printf("reading file %s", inputPath)
	if err != nil {
		fmt.Printf("failed to read input file: %s", err)
	}
	return strings.Split(string(content), "\n")

}

func Assert(condition bool, message string) {
	if !condition {
		fmt.Println("WARNING, OUTCOME IS NOT WHAT IS EXPECTED!")
		fmt.Println(message)
	}
}
