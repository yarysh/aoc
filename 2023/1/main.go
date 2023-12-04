// Advent of Code 2023, Day 1
// https://adventofcode.com/2023/day/1

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Seek(0, 0)
	fmt.Println("Part 1", part1(f))

	f.Seek(0, 0)
	fmt.Println("Part 2", part2(f))
}

func part1(f io.Reader) int {
	tens, ones := 0, 0

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Bytes()

		for i := 0; i < len(line); i++ {
			if line[i] >= 48 && line[i] <= 57 {
				tens += int(line[i]) - 48
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= 48 && line[i] <= 57 {
				ones += int(line[i]) - 48
				break
			}
		}
	}

	return tens*10 + ones
}

func part2(f io.Reader) int {
	m := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	tens, ones := 0, 0

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()

		for i := 0; i < len(line); i++ {
			if line[i] >= 48 && line[i] <= 57 {
				tens += int(line[i]) - 48
				break
			}

			found := false
			for digit, word := range m {
				if i+len(word) <= len(line) && line[i:i+len(word)] == word {
					tens += digit + 1
					found = true
					break
				}
			}

			if found {
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= 48 && line[i] <= 57 {
				ones += int(line[i]) - 48
				break
			}

			found := false
			for digit, word := range m {
				j := i + 1
				if j-len(word) >= 0 && line[j-len(word):j] == word {
					ones += digit + 1
					found = true
					break
				}
			}

			if found {
				break
			}
		}
	}

	return tens*10 + ones
}
