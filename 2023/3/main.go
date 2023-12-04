// Advent of Code 2023, Day 1
// https://adventofcode.com/2023/day/3

package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Seek(0, 0)
	fmt.Println("Part 1", part1(f))
}

func part1(f io.Reader) int {
	check := func(i int, line string) bool {
		return !unicode.IsDigit(rune(line[i]))
	}

	result := 0

	b, _ := io.ReadAll(f)
	lines := strings.Split(string(b), "\n")

	lineLen := len(lines[0])

	prevSymbols := make([]bool, lineLen)
	nextSymbols := make([]bool, lineLen)

	nextNums := make([]int, lineLen)

	currNums, currSymbols := parseLine(lines[0], check)

	for i := 1; i <= len(lines); i++ {
		nextNums, nextSymbols = parseLine(lines[min(i, len(lines)-1)], check)

		var found bool
		var num []int

		for pos := 0; pos < len(currNums); pos++ {
			if currNums[pos] == -1 {
				if !found {
					num = nil
					continue
				}

				result += aToD(num)

				found, num = false, nil
				continue
			}

			num = append(num, currNums[pos])

			prevPos, nextPos := max(pos-1, 0), min(pos+1, lineLen-1)
			found = found ||
				prevSymbols[prevPos] || prevSymbols[pos] || prevSymbols[nextPos] ||
				currSymbols[prevPos] || currSymbols[nextPos] ||
				nextSymbols[prevPos] || nextSymbols[pos] || nextSymbols[nextPos]

			if found && pos == len(currNums)-1 {
				result += aToD(num)
			}

		}

		prevSymbols = currSymbols[:]
		currNums, currSymbols = nextNums[:], nextSymbols[:]
	}

	return result
}

func parseLine(line string, check func(i int, line string) bool) ([]int, []bool) {
	var nums = make([]int, len(line))
	for i, _ := range nums {
		nums[i] = -1
	}

	var symbols = make([]bool, len(line))

	for i := 0; i < len(line); i++ {
		if line[i] == '.' {
			continue
		}

		if check(i, line) {
			symbols[i] = true
			continue
		}

		for j := i; j < len(line) && unicode.IsDigit(rune(line[j])); j++ {
			nums[j] = int(line[j] - 48)
		}
	}

	return nums, symbols
}

func aToD(nums []int) int {
	result := 0

	start := 0
	for i, n := range nums {
		if n != -1 {
			start = i
			break
		}
	}

	end := len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != -1 {
			end = i
			break
		}
	}

	filteredNums := nums[start : end+1]

	for k := len(filteredNums) - 1; k >= 0; k-- {
		if filteredNums[k] == -1 {
			return 0
		}

		result += filteredNums[k] * int(math.Pow10(len(filteredNums)-1-k))
	}

	return result
}
