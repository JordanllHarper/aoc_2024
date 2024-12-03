package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func filterCharacter(c rune, input string) string {
	newString := ""
	for _, v := range input {
		if v != c {
			newString = newString + string(v)
		}
	}
	return newString
}
func parseMultiplication(m string) int {
	rightOfOpen := strings.Split(m, "(")[1]
	parsed := filterCharacter(')', rightOfOpen)
	nums := strings.Split(parsed, ",")
	lhs, _ := strconv.Atoi(nums[0])
	rhs, _ := strconv.Atoi(nums[1])
	return lhs * rhs
}
func part1(input string) int {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllString(input, -1)
	fmt.Printf("Matches found: %v\n", matches)
	sum := 0
	for _, match := range matches {
		sum += parseMultiplication(match)
	}
	return sum
}
