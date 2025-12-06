package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/asgmel/aockit/input"
	"github.com/asgmel/aockit/utils"
)

type problem struct {
	numbers   []string
	operation byte
}

func (p problem) getNumbersAsInts() []int {
	intNumbers := []int{}
	for _, strNum := range p.numbers {
		intNum, _ := strconv.Atoi(strings.TrimSpace(strNum))
		intNumbers = append(intNumbers, intNum)
	}

	return intNumbers
}

func (p problem) getColumnNumbersAsInts() []int {
	numbers := []int{}
	for i := len(p.numbers[0]) - 1; i >= 0; i-- {
		columnNum := ""
		for _, strNum := range p.numbers {
			columnNum += string(strNum[i])
		}
		columnIntNum, _ := strconv.Atoi(strings.TrimSpace(columnNum))
		if columnIntNum > 0 {
			numbers = append(numbers, columnIntNum)
		}
	}
	return numbers
}

func main() {
	start := time.Now()
	taskOne("./input.txt")
	fmt.Printf("taskOne took %s\n", time.Since(start))

	start = time.Now()
	taskTwo("./input.txt")
	fmt.Printf("taskTwo took %s\n", time.Since(start))
}

func taskOne(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	problems := formatPuzzleInput(puzzleInput)
	sum := 0
	for _, problem := range problems {
		sum += performHorizontalCalculation(problem)
	}
	fmt.Printf("The sum of all problems is: %d\n", sum)
}

func taskTwo(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	problems := formatPuzzleInput(puzzleInput)
	sum := 0
	for _, problem := range problems {
		sum += performVerticalCalculation(problem)
	}
	fmt.Printf("The sum of all problems is: %d\n", sum)
}

func formatPuzzleInput(puzzleInput []string) []problem {
	problems := []problem{}
	runeRowIdx := len(puzzleInput) - 1
	numberStrings := make([]string, runeRowIdx)
	operation := byte(0)
	for opsIdx := 0; opsIdx < len(puzzleInput[0]); opsIdx++ {
		if puzzleInput[runeRowIdx][opsIdx] != ' ' && opsIdx != 0 {
			problems = append(problems, problem{numbers: numberStrings, operation: operation})
			numberStrings = make([]string, runeRowIdx)
			operation = byte(0)
		}
		if puzzleInput[runeRowIdx][opsIdx] != ' ' {
			operation = byte(puzzleInput[runeRowIdx][opsIdx])
		}
		for rowIdx := 0; rowIdx < runeRowIdx; rowIdx++ {
			numberStrings[rowIdx] += string(puzzleInput[rowIdx][opsIdx])
		}
	}
	problems = append(problems, problem{numbers: numberStrings, operation: operation})
	return problems
}

func performCalculation(p problem, numbers []int) int {
	switch p.operation {
	case '*':
		return MultiplyIntSlice(numbers)
	case '+':
		return utils.SumIntSlice(numbers)
	default:
		panic(fmt.Sprintf("Invalid problem operation %c", p.operation))
	}
}

func performHorizontalCalculation(p problem) int {
	return performCalculation(p, p.getNumbersAsInts())
}

func performVerticalCalculation(p problem) int {
	return performCalculation(p, p.getColumnNumbersAsInts())
}

func MultiplyIntSlice(slice []int) int {
	sum := 1
	for _, val := range slice {
		sum *= val
	}
	return sum
}
