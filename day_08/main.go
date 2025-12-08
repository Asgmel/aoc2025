package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/asgmel/aockit/input"
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

func main() {
	taskOne("./example.txt")
}

func taskOne(inputPath string) {
	puzzleInput := input.ReadInputLines(inputPath)
	junctionBoxes := createJunctionBoxesFromInput(puzzleInput)
	connections := calculateConnections(junctionBoxes)
	for _, conn := range connections {
		fmt.Printf("From (%d,%d,%d) to (%d,%d,%d): Distance %.2f\n",
			conn.from.x, conn.from.y, conn.from.z,
			conn.to.x, conn.to.y, conn.to.z,
			conn.distance)
	}
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
