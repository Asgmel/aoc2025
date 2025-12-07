package main

import (
	"fmt"
	"time"

	"github.com/asgmel/aockit/input"
	"github.com/asgmel/aockit/utils"
)

func main() {
	start := time.Now()
	taskOne("./input.txt")
	fmt.Printf("taskOne took %s\n", time.Since(start))

	start = time.Now()
	taskTwo("./input.txt")
	fmt.Printf("taskTwo took %s\n", time.Since(start))
}

func taskOne(inputPath string) {
	puzzleInput := input.ReadInputLetters(inputPath)
	count := loopTachyonManifold(puzzleInput)
	fmt.Printf("Task 1: %d\n", count)
}

func taskTwo(inputPath string) {
	puzzleInput := input.ReadInputLetters(inputPath)
	cache := make(map[string]int)
	count := traverseTachyonManifoldWithCache(puzzleInput, 1, getInitialBeamIndex(puzzleInput[0]), cache)
	fmt.Printf("Task 2: %d\n", count)
}

func loopTachyonManifold(puzzleInput [][]string) int {
	beamIndexes := []int{getInitialBeamIndex(puzzleInput[0])}
	splitCounter := 0
	for _, row := range puzzleInput {
		newBeamIndexes, count := placeBeams(row, beamIndexes)
		beamIndexes = newBeamIndexes
		splitCounter += count
	}
	return splitCounter
}

func traverseTachyonManifoldWithCache(puzzleInput [][]string, rowIdx int, beamIdx int, cache map[string]int) int {
	cacheKey := fmt.Sprintf("%d,%d", rowIdx, beamIdx)

	if result, exists := cache[cacheKey]; exists {
		return result
	}

	if rowIdx == len(puzzleInput)-1 {
		cache[cacheKey] = 1
		return 1
	}

	beamIndexes, _ := placeBeams(puzzleInput[rowIdx], []int{beamIdx})
	count := 0
	for _, idx := range beamIndexes {
		count += traverseTachyonManifoldWithCache(puzzleInput, rowIdx+1, idx, cache)
	}

	cache[cacheKey] = count
	return count
}

func getInitialBeamIndex(row []string) int {
	for idx := range row {
		if row[idx] == "S" {
			return idx
		}
	}
	return 0
}

func placeBeams(row []string, beamIndexes []int) (newBeamIndexes []int, splitCount int) {
	for _, beamIdx := range beamIndexes {
		if row[beamIdx] == "^" {
			splitCount++
			newBeamIndexes = append(newBeamIndexes, beamIdx-1)
			newBeamIndexes = append(newBeamIndexes, beamIdx+1)
		} else {
			newBeamIndexes = append(newBeamIndexes, beamIdx)
		}

	}
	newBeamIndexes = utils.FilterDuplicates(newBeamIndexes)
	return newBeamIndexes, splitCount
}
