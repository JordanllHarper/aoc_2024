package main

import (
	"fmt"
	"strings"
)

func getIssues(report []int) int {

	isIncreasing := true

	// init case

	left, right := report[0], report[1]
	switch {
	case left > right:
		isIncreasing = false
	case left < right:
		isIncreasing = true
	case left == right:
		return 0
	}

	for i, left := range report {
		nextIndex := i + 1

		if nextIndex > len(report)-1 {
			break
		}
		right := report[nextIndex]

		if left == right {
			return i
		}

		leftRightDifference := findDifference(left, right)
		if leftRightDifference <= 0 || leftRightDifference > 3 {
			return i
		}

		switch {
		case left > right && isIncreasing:
			fallthrough
		case left < right && !isIncreasing:
			return i
		}
	}
	return -1
}

func removeAtIndex(s []int, index int) []int {
	rhsIndex := index + 1 
	if len(s) > 
	return append(s[:index], s[index+1:]...)
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	count := 0
	for _, line := range lines {
		fmt.Printf("Parsing line %v\n", line)
		report := parseReport(line)
		leftIssueIndex := getIssues(report)
		if leftIssueIndex == -1 {
			count++
			continue
		}

		rightIssueIndex := leftIssueIndex + 1

		var variantOne []int
		copy(report, variantOne)
		variantOne = removeAtIndex(variantOne, leftIssueIndex)

		var variantTwo []int
		copy(report, variantTwo)
		variantTwo = removeAtIndex(variantTwo, rightIssueIndex)

		fmt.Printf("Variant One: %v\n", variantOne)
		fmt.Printf("Variant Two: %v\n", variantTwo)

		variantOneIssues := getIssues(variantOne)
		variantTwoIssues := getIssues(variantTwo)

		if variantOneIssues == -1 || variantTwoIssues == -1 {
			count++
		} else {
			fmt.Printf("Variant One Issue Index: %v\n", variantOneIssues)
			fmt.Printf("Variant Two Issue Index: %v\n", variantTwoIssues)
		}
	}
	return count
}
