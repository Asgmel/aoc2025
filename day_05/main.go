package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/asgmel/aockit/input"
)

type freshRange struct {
	start int
	end   int
}

func (r freshRange) InRange(val int) bool {
	return val >= r.start && val <= r.end
}

func main() {
	taskOne("./input.txt")
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
