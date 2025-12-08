package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/asgmel/aockit/input"
	"github.com/asgmel/aockit/utils"
)

type junctionBox struct {
	x int
	y int
	z int
}

type junctionConnection struct {
	from     junctionBox
	to       junctionBox
	distance float64
}

type circuit struct {
	junctionBoxes []junctionBox
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
	junctionBoxes := createJunctionBoxesFromInput(puzzleInput)
	circuits := createInitialCircuits(junctionBoxes)
	connections := calculateConnections(junctionBoxes)
	sortedConnections := sortConnectionsByDistanceAscending(connections)
	for _, connection := range sortedConnections[:1000] {
		circuits = connectCircuits(connection, circuits)
	}
	uniqueCircuits := filterDuplicateCircuits(circuits)
	sort.Slice(uniqueCircuits, func(i, j int) bool {
		return len(uniqueCircuits[i].junctionBoxes) > len(uniqueCircuits[j].junctionBoxes)
	})
	sum := len(uniqueCircuits[0].junctionBoxes) * len(uniqueCircuits[1].junctionBoxes) * len(uniqueCircuits[2].junctionBoxes)
	fmt.Printf("Task 1: %d\n", sum)
}

func taskTwo(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	junctionBoxes := createJunctionBoxesFromInput(puzzleInput)
	circuits := createInitialCircuits(junctionBoxes)
	connections := calculateConnections(junctionBoxes)
	sortedConnections := sortConnectionsByDistanceAscending(connections)
	result := 0
	for _, connection := range sortedConnections {
		circuits = connectCircuits(connection, circuits)
		if len(filterDuplicateCircuits(circuits)) == 1 {
			result = connection.to.x * connection.from.x
			break
		}
	}
	fmt.Printf("Task 2: %d\n", result)
}

func filterDuplicateCircuits(circuits map[junctionBox]*circuit) []circuit {
	uniqueCircuits := make(map[*circuit]bool)

	for _, circuitPtr := range circuits {
		uniqueCircuits[circuitPtr] = true
	}

	result := make([]circuit, 0, len(uniqueCircuits))
	for circuitPtr := range uniqueCircuits {
		result = append(result, *circuitPtr)
	}

	return result
}

func createInitialCircuits(jb []junctionBox) map[junctionBox]*circuit {
	circuits := make(map[junctionBox]*circuit)
	for _, box := range jb {
		circuits[box] = &circuit{
			junctionBoxes: []junctionBox{box},
		}
	}
	return circuits
}

func connectCircuits(connection junctionConnection, circuits map[junctionBox]*circuit) map[junctionBox]*circuit {
	merged := append(circuits[connection.from].junctionBoxes, circuits[connection.to].junctionBoxes...)
	filtered := utils.FilterDuplicates(merged)
	mergedCircuit := &circuit{
		junctionBoxes: filtered,
	}
	for _, jb := range mergedCircuit.junctionBoxes {
		circuits[jb] = mergedCircuit
	}
	return circuits
}

func sortConnectionsByDistanceAscending(sortedConnections []junctionConnection) []junctionConnection {
	sort.Slice(sortedConnections, func(i, j int) bool {
		return sortedConnections[i].distance < sortedConnections[j].distance
	})
	return sortedConnections
}

func calculateConnections(junctionBoxes []junctionBox) []junctionConnection {
	connections := make([]junctionConnection, 0, len(junctionBoxes))
	for i := range junctionBoxes {
		for j := i + 1; j < len(junctionBoxes); j++ {
			connections = append(connections, junctionConnection{
				from:     junctionBoxes[i],
				to:       junctionBoxes[j],
				distance: calculateEuclideanDistance(junctionBoxes[i], junctionBoxes[j]),
			})
		}
	}
	return connections
}

func puzzleInputLineToJunctionBox(line string) junctionBox {
	values := strings.Split(line, ",")
	// Since we know the shape of the input, we can ignore errors here as they won't occur
	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])
	z, _ := strconv.Atoi(values[2])
	return junctionBox{x: x, y: y, z: z}
}

func createJunctionBoxesFromInput(puzzleInput []string) []junctionBox {
	junctionBoxes := make([]junctionBox, 0, len(puzzleInput))
	for _, line := range puzzleInput {
		junctionBox := puzzleInputLineToJunctionBox(line)
		junctionBoxes = append(junctionBoxes, junctionBox)
	}
	return junctionBoxes
}

func calculateEuclideanDistance(jb1, jb2 junctionBox) float64 {
	dx := float64(jb2.x - jb1.x)
	dy := float64(jb2.y - jb1.y)
	dz := float64(jb2.z - jb1.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
