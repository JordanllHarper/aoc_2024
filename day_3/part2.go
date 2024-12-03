package main

import (
	"fmt"
	"regexp"
)

func part2(input string) int {
	r := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)
	matches := r.FindAllString(input, -1)
	fmt.Printf("Matches found: %v\n", matches)
	sum := 0
	enableSum := true
	for _, match := range matches {
		if m, _ := regexp.Match(`don't\(\)`, []byte(match)); m {
			enableSum = false
			continue
		}
		if m, _ := regexp.Match(`do\(\)`, []byte(match)); m {
			enableSum = true
			continue
		}

		if enableSum {
			sum += parseMultiplication(match)
		}
	}
	return sum
}
