package main

import (
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 8},
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 999: 1 blue, 10 green, 10 green, 4 blue", 1},
	} {
		r := strings.NewReader(tc.input)

		if got, want := part1(r), tc.want; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func Test_part2(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 2286},
	} {
		r := strings.NewReader(tc.input)

		if got, want := part2(r), tc.want; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
