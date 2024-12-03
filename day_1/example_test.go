package main

import "testing"

func TestPart1(t *testing.T) {
	data := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 11
	actual := part1(data)
	if expected != actual {
		t.Fatalf(`Error: Wanted %v but got %v`, expected, actual)
	}
}

func TestPart2(t *testing.T) {
	data := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 31
	actual := part2(data)
	if expected != actual {
		t.Fatalf(`Error: Wanted %v but got %v`, expected, actual)
	}
}
