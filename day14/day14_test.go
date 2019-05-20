package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	input := []byte{3, 7}
	recipes := 10

	tests := []struct {
		name   string
		before int
		want   []byte
	}{
		{"9", 9, []byte{5, 1, 5, 8, 9, 1, 6, 7, 7, 9}},
		{"5", 5, []byte{0, 1, 2, 4, 5, 1, 5, 8, 9, 1}},
		{"18", 18, []byte{9, 2, 5, 1, 0, 7, 1, 0, 8, 5}},
		{"2018", 2018, []byte{5, 9, 4, 1, 4, 2, 9, 8, 8, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(input, tt.before, recipes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	input := []byte{3, 7}

	tests := []struct {
		name  string
		score []byte
		want  int
	}{
		{"51589", []byte{5, 1, 5, 8, 9}, 9},
		{"01245", []byte{0, 1, 2, 4, 5}, 5},
		{"92510", []byte{9, 2, 5, 1, 0}, 18},
		{"59414", []byte{5, 9, 4, 1, 4}, 2018},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(input, tt.score); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
