package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func day1_1() {
	result := GetPuzzleInput("inputs/1.txt")

	left := make([]int, len(result))
	right := make([]int, len(result))
	for idx, line := range result {
		chunks := strings.Split(line, "   ")
		left_number, _ := strconv.Atoi(chunks[0])
		right_number, _ := strconv.Atoi(chunks[1])
		left[idx] = left_number
		right[idx] = right_number

	}
	slices.Sort(left)
	slices.Sort(right)

	total_distance := 0
	for idx, left_number := range left {
		right_number := right[idx]
		fmt.Println(left_number, right_number)
		distance := (right_number - left_number)
		if distance < 0 {
			distance = distance * -1
		}
		total_distance += distance

	}
	fmt.Println("distance", total_distance)
}
