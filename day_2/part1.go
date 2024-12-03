package main

import (
	"strconv"
	"strings"
)

func parseReport(report string) []int {
	levels := strings.Split(report, " ")
	levelsToInt := func() []int {
		l := []int{}
		for _, level := range levels {
			lInt, _ := strconv.Atoi(level)
			l = append(l, lInt)
		}
		return l
	}()
	return levelsToInt
}

func windows(nums []int) [][]int {
	windows := [][]int{}

	for i := range nums {
		if i == len(nums)-1 {
			break
		}
		windows = append(windows, []int{nums[i], nums[i+1]})
	}
	return windows
}

func findDifference(numOne int, numTwo int) int {
	if numOne > numTwo {
		return numOne - numTwo
	} else if numOne < numTwo {
		return numTwo - numOne
	} else {
		return 0
	}
}

func isSafe(report []int) bool {
	reportWindows := windows(report )
	isIncreasing := true

	first := reportWindows[0]
	left := first[0]
	right := first[1]

	// handle initialization - which direction do we start and is it good?
	if left > right {
		isIncreasing = false
	} else if left < right {
		isIncreasing = true
	} else {
		return false
	}

	// // guard for asserting number difference
	leftRightDifference := findDifference(left, right)
	if leftRightDifference <= 0 || leftRightDifference > 3 {
		return false
	}

	for _, window := range reportWindows[1:] {
		left := window[0]
		right := window[1]

		// guard for asserting number order
		switch {
		case left > right && isIncreasing:
			return false
		case left < right && !isIncreasing:
			return false
		case left == right:
			return false
		}

		// guard for asserting number difference
		leftRightDifference := findDifference(left, right)
		if leftRightDifference <= 0 || leftRightDifference > 3 {
			return false
		}
	}

	return true
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	count := 0
	for _, line := range lines {
		report := parseReport(line)
		if safe := isSafe(report); safe {
			count++
		}
	}
	return count
}
