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
		{"1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\n", 142},
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
		{"two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen\n", 281},
	} {
		r := strings.NewReader(tc.input)

		if got, want := part2(r), tc.want; got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
