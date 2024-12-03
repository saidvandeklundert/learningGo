package test

import (
	"fmt"
	"os"
	"path/filepath"
)

// get the input for the puzzle
func GetPuzzleInput(year string, puzzle_number string) string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	inputPath := filepath.Join(wd, year, puzzle_number+".txt")
	content, err := os.ReadFile(inputPath)
	fmt.Printf("reading file %s", inputPath)
	if err != nil {
		fmt.Printf("failed to read input file: %s", err)
	}
	return string(content)

}

func Assert(condition bool, message string) {
	if !condition {
		fmt.Println("WARNING, OUTCOME IS NOT WHAT IS EXPECTED!")
		fmt.Println(message)
	}
}
