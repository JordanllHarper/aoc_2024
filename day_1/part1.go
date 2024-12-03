package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseSides(input string) ([]string, []string) {

	lines := strings.Split(input, "\n")

	lhs := []string{}
	rhs := []string{}

	for _, line := range lines {
		fmt.Println(line)
		split := strings.Split(line, "   ")
		if len(split) < 2 {
			// end of line
			break
		}
		left := split[0]
		right := split[1]
		lhs = append(lhs, left)
		rhs = append(rhs, right)
	}

	return lhs, rhs
}

func mapToInt(values []string) []int {
	ints := []int{}
	for _, v := range values {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("%v was not an int\n", v)
		}
		ints = append(ints, vInt)
	}
	return ints
}

func part1(input string) int {
	lhs, rhs := parseSides(input)
	lhsInt := mapToInt(lhs)
	rhsInt := mapToInt(rhs)

	{
		slices.Sort(lhsInt)
		slices.Sort(rhsInt)
	}

	differences := []int{}
	for i := 0; i < len(lhs); i++ {
		left := lhsInt[i]
		right := rhsInt[i]

		if left > right {
			differences = append(differences, left-right)
		} else if right > left {
			differences = append(differences, right-left)
		} else {
			differences = append(differences, 0)
		}
	}

	sum := func() int {
		s := 0
		for _, v := range differences {
			s += v
		}
		return s
	}()
	fmt.Printf("Answer: %v\n", sum)

	return sum
}
