// Advent of Code 2023, Day 10
// https://adventofcode.com/2023/day/10

package main

import (
	"fmt"
	"io"
	"os"
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
}

func part1(f *os.File) int {
	puzzle, _ := io.ReadAll(f)
	lines := strings.Split(string(puzzle), "\n")

	line1 := strings.Split(lines[0], "")

	emptyRow := make([]string, len(line1)+2)
	for i := 0; i < len(emptyRow); i++ {
		emptyRow[i] = "."
	}

	rows := make([][]string, len(lines)+2)
	rows[0] = emptyRow
	for i := 1; i < len(lines)+1; i++ {
		rows[i] = strings.Split("."+lines[i-1]+".", "")
	}
	rows[len(rows)-1] = emptyRow

	start := findStart(rows)
	next := findNext(start, [2]int{-1, -1}, rows)
	if len(next) != 2 {
		panic("couldn't find start")
	}

	steps := 1
	prevCurr1, prevCurr2 := start, start
	curr1, curr2 := next[0], next[1]

	for curr1 != curr2 {
		curr1Next := findNext(curr1, prevCurr1, rows)
		if len(curr1Next) != 1 {
			panic(fmt.Sprintf("curr1: more than 2 possible steps %v", curr1Next))
		}
		prevCurr1, curr1 = curr1, curr1Next[0]

		curr2Next := findNext(curr2, prevCurr2, rows)
		if len(curr2Next) != 1 {
			panic(fmt.Sprintf("curr2: more than 2 possible steps %v", curr2Next))
		}
		prevCurr2, curr2 = curr2, curr2Next[0]

		steps++
	}

	return steps
}

func findStart(rows [][]string) [2]int {
	for i, row := range rows {
		for j, pipe := range row {
			if pipe == "S" {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}

func findNext(pos [2]int, prev [2]int, rows [][]string) [][2]int {
	var res [][2]int

	add := func(item [2]int) {
		if item != prev {
			res = append(res, item)
		}
	}

	top := [2]int{pos[0] - 1, pos[1]}
	right := [2]int{pos[0], pos[1] + 1}
	bottom := [2]int{pos[0] + 1, pos[1]}
	left := [2]int{pos[0], pos[1] - 1}

	m := map[string][4]string{
		"|": {"|7F", "", "|LJ", ""},
		"-": {"", "-J7", "", "-LF"},
		"L": {"|7F", "|-7J", "", ""},
		"J": {"|7F", "", "", "|-LF"},
		"7": {"", "", "|LJ", "|-LF"},
		"F": {"", "|-7J", "|LJ", ""},
		"S": {"|7F", "-J7", "|LJ", "-LF"},
	}

	ch := rows[pos[0]][pos[1]]

	topCh := rows[top[0]][top[1]]
	if strings.Index(m[ch][0], topCh) != -1 {
		add(top)
	}

	rightCh := rows[right[0]][right[1]]
	if strings.Index(m[ch][1], rightCh) != -1 {
		add(right)
	}

	bottomCh := rows[bottom[0]][bottom[1]]
	if strings.Index(m[ch][2], bottomCh) != -1 {
		add(bottom)
	}

	leftCh := rows[left[0]][left[1]]
	if strings.Index(m[ch][3], leftCh) != -1 {
		add(left)
	}

	return res
}
