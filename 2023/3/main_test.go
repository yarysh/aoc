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
		{"467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..", 4361},
	} {
		r := strings.NewReader(tc.input)

		if got, want := part1(r), tc.want; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
