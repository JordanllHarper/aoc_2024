package main

import "testing"
	const data = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPart1(t *testing.T) {
	expected := 18
	actual := part1(data)
	if expected != actual {
		t.Fatalf(`Error: Wanted %v but got %v`, expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	actual := part2(data)
	if expected != actual {
		t.Fatalf(`Error: Wanted %v but got %v`, expected, actual)
	}
}
