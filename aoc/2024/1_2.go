package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day1_2() {
	result := GetPuzzleInput("inputs/1.txt")
	left_numbers := make(map[int]int)
	left := make([]int, len(result))
	right := make([]int, len(result))
	for idx, line := range result {
		chunks := strings.Split(line, "   ")
		left_number, _ := strconv.Atoi(chunks[0])
		right_number, _ := strconv.Atoi(chunks[1])
		left[idx] = left_number
		right[idx] = right_number
		left_numbers[left_number] = 0
	}
	// set occurences
	for _, number := range right {
		_, ok := left_numbers[number]
		if ok {
			left_numbers[number] += 1
		}

	}
	// calculate similarity_score
	similarity_score := 0
	for number, occurrences := range left_numbers {
		if occurrences > 0 {
			result := number * occurrences
			similarity_score += result
		}

	}
	fmt.Println("similarity_score", similarity_score)
}
