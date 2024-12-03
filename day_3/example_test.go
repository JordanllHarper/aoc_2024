package main

import "testing"

const data = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
func TestPart1(t *testing.T) {
	expected := 161
	actual := part1(data)
	if expected != actual {
		t.Fatalf(`Error: Wanted %v but got %v`, expected, actual)
	}
}

func TestPart2(t *testing.T) {
	data := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := 48
	actual := part2(data)
	if expected != actual {
		t.Fatalf(`Error: Wanted %v but got %v`, expected, actual)
	}
}
