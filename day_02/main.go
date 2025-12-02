package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/asgmel/aockit/input"
	"github.com/asgmel/aockit/utils"
)

type IdRange struct {
	start int
	end   int
}

func main() {
	taskOne("./input.txt")
	taskTwo("./input.txt")
}

func taskOne(inputPath string) {
	puzzleInput := input.ReadInputString(inputPath)
	ids := getIds(puzzleInput)
	invalidIds := []int{}
	for _, id := range ids {
		invalidIds = append(invalidIds, hasDoubleSequence(id)...)
	}
	fmt.Printf("Task 1: %d\n", utils.SumIntSlice(invalidIds))

}

func taskTwo(inputPath string) {
	puzzleInput := input.ReadInputString(inputPath)
	ids := getIds(puzzleInput)
	invalidIds := []int{}
	for _, id := range ids {
		invalidIds = append(invalidIds, hasNSequence(id)...)
	}
	fmt.Printf("Task 2: %d\n", utils.SumIntSlice(invalidIds))
}

func hasDoubleSequence(idRange IdRange) []int {
	invalidIds := []int{}
	for i := idRange.start; i <= idRange.end; i++ {
		idString := strconv.Itoa(i)
		firstHalf := idString[0 : len(idString)/2]
		secondHalf := idString[len(idString)/2:]
		if firstHalf == secondHalf {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func hasNSequence(idRange IdRange) []int {
	invalidIds := []int{}
idLoop:
	for i := idRange.start; i <= idRange.end; i++ {
		idString := strconv.Itoa(i)
	comparisonLoop:
		for n := 1; n <= len(idString)/2; n++ {
			if len(idString)%n != 0 {
				continue
			}
			substrings := splitByN(idString, n)
			for j := 0; j < len(substrings)-1; j++ {
				if substrings[j] != substrings[j+1] {
					continue comparisonLoop
				}
			}
			invalidIds = append(invalidIds, i)
			continue idLoop
		}
	}
	return invalidIds
}

func splitByN(s string, n int) []string {
	var result []string
	for i := 0; i <= len(s)-n; i += n {
		result = append(result, s[i:i+n])
	}
	return result
}

func getIds(puzzleInput string) []IdRange {
	// given that the only input is the puzzle input, we do not need to validate it
	stringIdRanges := strings.Split(puzzleInput, ",")
	idRanges := make([]IdRange, 0, len(stringIdRanges))

	for _, stringIdRange := range stringIdRanges {
		stringIdList := strings.Split(stringIdRange, "-")
		start, err := strconv.Atoi(stringIdList[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(stringIdList[1])
		if err != nil {
			panic(err)
		}
		idRange := IdRange{
			start: start,
			end:   end,
		}
		idRanges = append(idRanges, idRange)
	}
	return idRanges
}
