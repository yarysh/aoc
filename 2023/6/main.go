// Advent of Code 2023, Day 6
// https://adventofcode.com/2023/day/6

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 1", solution([][]int{{7, 9}, {15, 40}, {30, 200}}))
	fmt.Println("Part 2", solution([][]int{{71530, 940200}}))
}

func solution(input [][]int) int {
	result := 1

	for _, item := range input {
		time, distance := item[0], item[1]

		currResult := 0

		found := false
		for i := 1; i < time; i++ {
			if i*(time-i) > distance {
				currResult++
				found = true
				continue
			}

			if found {
				break
			}

		}

		result *= max(currResult, 1)
	}

	return result
}
