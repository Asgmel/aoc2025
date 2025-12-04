package main

import (
	"fmt"

	"github.com/asgmel/aockit/grid"
	"github.com/asgmel/aockit/input"
)

func main() {
	taskOne("./input.txt")
	taskTwo("./input.txt")
}

func taskOne(inputPath string) {
	puzzleInput := input.ReadInputLetters(inputPath)
	markedInput := markForkliftAccessablePaper(puzzleInput)
	accessablePaperCount := countRemovablePaper(markedInput)
	println("Accessable paper count:", accessablePaperCount)
}

func taskTwo(inputPath string) {
	puzzleInput := input.ReadInputLetters(inputPath)
	count := recursiveCountMarkedPaper(puzzleInput)
	println("Accessable paper count:", count)
}

func recursiveCountMarkedPaper(matrix [][]string) int {
	count := 0
	markedMatrix := markForkliftAccessablePaper(matrix)
	count += countRemovablePaper(markedMatrix)
	newMatrix := removeMarkedPaper(markedMatrix)

	if countRemovablePaper(markedMatrix) > 0 {
		count += recursiveCountMarkedPaper(newMatrix)
	}
	return count
}

func countRemovablePaper(matrix [][]string) (count int) {
	for y, row := range matrix {
		for x := range row {
			if matrix[y][x] == "x" {
				count++
			}
		}
	}
	return count
}

func removeMarkedPaper(matrix [][]string) [][]string {
	removedMatrix := make([][]string, len(matrix))
	for i := range matrix {
		removedMatrix[i] = make([]string, len(matrix[i]))
		copy(removedMatrix[i], matrix[i])
	}
	for y, row := range matrix {
		for x := range row {
			if matrix[y][x] == "x" {
				removedMatrix[y][x] = "."
			}
		}
	}
	return removedMatrix
}

func markForkliftAccessablePaper(matrix [][]string) [][]string {
	markedMatrix := make([][]string, len(matrix))
	for i := range matrix {
		markedMatrix[i] = make([]string, len(matrix[i]))
		copy(markedMatrix[i], matrix[i])
	}
	for y, row := range matrix {
		for x := range row {
			position := grid.Position{X: x, Y: y}
			neighbourCount := countNeighboursWithPaper(position, matrix)
			if neighbourCount < 4 && matrix[y][x] == "@" {
				markedMatrix[y][x] = "x"
			}
		}
	}
	for _, row := range markedMatrix {
		fmt.Printf("%s\n", row)
	}
	return markedMatrix
}

func countNeighboursWithPaper(position grid.Position, matrix [][]string) (count int) {
	neighbours := grid.GetNeighbouringPositions(position, matrix)
	for _, neighbour := range neighbours {
		if matrix[neighbour.Y][neighbour.X] == "@" {
			count++
		}
	}
	return count
}
