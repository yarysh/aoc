// Advent of Code 2023, Day 9
// https://adventofcode.com/2023/day/9

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func part1(f *os.File) int {
	res := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		in := strings.Split(s.Text(), " ")

		line := make([]int, len(in))
		for i, str := range in {
			v, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			line[i] = v
		}

		for {
			zeroes := 0

			res += line[len(line)-1]

			for i := 0; i < len(line)-1; i++ {
				line[i] = line[i+1] - line[i]
				if line[i] == 0 {
					zeroes++
				}
			}

			if zeroes == len(line)-1 {
				break
			}

			line = line[:len(line)-1]
		}
	}

	return res
}

func part2(f *os.File) int {
	res := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		in := strings.Split(s.Text(), " ")

		line := make([]int, len(in))
		for i, str := range in {
			v, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			line[i] = v
		}

		var tmp []int

		for {
			zeroes := 0
			tmp = append(tmp, line[0])

			for i := 0; i < len(line)-1; i++ {
				line[i] = line[i+1] - line[i]
				if line[i] == 0 {
					zeroes++
				}
			}

			if zeroes == len(line)-1 {
				break
			}

			line = line[:len(line)-1]
		}

		currRes := tmp[len(tmp)-1]
		for i := len(tmp) - 2; i >= 0; i-- {
			currRes = tmp[i] - currRes
		}

		res += currRes
	}

	return res
}
