// Advent of Code 2023, Day 5
// https://adventofcode.com/2023/day/5

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func part1(f io.Reader) int {
	var seeds []int
	m := [7][][4]int{}

	s := bufio.NewScanner(f)

	s.Scan()
	for _, seed := range parseDigits(s.Text()) {
		seeds = append(seeds, seed)
	}

	mapPos := -1
	for s.Scan() {
		line := s.Text()

		digits := true
		switch line {
		case "seed-to-soil map:":
			mapPos, digits = 0, false
		case "soil-to-fertilizer map:":
			mapPos, digits = 1, false
		case "fertilizer-to-water map:":
			mapPos, digits = 2, false
		case "water-to-light map:":
			mapPos, digits = 3, false
		case "light-to-temperature map:":
			mapPos, digits = 4, false
		case "temperature-to-humidity map:":
			mapPos, digits = 5, false
		case "humidity-to-location map:":
			mapPos, digits = 6, false
		case "":
			digits = false
		}

		if !digits {
			continue
		}

		nums := parseDigits(line)
		if len(nums) != 3 {
			panic(fmt.Sprintf("invalid line: %s => %v", line, nums))
		}

		m[mapPos] = append(m[mapPos], [4]int{
			nums[0], nums[0] + nums[2] - 1,
			nums[1], nums[1] + nums[2] - 1,
		})
	}

	minLoc := math.Inf(+1)

	for _, seed := range seeds {
		pos := seed

		for p := 0; p < 7; p++ {
			for j := 0; j < len(m[p]); j++ {
				if pos >= m[p][j][2] && pos <= m[p][j][3] {
					offset := pos - m[p][j][2]
					pos = m[p][j][0] + offset
					break
				}
			}
		}

		if float64(pos) < minLoc {
			minLoc = float64(pos)
		}
	}

	return int(minLoc)
}

func part2(f io.Reader) int {
	var seeds []int
	m := [7][][4]int{}

	s := bufio.NewScanner(f)

	s.Scan()
	for _, seed := range parseDigits(s.Text()) {
		seeds = append(seeds, seed)
	}

	if len(seeds)%2 != 0 {
		panic("odd number of seeds")
	}

	mapPos := -1
	for s.Scan() {
		line := s.Text()

		digits := true
		switch line {
		case "seed-to-soil map:":
			mapPos, digits = 0, false
		case "soil-to-fertilizer map:":
			mapPos, digits = 1, false
		case "fertilizer-to-water map:":
			mapPos, digits = 2, false
		case "water-to-light map:":
			mapPos, digits = 3, false
		case "light-to-temperature map:":
			mapPos, digits = 4, false
		case "temperature-to-humidity map:":
			mapPos, digits = 5, false
		case "humidity-to-location map:":
			mapPos, digits = 6, false
		case "":
			digits = false
		}

		if !digits {
			continue
		}

		nums := parseDigits(line)
		if len(nums) != 3 {
			panic(fmt.Sprintf("invalid line: %s => %v", line, nums))
		}

		m[mapPos] = append(m[mapPos], [4]int{
			nums[0], nums[0] + nums[2] - 1,
			nums[1], nums[1] + nums[2] - 1,
		})
	}

	minLoc := math.Inf(+1)

	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			pos := seed

			for p := 0; p < 7; p++ {
				for j := 0; j < len(m[p]); j++ {
					if pos >= m[p][j][2] && pos <= m[p][j][3] {
						offset := pos - m[p][j][2]
						pos = m[p][j][0] + offset
						break
					}
				}
			}

			if float64(pos) < minLoc {
				minLoc = float64(pos)
			}
		}
	}

	return int(minLoc)
}

func parseDigits(line string) []int {
	var result []int
	for _, numStr := range strings.Split(strings.ReplaceAll(line, "seeds: ", ""), " ") {
		num, _ := strconv.Atoi(numStr)
		result = append(result, num)
	}
	return result
}
