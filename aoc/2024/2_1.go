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

func day2_2() {
	result := GetPuzzleInput("inputs/2.txt")
	total_safe := 0
	for _, line := range result {
		isSafe := isLineSafeDampened(line)
		if isSafe {

			total_safe += 1
		}
	}
	fmt.Printf("\ntotal safe lines dampened %d (should be 663)\n", total_safe)

}

func isLineSafeDampened(line string) bool {

	chunks := strings.Split(line, " ")
	slice := make([]int, 0)
	for _, char := range chunks {
		number, _ := strconv.Atoi(char)
		slice = append(slice, number)

	}
	//fmt.Println(line)
	//fmt.Println(slice)
	if isSafeSlice(slice) {
		return true
	} else {
		for idx := range len(slice) {
			temp_slice := removeIndex(slice, idx)
			if isSafeSlice(temp_slice) {
				return true
			}
		}

	}
	return false
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isSafeSlice(slice []int) bool {
	for idx := range len(slice) - 1 {
		if idx == 0 {
			continue
		}
		integer := slice[idx]
		previous := slice[idx-1]
		next := slice[idx+1]
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
