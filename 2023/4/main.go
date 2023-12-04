// Advent of Code 2023, Day 4
// https://adventofcode.com/2023/day/4

package main

import (
	"bufio"
	"bytes"
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
	result := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Bytes()

		cards := map[int]struct{}{}

		i := len(line) - 1
		var num []int

		for ; i >= 0 && line[i] != '|'; i-- {
			if line[i] == ' ' && len(num) > 0 {
				cards[a2D(num)] = struct{}{}
				num = nil
			} else if line[i] >= '0' && line[i] <= '9' {
				num = append(num, int(line[i]-'0'))
			}
		}

		num = nil
		wins := float64(-1)

		for ; i >= 0 && line[i] != ':'; i-- {
			if line[i] == ' ' && len(num) > 0 {
				if _, ok := cards[a2D(num)]; ok {
					wins++
				}
				num = nil
			} else if line[i] >= '0' && line[i] <= '9' {
				num = append(num, int(line[i]-'0'))
			}
		}

		if wins != -1 {
			result += int(math.Pow(2, wins))
		}
	}

	return result
}

func part2(f io.Reader) int {
	b, _ := io.ReadAll(f)
	lines := bytes.Split(b, []byte("\n"))

	scores := make([]int, len(lines))
	for i := 0; i < len(scores); i++ {
		scores[i] = 1
	}

	for n := 0; n < len(lines); n++ {
		line := lines[n]

		cards := map[int]struct{}{}

		i := len(line) - 1
		var num []int

		for ; i >= 0 && line[i] != '|'; i-- {
			if line[i] == ' ' && len(num) > 0 {
				cards[a2D(num)] = struct{}{}
				num = nil
			} else if line[i] >= '0' && line[i] <= '9' {
				num = append(num, int(line[i]-'0'))
			}
		}

		num = nil
		wins := 0

		for ; i >= 0 && line[i] != ':'; i-- {
			if line[i] == ' ' && len(num) > 0 {
				if _, ok := cards[a2D(num)]; ok {
					wins++
				}
				num = nil
			} else if line[i] >= '0' && line[i] <= '9' {
				num = append(num, int(line[i]-'0'))
			}
		}

		for j := max(scores[n], 1); j > 0; j-- {
			k := n + 1
			for ; k <= wins+n; k++ {
				scores[k] += 1
			}
		}
	}

	result := 0
	for _, score := range scores {
		result += score
	}

	return result
}

func a2D(num []int) int {
	result := 0
	for i := 0; i < len(num); i++ {
		result += num[i] * int(math.Pow10(i))
	}
	return result
}
