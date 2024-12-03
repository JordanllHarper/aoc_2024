package main

func countRhs(lhs int, rhs []int) int {
	count := 0
	for _, v := range rhs {
		if lhs == v {
			count += 1
		}
	}
	return count
}

func part2(input string) int {
	lhs, rhs := parseSides(input)
	lhsInt := mapToInt(lhs)
	rhsInt := mapToInt(rhs)

	similarityScore := 0
	for _, v := range lhsInt {
		matchingRhs := countRhs(v, rhsInt)
		similarityScore += (v * matchingRhs)
	}

	return similarityScore
}
