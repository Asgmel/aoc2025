package main

import (
	"fmt"
	"strconv"

	"github.com/asgmel/aockit/input"
)

func main() {
	taskOne("./input.txt")
}

func taskOne(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	totalJoltage := 0
	for _, bank := range puzzleInput {
		totalJoltage += findMaxJoltageInBank(bank)
	}
	fmt.Printf("Total joltage: %d\n", totalJoltage)
}

func findMaxJoltageInBank(bank string) (joltage int) {
	idx, firstJoltage := findMaxJoltage(bank[:len(bank)-1])
	_, secondJoltage := findMaxJoltage(bank[idx+1:])
	fmt.Printf("Bank %s: max joltages are %d and %d\n", bank, firstJoltage, secondJoltage)
	result, err := strconv.Atoi(string(firstJoltage) + string(secondJoltage))
	if err != nil {
		panic(err)
	}
	return result
}

func findMaxJoltage(bank string) (index int, joltage rune) {
	for idx, jolt := range bank {
		if jolt > joltage {
			joltage = jolt
			index = idx
		}
	}
	return index, joltage
}
