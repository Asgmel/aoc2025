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

	start = time.Now()
	taskTwo("./input.txt")
	fmt.Printf("taskTwo took %s\n", time.Since(start))
}

func taskOne(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	positions := formatPuzzleInput(puzzleInput)
	rectangles := calculateRectangles(positions)
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].area > rectangles[j].area
	})
	fmt.Printf("Task 1: %d\n", rectangles[0].area)
}

func taskTwo(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	positions := formatPuzzleInput(puzzleInput)
	rectangles := calculateRectangles(positions)
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].area > rectangles[j].area
	})

	cache := make(map[grid.Position]bool)

	for _, rect := range rectangles {
		if rectangleInsideBorder(rect, positions, cache) {
			fmt.Printf("Task 2: %d\n", rect.area)
			return
		}
	}
	fmt.Printf("No rectangle found for Task 2\n")
}

func rectangleInsideBorder(rect rectangle, positions []grid.Position, cache map[grid.Position]bool) bool {
	minX := min(rect.p1.X, rect.p2.X)
	maxX := max(rect.p1.X, rect.p2.X)
	minY := min(rect.p1.Y, rect.p2.Y)
	maxY := max(rect.p1.Y, rect.p2.Y)

	pos1 := grid.Position{X: rect.p1.X, Y: rect.p2.Y}
	pos2 := grid.Position{X: rect.p2.X, Y: rect.p1.Y}
	if !posInsideBorderCached(pos1, positions, cache) || !posInsideBorderCached(pos2, positions, cache) {
		return false
	}

	for x := minX + 1; x < maxX; x++ {
		if !posInsideBorderCached(grid.Position{X: x, Y: minY}, positions, cache) {
			return false
		}
		if !posInsideBorderCached(grid.Position{X: x, Y: maxY}, positions, cache) {
			return false
		}
	}

	for y := minY + 1; y < maxY; y++ {
		if !posInsideBorderCached(grid.Position{X: minX, Y: y}, positions, cache) {
			return false
		}
		if !posInsideBorderCached(grid.Position{X: maxX, Y: y}, positions, cache) {
			return false
		}
	}

	return true
}

func posInsideBorderCached(pos grid.Position, positions []grid.Position, cache map[grid.Position]bool) bool {
	if result, exists := cache[pos]; exists {
		return result
	}

	result := posInsideBorder(pos, positions)
	cache[pos] = result
	return result
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

func posInsideBorder(pos grid.Position, positions []grid.Position) bool {
	crossings := 0

	for i := 0; i < len(positions); i++ {
		j := (i + 1) % len(positions)

		// Check if the point is exactly on a border line
		if positions[i].X == positions[j].X && positions[i].X == pos.X {
			minY := min(positions[i].Y, positions[j].Y)
			maxY := max(positions[i].Y, positions[j].Y)
			if pos.Y >= minY && pos.Y <= maxY {
				return true
			}
		} else if positions[i].Y == positions[j].Y && positions[i].Y == pos.Y {
			minX := min(positions[i].X, positions[j].X)
			maxX := max(positions[i].X, positions[j].X)
			if pos.X >= minX && pos.X <= maxX {
				return true
			}
		}

		// If we are not on a border line, check for crossings
		// Skip horizontal border, as we will never cross them with a horizontal ray
		if positions[i].Y == positions[j].Y {
			continue
		}

		posX := positions[i].X
		minY := min(positions[i].Y, positions[j].Y)
		maxY := max(positions[i].Y, positions[j].Y)

		if pos.Y >= minY && pos.Y < maxY && posX > pos.X {
			crossings++
		}
	}
	return crossings%2 == 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
