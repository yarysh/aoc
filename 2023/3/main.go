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

	f.Seek(0, 0)
	fmt.Println("Part 2", part2(f))
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

func part2(f io.Reader) int {
	result := 0

	b, _ := io.ReadAll(f)
	lines := strings.Split(string(b), "\n")

	top := make([]int, len(lines[0]))
	curr := l2arr(lines[0])
	btm := make([]int, len(lines[0]))

	for i := 0; i < len(lines)-1; i++ {
		btm = l2arr(lines[i+1])

		for j, d := range curr {
			if d != -1 {
				continue
			}

			var nums []int

			nums = append(nums, findNums(top, j)...)
			nums = append(nums, findNums(curr, j)...)
			nums = append(nums, findNums(btm, j)...)

			if len(nums) == 2 {
				result += nums[0] * nums[1]
			}
		}

		top, curr = curr[:], btm[:]
	}

	return result
}

func l2arr(line string) []int {
	result := make([]int, len(line))

	for i, ch := range line {
		if ch >= '0' && ch <= '9' {
			result[i] = int(ch - '0')
			continue
		}

		if ch == '*' {
			result[i] = -1
			continue
		}

		result[i] = -2
	}

	return result
}

func findNums(arr []int, pos int) []int {
	var result []int

	left, right := max(pos-1, 0), min(pos+1, len(arr)-1)
	for ; left > 0 && arr[left] >= 0; left-- {
	}

	for ; right < len(arr)-1 && arr[right] >= 0; right++ {
	}

	var nums [][]int
	var num []int

	for _, n := range arr[left : right+1] {
		if n < 0 {
			if len(num) > 0 {
				nums = append(nums, num)
				num = nil
			}
			continue
		}

		num = append(num, n)
	}

	if len(num) > 0 {
		nums = append(nums, num)
		num = nil
	}

	for _, n := range nums {
		result = append(result, aToD(n))
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

	for i := len(nums) - 1; i >= 0; i-- {
		result += nums[i] * int(math.Pow10(len(nums)-1-i))
	}

	return result
}
