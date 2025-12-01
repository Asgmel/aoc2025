package main

import (
	"fmt"
	"strconv"

	"github.com/asgmel/aockit/input"
	"github.com/asgmel/aockit/utils"
)

func main() {
	taskOne("./input.txt")
	taskTwo("./input.txt")
}

func rotateDial(position int, direction string, steps int) (newPosition int, zeroCount int) {
	switch direction {
	case "L":
		newPosition = position - steps
	case "R":
		newPosition = position + steps
	}

	rotationCount := utils.AbsoluteValue(newPosition / 100)
	newPosition = newPosition % 100
	if newPosition < 0 {
		newPosition = newPosition + 100
		rotationCount++
	}

	if position == 0 && rotationCount > 0 && direction == "L" {
		return newPosition, rotationCount - 1
	}

	if newPosition == 0 && rotationCount > 0 && direction == "L" {
		return newPosition, rotationCount + 1
	}

	if newPosition == 0 && rotationCount == 0 {
		return 0, 1
	}

	return newPosition, rotationCount
}

func taskOne(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	zeroCounter := 0
	position := 50
	for _, line := range puzzleInput {
		direction := line[0:1]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err) // Since this is puzzle input, we can assume it's always valid
		}
		position, _ = rotateDial(position, direction, steps)
		// fmt.Printf("Moved %s%d to position %d\n", direction, steps, position)
		if position == 0 {
			zeroCounter++
		}
	}
	fmt.Printf("Task 1: %d\n", zeroCounter)

}

func taskTwo(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	zeroCounter := 0
	position := 50
	for _, line := range puzzleInput {
		zeroCount := 0
		direction := line[0:1]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err) // Since this is puzzle input, we can assume it's always valid
		}
		position, zeroCount = rotateDial(position, direction, steps)
		zeroCounter += zeroCount
		// fmt.Printf("Moved %s%d to position %d. Added %d zeroes, total: %d.\n", direction, steps, position, zeroCount, zeroCounter)
	}
	fmt.Printf("Task 2: %d\n", zeroCounter)
}
