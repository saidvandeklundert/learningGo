/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"aoc/pkg/test"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func toInteger(number string) string {
	var wordToNumber = map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	if value, exists := wordToNumber[number]; exists {
		return value
	} else {
		return number
	}
}

func getLineNumbers(line string) (string, string) {
	interesting := [20]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	tracker := make(map[int]string)

	for _, number := range interesting {
		res := strings.Index(line, number)
		if res != -1 {
			tracker[res] = number
		}
		last_occurence := strings.LastIndex(line, number)
		if res != -1 {
			tracker[last_occurence] = number
		}

	}

	fmt.Println(tracker)
	keys := make([]int, 0, len(tracker))

	for k := range tracker {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	first_number_key := keys[0]
	second_number_key := keys[len(keys)-1]
	first_number := tracker[first_number_key]
	second_number := tracker[second_number_key]
	return toInteger(first_number), toInteger(second_number)

}

// takes the Number1 and Number2 rune pointer and converts it into an integer value.
func GenerateNumber(number1 string, number2 string) int {
	number := string(number1) + string(number2)
	num, _ := strconv.Atoi(number)
	return num
}
func Puzzle20231b() {

	assignmentInput := test.GetPuzzleInput("2023", "1")
	fmt.Println((assignmentInput))
	fmt.Println("aoc20231b ran")
	solution := 0
	for _, line := range strings.Split(assignmentInput, "\n") {
		fmt.Println(line)
		number1, number2 := getLineNumbers(line)
		line_result := GenerateNumber(number1, number2)

		solution += line_result
	}
	fmt.Println(solution)
	expected := 55218
	test.Assert(solution == expected, fmt.Sprintf("got %d want %d", solution, expected))

}

// aoc20231bCmd represents the aoc20231b command
var aoc20231bCmd = &cobra.Command{
	Use:   "aoc20231b",
	Short: "aoc20231b",
	Long:  `aoc20231b`,
	Run: func(cmd *cobra.Command, args []string) {
		Puzzle20231b()

	},
}

func init() {
	rootCmd.AddCommand(aoc20231bCmd)
}
