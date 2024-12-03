package main

import (
	"strings"
)

func removeAtIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func leftRightDifferenceInBound(left, right int) bool {
	leftRightDifference := findPositiveDifference(left, right)
	return leftRightDifference > 0 && leftRightDifference < 4
}

func checkLeftRight(left, right int, isIncreasing bool) bool {

	if !leftRightDifferenceInBound(left, right) {
		return false
	}

	switch {
	case left > right && isIncreasing:
		return false
	case left < right && !isIncreasing:
		return false
	case left == right:
		return false
	}

	return true
}

func getNextIssueIndex(report []int, isIncreasing bool) int {
	{

		left := report[0]
		right := report[1]

		if !leftRightDifferenceInBound(left, right) {
			return 0
		}

	}

	for i, left := range report {
		rightIndex := i + 1
		if outOfRange(report, rightIndex) {
			break // out of range, this has already been checked
		}
		right := report[rightIndex]
		success := checkLeftRight(left, right, isIncreasing)
		if !success {
			return i // return the lhs index
		}
	}
	return -1
}

func outOfRange(s []int, index int) bool {
	return index > len(s)-1
}

func testRemoveAtIndex(s []int, indexToRemove int, isIncreasing bool) bool {
	variant := removeAtIndex(s, indexToRemove)
	nextIssue := getNextIssueIndex(variant, isIncreasing)
	return nextIssue == -1
}

func copySlice(s []int) []int {
	cpy := make([]int, len(s))
	copy(cpy, s)
	return cpy
}

func considerInDirection(report []int, isIncreasing bool) bool {

	lhsIssueIndex := getNextIssueIndex(report, isIncreasing)

	// no issues
	if lhsIssueIndex == -1 {
		return true
	}

	rhsIssueIndex := lhsIssueIndex + 1

	lhsTestCpy, rhsTestCpy := copySlice(report), copySlice(report)

	return testRemoveAtIndex(lhsTestCpy, lhsIssueIndex, isIncreasing) || testRemoveAtIndex(rhsTestCpy, rhsIssueIndex, isIncreasing)
}
func part2(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	count := 0
	for _, line := range lines {
		report := parseReport(line)
		if considerInDirection(report, true) || considerInDirection(report, false) {
			count++
		}
	}
	return count
}
