package main

import (
	"fmt"
	"strings"
)

/*
	This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words.


It's a little unusual, though, as you don't merely need to find one instance of XMAS - you need to find all of them. Here are a few ways XMAS might appear, where irrelevant characters have been replaced with .:
*/

/*

Search for word "XMAS"
* XMAS can be forwards or backwards "SAMX"
* XMAS can be vertical forwards and backwards
* XMAS can be diagonal forwards and backwards
* XMAS can overlap other words

Letters only cared about are X,M,A,S -

NOTE: these are also the only letters in the grid

X M A S (or alterations) are a sequence of letters, so order matters

Proposed solution:

Scan the grid multiple times looking for a particular direction

We need to make sure we don't count the same word twice

We could store indexes of looked at letters for a particular direction

Sum the result of all the iterations and return as our answer.

*/

const XMAS = "XMAS"
const SAMX = "SAMX"

func scanDown(grid [][]rune, startingY, startingX int) bool {
	start := ""
	for i := startingY; i < startingY+4 && i < len(grid)-1; i++ {
		start = start + string(grid[i][startingX])
	}
	println(start)
	return start == XMAS || start == SAMX
}

func scanVertical(grid [][]rune) int {
	counter := 0
	for y, line := range grid {
		for x, c := range line {
			if c == 'X' || c == 'S' {
				success := scanDown(grid, y, x)
				if success {
					counter++
				}
			}
		}
	}
	return counter
}

// We have x
func scanForward(line []rune, startingIndex int) bool {
	start := ""

	for i := startingIndex; i < startingIndex+4 && i < len(line)-1; i++ {
		start = start + string(line[i])
	}

	return start == XMAS || start == SAMX
}
func scanHorizontal(grid [][]rune) int {
	// TODO: count words

	counter := 0
	for _, line := range grid {
		for i, c := range line {
			if c == 'X' || c == 'S' {
				success := scanForward(line, i)
				if success {
					counter++
				}
			}
		}
	}
	return counter
}

func scanDiagonal(grid [][]rune) int {
	return 0
}

func transformLine(line string) []rune {
	chars := []rune{}
	for _, char := range line {
		chars = append(chars, char)
	}
	return chars
}

func transformGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, transformLine(line))
	}
	return grid
}

func part1(input string) int {
	grid := transformGrid(input)
	horizontalOccurrences := scanHorizontal(grid)
	fmt.Println("Horizontal Occurrences: ", horizontalOccurrences)
	verticalOccurrences := scanVertical(grid)
	fmt.Println("Vertical Occurrences: ", verticalOccurrences)
	diagonalOccurrences := scanDiagonal(grid)
	return verticalOccurrences + horizontalOccurrences + diagonalOccurrences
}
