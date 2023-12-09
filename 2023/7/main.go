// Advent of Code 2023, Day 7
// https://adventofcode.com/2023/day/7

package main

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
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

type Hand struct {
	CardsStr string
	Cards    []int
	Bid      int
	Score    int
}

func part1(f io.Reader) int {
	var deck = map[rune]int{'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12}

	h2a := func(hand string) []int {
		r := make([]int, len(deck))

		for _, ch := range hand {
			r[deck[ch]]++
		}

		return r
	}

	var hands []Hand

	s := bufio.NewScanner(f)

	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		hand := parts[0]
		cards := h2a(hand)

		hands = append(hands, Hand{
			CardsStr: parts[0],
			Cards:    cards,
			Bid:      bid,
			Score:    score(cards, false),
		})
	}

	slices.SortStableFunc(hands, func(i, j Hand) int {
		if n := cmp.Compare(i.Score, j.Score); n != 0 {
			return n
		}

		for k := 0; k < len(i.Cards); k++ {
			if n := cmp.Compare(deck[rune(i.CardsStr[k])], deck[rune(j.CardsStr[k])]); n != 0 {
				return n
			}
		}

		return 0
	})

	res := 0
	for n, hand := range hands {
		res += (n + 1) * hand.Bid
	}

	return res
}

func part2(f io.Reader) int {
	var deck = map[rune]int{'J': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'T': 9, 'Q': 10, 'K': 11, 'A': 12}

	h2a := func(hand string) []int {
		r := make([]int, len(deck))

		for _, ch := range hand {
			r[deck[ch]]++
		}

		return r
	}

	var hands []Hand

	s := bufio.NewScanner(f)

	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		hand := parts[0]
		cards := h2a(hand)

		hands = append(hands, Hand{
			CardsStr: parts[0],
			Cards:    cards,
			Bid:      bid,
			Score:    score(cards, true),
		})
	}

	slices.SortStableFunc(hands, func(i, j Hand) int {
		if n := cmp.Compare(i.Score, j.Score); n != 0 {
			return n
		}

		for k := 0; k < len(i.Cards); k++ {
			if n := cmp.Compare(deck[rune(i.CardsStr[k])], deck[rune(j.CardsStr[k])]); n != 0 {
				return n
			}
		}

		return 0
	})

	res := 0
	for n, hand := range hands {
		res += (n + 1) * hand.Bid
	}

	return res
}

func score(hand []int, jokerMode bool) int {
	res := 0

	if j := hand[0]; jokerMode && j > 0 {
		maxIndex, maxValue := 0, 0
		for i := 1; i < len(hand); i++ {
			if hand[i] >= maxValue {
				maxIndex, maxValue = i, hand[i]
			}
		}

		hand[maxIndex] = min(hand[maxIndex]+j, 5)
		hand[0] = 0
	}

	for _, n := range hand {
		if n == 0 {
			continue
		}

		if n >= 2 {
			res += int(math.Pow(2, float64(n))) - n
		}
	}

	return max(res, 1)
}
