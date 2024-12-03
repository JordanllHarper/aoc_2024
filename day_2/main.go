package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	input := string(b)

	answer1 := part1(input)
	fmt.Printf("Answer 1: %v\n", answer1)

	answer2 := part2(input)
	fmt.Printf("Answer 2: %v\n", answer2)
}
