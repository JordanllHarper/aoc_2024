package main

import "testing"

const data = "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9\n"

func TestPart1(t *testing.T) {
	expected := 2
	actual := part1(data)
	if expected != actual {
		t.Fatalf(`Error: Wanted %v but got %v`, expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 4
	actual := part2(data)
	if expected != actual {
		t.Fatalf(`Error: Wanted %v but got %v`, expected, actual)
	}
}
