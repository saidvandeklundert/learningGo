package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2_1() {
	result := GetPuzzleInput("inputs/2.txt")
	total_safe := 0
	for _, line := range result {
		isSafe := isLineSafe(line)
		if isSafe {
			total_safe += 1
		}
	}
	fmt.Printf("\ntotal safe lines %d (should be 663)\n", total_safe)
	// 663 is correct!!
	// 496 is too low
	// 690 is too high
	// 778 too high

}

func isLineSafe(line string) bool {

	chunks := strings.Split(line, " ")
	for idx := range len(chunks) - 1 {
		if idx == 0 {
			continue
		}
		integer, _ := strconv.Atoi(chunks[idx])
		previous, _ := strconv.Atoi(chunks[idx-1])
		next, _ := strconv.Atoi(chunks[idx+1])
		if integer == previous || integer == next {
			return false
		}

		if integer < previous && integer < next {
			return false
		} else if integer > previous && integer > next {
			return false
		}
		diff_prev := next - integer
		if diff_prev < 0 {
			diff_prev = diff_prev * -1
		}
		if diff_prev > 3 {
			return false
		}
		diff_next := previous - integer
		if diff_next < 0 {
			diff_next = diff_next * -1
		}
		if diff_next > 3 {
			return false
		}

	}
	return true

}
