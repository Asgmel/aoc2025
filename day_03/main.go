package main

import (
	"fmt"
	"strconv"

	"github.com/asgmel/aockit/input"
)

func main() {
	taskOne("./input.txt")
	taskTwo("./input.txt")
}

func taskOne(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	totalJoltage := 0
	for _, bank := range puzzleInput {
		totalJoltage += findMaxJoltageInBank(bank, 2)
	}
	fmt.Printf("Total joltage: %d\n", totalJoltage)
}

func taskTwo(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	totalJoltage := 0
	for _, bank := range puzzleInput {
		totalJoltage += findMaxJoltageInBank(bank, 12)
	}
	fmt.Printf("Total joltage: %d\n", totalJoltage)
}

func findMaxJoltageInBank(bank string, numOfBatteries int) (joltage int) {
	joltages := ""
	currentIndex := 0
	for i := 1; i <= numOfBatteries; i++ {
		idx, maxJoltage := findMaxJoltage(bank[currentIndex : len(bank)-numOfBatteries+i])
		currentIndex += idx + 1
		joltages += string(maxJoltage)
	}
	result, err := strconv.Atoi(joltages)
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
