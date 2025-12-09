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
	border := generateBorderPositions(positions)
	rectangles := calculateRectangles(positions)
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].area > rectangles[j].area
	})

	// Cache for inside/outside checks
	cache := make(map[grid.Position]bool)

	for idx, rect := range rectangles {
		fmt.Printf("Checking rectangle %d with area %d\n", idx, rect.area)
		if rectangleFullyInsideBorderCached(rect, positions, border, cache) {
			fmt.Printf("Task 2: %d\n", rect.area)
			return
		}
	}
	fmt.Printf("No rectangle found for Task 2\n")
}

func rectangleFullyInsideBorderCached(rect rectangle, positions []grid.Position, border map[grid.Position]struct{}, cache map[grid.Position]bool) bool {
	minX := min(rect.p1.X, rect.p2.X)
	maxX := max(rect.p1.X, rect.p2.X)
	minY := min(rect.p1.Y, rect.p2.Y)
	maxY := max(rect.p1.Y, rect.p2.Y)

	pos1 := grid.Position{X: rect.p1.X, Y: rect.p2.Y}
	pos2 := grid.Position{X: rect.p2.X, Y: rect.p1.Y}
	if !posInsideBorderCached(pos1, positions, border, cache) || !posInsideBorderCached(pos2, positions, border, cache) {
		return false
	}

	for x := minX + 1; x < maxX; x++ {
		if !posInsideBorderCached(grid.Position{X: x, Y: minY}, positions, border, cache) {
			return false
		}
		if !posInsideBorderCached(grid.Position{X: x, Y: maxY}, positions, border, cache) {
			return false
		}
	}

	for y := minY + 1; y < maxY; y++ {
		if !posInsideBorderCached(grid.Position{X: minX, Y: y}, positions, border, cache) {
			return false
		}
		if !posInsideBorderCached(grid.Position{X: maxX, Y: y}, positions, border, cache) {
			return false
		}
	}

	return true
}

func posInsideBorderCached(pos grid.Position, positions []grid.Position, border map[grid.Position]struct{}, cache map[grid.Position]bool) bool {
	if result, exists := cache[pos]; exists {
		return result
	}

	result := posInsideBorder(pos, positions, border)
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

func generateBorderPositions(p []grid.Position) map[grid.Position]struct{} {
	border := make(map[grid.Position]struct{})
	for i := 0; i < len(p); i++ {
		j := (i + 1) % len(p)
		border[p[i]] = struct{}{}

		if p[i].X == p[j].X {
			for k := min(p[i].Y, p[j].Y) + 1; k < max(p[i].Y, p[j].Y); k++ {
				border[grid.Position{X: p[i].X, Y: k}] = struct{}{}
			}
		} else if p[i].Y == p[j].Y {
			for k := min(p[i].X, p[j].X) + 1; k < max(p[i].X, p[j].X); k++ {
				border[grid.Position{X: k, Y: p[i].Y}] = struct{}{}
			}
		}
	}
	return border
}

func posInsideBorder(pos grid.Position, positions []grid.Position, border map[grid.Position]struct{}) bool {
	if _, onBorder := border[pos]; onBorder {
		return true
	}

	crossings := 0

	for i := 0; i < len(positions); i++ {
		j := (i + 1) % len(positions) // Wrap around to the first position

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
