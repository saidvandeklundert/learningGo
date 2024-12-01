/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"aoc/pkg/test"
	"fmt"

	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Digits struct {
	Number1 *rune
	Number2 *rune
}

// takes the Number1 and Number2 rune pointer and converts it into an integer value.
func (d *Digits) GenerateNumber() int {
	number := string(*d.Number1) + string(*d.Number2)
	num, _ := strconv.Atoi(number)
	return num
}

func isDigit(r rune) bool {
	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func getDigitsFromLine(line string) Digits {
	var isDigitResult bool
	digits := Digits{}

	for _, r := range line {

		isDigitResult = isDigit(r)
		if isDigitResult {

			if digits.Number1 == nil {
				digits.Number1 = &r
			} else {
				digits.Number2 = &r
			}
		}

	}
	if digits.Number2 == nil {
		digits.Number2 = digits.Number1
	}
	return digits
}

// aoc20231aCmd represents the aoc20231a command
var aoc20231aCmd = &cobra.Command{
	Use:   "aoc20231a",
	Short: "Solution to aoc20231a",
	Long:  `Solution to the aforementioned aoc challenge.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aoc20231a called")
		total := 0
		assignmentInput := test.GetPuzzleInput("2023", "1")
		for _, line := range strings.Split(assignmentInput, "\n") {
			digits := getDigitsFromLine(line)
			fmt.Println(*digits.Number1, *digits.Number2)
			fmt.Printf("The rune value is: %c\n", *digits.Number1)
			lineNumber := digits.GenerateNumber()
			total += lineNumber

		}
		fmt.Println(total) //54951

	},
}

func init() {
	rootCmd.AddCommand(aoc20231aCmd)
}
