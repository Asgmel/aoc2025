package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/asgmel/aockit/input"
)

type freshRange struct {
	start int
	end   int
}

func (r freshRange) ElementsInRange() int {
	return r.end - r.start + 1
}

func (r freshRange) InRange(val int) bool {
	return val >= r.start && val <= r.end
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
	freshIds, ids := formatPuzzleInput(puzzleInput)
	count := countFreshIds(freshIds, ids)
	fmt.Printf("Fresh id count: %d\n", count)
}

func countFreshIds(freshIds []freshRange, ids []int) (count int) {
	for _, id := range ids {
		if checkIdFreshness(freshIds, id) {
			count++
		}
	}
	return count
}

func checkIdFreshness(freshIds []freshRange, id int) bool {
	for _, r := range freshIds {
		if r.InRange(id) {
			return true
		}
	}
	return false
}

func formatPuzzleInput(input []string) (freshIds []freshRange, ids []int) {
	readingIds := false
	for _, line := range input {
		if line == "" {
			readingIds = true
			continue
		}
		if readingIds {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		} else {
			splitLine := strings.Split(line, "-")
			start, _ := strconv.Atoi(splitLine[0])
			end, _ := strconv.Atoi(splitLine[1])
			freshIds = append(freshIds, freshRange{start: start, end: end})
		}
	}
	return freshIds, ids
}

func taskTwo(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	freshRanges, _ := formatPuzzleInput(puzzleInput)
	uniqueRanges := getUniqueRanges(freshRanges)
	totalFreshCount := 0
	for _, r := range uniqueRanges {
		totalFreshCount += r.ElementsInRange()
	}
	fmt.Printf("Total fresh id count: %d\n", totalFreshCount)
}

func rangeOverlapsStart(r1, r2 freshRange) bool {
	return r1.start <= r2.start && r1.end >= r2.start
}

func rangeOverlapsEnd(r1, r2 freshRange) bool {
	return r1.start <= r2.end && r1.end >= r2.end
}

func rangeContains(r1, r2 freshRange) bool {
	return r1.start <= r2.start && r1.end >= r2.end
}

func recursiveGetUniqueRanges(currentRange freshRange, uniqueRanges []freshRange) []freshRange {
	for i := 0; i < len(uniqueRanges); i++ {
		uniqueRange := uniqueRanges[i]
		if rangeContains(uniqueRange, currentRange) {
			return uniqueRanges
		} else if rangeContains(currentRange, uniqueRange) {
			uniqueRanges = append(uniqueRanges[:i], uniqueRanges[i+1:]...)
			return recursiveGetUniqueRanges(currentRange, uniqueRanges)
		} else if rangeOverlapsStart(currentRange, uniqueRange) {
			currentRange.end = uniqueRange.start - 1
			return recursiveGetUniqueRanges(currentRange, uniqueRanges)
		} else if rangeOverlapsEnd(currentRange, uniqueRange) {
			currentRange.start = uniqueRange.end + 1
			return recursiveGetUniqueRanges(currentRange, uniqueRanges)
		}
	}
	uniqueRanges = append(uniqueRanges, currentRange)
	return uniqueRanges
}

func getUniqueRanges(ranges []freshRange) (uniqueRanges []freshRange) {
	for _, r := range ranges {
		uniqueRanges = recursiveGetUniqueRanges(r, uniqueRanges)
	}
	return uniqueRanges
}
