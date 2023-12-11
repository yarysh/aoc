package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Seek(0, 0)
	fmt.Println("Part 1", solution(f, 1))

	f.Seek(0, 0)
	fmt.Println("Part 2", solution(f, 1000000-1))
}

type Item struct {
	ID int
	X  int
	Y  int
}

func solution(f *os.File, shrinkRatio int) int {
	var items []Item

	sizeX := 0
	id, y := 1, 0

	s := bufio.NewScanner(f)

	s.Scan()
	line := s.Text()
	sizeX = len(line)

	for {
		for x, ch := range line {
			if ch == '#' {
				items = append(items, Item{
					ID: id,
					X:  x,
					Y:  y,
				})
				id++
			}
		}

		if s.Scan() {
			line = s.Text()
			y++
			continue
		}

		break
	}

loopX:
	for x := 0; x < sizeX; x++ {
		for _, item := range items {
			if item.X == x {
				continue loopX
			}
		}

		for i, item := range items {
			if item.X > x {
				items[i].X += shrinkRatio
			}
		}

		x += shrinkRatio
		sizeX += shrinkRatio
	}

	sizeY := y

loopY:
	for y := 0; y < sizeY; y++ {
		for _, item := range items {
			if item.Y == y {
				continue loopY
			}
		}

		for i, item := range items {
			if item.Y > y {
				items[i].Y += shrinkRatio
			}
		}

		y += shrinkRatio
		sizeY += shrinkRatio
	}

	res := 0
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			res += max(items[i].X, items[j].X) - min(items[i].X, items[j].X) +
				max(items[i].Y, items[j].Y) - min(items[i].Y, items[j].Y)
		}
	}

	return res
}
