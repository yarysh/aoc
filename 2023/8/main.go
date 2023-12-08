// Advent of Code 2023, Day 8
// https://adventofcode.com/2023/day/8

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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

	path := []int{}
	m := map[string][2]string{}

	s := bufio.NewScanner(f)

	s.Scan()
	for _, r := range s.Bytes() {
		if r == 'L' {
			path = append(path, 0)
		} else {
			path = append(path, 1)
		}
	}

	s.Scan()

	regex := regexp.MustCompile(`^(\w+)\s*=\s*[(](\w+)\s*[,]\s*(\w+)[)]$`)
	for s.Scan() {
		line := s.Text()
		parts := regex.FindAllStringSubmatch(line, -1)
		m[parts[0][1]] = [2]string{parts[0][2], parts[0][3]}
	}

	pos := "AAA"
	for {
		for _, d := range path {
			pos = m[pos][d]
			res++

			if pos == "ZZZ" {
				return res
			}
		}
	}
}

func part2(f io.Reader) int64 {
	path := []int{}
	m := map[string][2]string{}

	s := bufio.NewScanner(f)

	s.Scan()
	for _, r := range s.Bytes() {
		if r == 'L' {
			path = append(path, 0)
		} else {
			path = append(path, 1)
		}
	}

	s.Scan()

	regex := regexp.MustCompile(`^(\w+)\s*=\s*[(](\w+)\s*[,]\s*(\w+)[)]$`)
	for s.Scan() {
		line := s.Text()
		parts := regex.FindAllStringSubmatch(line, -1)
		m[parts[0][1]] = [2]string{parts[0][2], parts[0][3]}
	}

	var positions []string
	for k, _ := range m {
		if k[2] == 'A' {
			positions = append(positions, k)
		}
	}

	finishes := make([]int64, len(positions))

	for i := 0; i < len(positions); i++ {
		steps := 0

		for positions[i][2] != 'Z' {
			for j := 0; j < len(path); j++ {
				positions[i] = m[positions[i]][path[j]]
				steps++
			}
		}

		finishes[i] = int64(steps)
	}

	tmp := make([]int64, len(finishes))
	copy(tmp, finishes)

	for {
		matches := 1
		minI, minVal := 0, tmp[0]

		for i := 1; i < len(tmp); i++ {
			if tmp[i] < minVal {
				minI, minVal = i, tmp[i]
			} else if tmp[i] == minVal {
				matches++
			}
		}

		if matches == len(tmp) {
			return tmp[0]
		}

		tmp[minI] += finishes[minI]
	}
}
