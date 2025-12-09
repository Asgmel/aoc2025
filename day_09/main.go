package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/asgmel/aockit/grid"
	"github.com/asgmel/aockit/input"
	"github.com/asgmel/aockit/utils"
)

type rectangle struct {
	p1   grid.Position
	p2   grid.Position
	area int
}

func main() {
	start := time.Now()
	taskOne("./input.txt")
	fmt.Printf("taskOne took %s\n", time.Since(start))

	// start = time.Now()
	// taskTwo("./input.txt")
	// fmt.Printf("taskTwo took %s\n", time.Since(start))
}

func taskOne(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	positions := formatPuzzleInput(puzzleInput)
	rectangles := calculateRectangles(positions)
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].area > rectangles[j].area
	})
	fmt.Printf("Largest rectangle area: %d\n", rectangles[0].area)
}

func formatPuzzleInput(input []string) []grid.Position {
	positions := make([]grid.Position, 0, len(input))
	for _, line := range input {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		positions = append(positions, grid.Position{
			X: x,
			Y: y,
		})
	}
	return positions
}

func calculateRectangles(positions []grid.Position) []rectangle {
	rectangles := []rectangle{}
	for p1 := 0; p1 < len(positions); p1++ {
		for p2 := p1 + 1; p2 < len(positions); p2++ {
			rectangles = append(rectangles, rectangle{
				p1:   positions[p1],
				p2:   positions[p2],
				area: (utils.AbsoluteValue(positions[p1].X-positions[p2].X) + 1) * (utils.AbsoluteValue(positions[p1].Y-positions[p2].Y) + 1),
			})
		}
	}
	return rectangles
}
