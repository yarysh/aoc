// Advent of Code 2023, Day 2
// https://adventofcode.com/2023/day/2

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	res := 0

	// rgb
	limit := [3]int{12, 13, 14}

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()

		curr := [3]int{0, 0, 0}

	loop:
		for i := len(line) - 1; i >= 0; {
			var p int

			switch line[i] {
			case 'd': // red
				p = 0
				i -= 4
			case 'n': // green
				p = 1
				i -= 6
			case 'e': // blue
				p = 2
				i -= 5
			case ':':
				p = -1
				i--
			case ';':
				curr = [3]int{0, 0, 0}
				i--
				continue loop
			default:
				i--
				continue loop
			}

			n := 0
			for j := 0; i >= 0 && line[i] >= 48 && line[i] <= 57; j++ {
				n += int(line[i]-48) * int(math.Pow10(j))
				i--
			}

			switch p {
			case -1:
				res += n
				break loop
			default:
				curr[p] += n
				if curr[0] > limit[0] || curr[1] > limit[1] || curr[2] > limit[2] {
					break loop
				}
			}
		}
	}

	return res
}

func part2(f io.Reader) int {
	res := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()

		curr := [3]int{0, 0, 0}

	loop:
		for i := len(line) - 1; i >= 0; {
			var p int

			switch line[i] {
			case 'd': // red
				p = 0
				i -= 4
			case 'n': // green
				p = 1
				i -= 6
			case 'e': // blue
				p = 2
				i -= 5
			case ':':
				res += curr[0] * curr[1] * curr[2]
				break loop
			default:
				i--
				continue loop
			}

			n := 0
			for j := 0; i >= 0 && line[i] >= 48 && line[i] <= 57; j++ {
				n += int(line[i]-48) * int(math.Pow10(j))
				i--
			}

			if curr[p] < n {
				curr[p] = n
			}
		}
	}

	return res
}
