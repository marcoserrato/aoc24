package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) inGrid(height int, width int) bool {
	return p.X >= 0 && p.X < height && p.Y >= 0 && p.Y < width
}

func main() {
	file, err := os.Open("cmd/day_08/input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)

	pointsByType := make(map[string][]Point)
	grid := make(map[int]map[int]string)
	row := 0

	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			break
		}

		grid[row] = make(map[int]string)

		for col, v := range strings.Split(string(line[:len(line)-1]), "") {
			if v != "." {
				if arr, ok := pointsByType[v]; ok {
					pointsByType[v] = append(arr, Point{row, col})
				} else {
					temp := make([]Point, 0)
					pointsByType[v] = append(temp, Point{row, col})
				}
			}
			grid[row][col] = v
		}
		row += 1
	}

	height := len(grid)
	width := len(grid[0])
	visited := make(map[Point]bool)
	visited2 := make(map[Point]bool)

	for _, points := range pointsByType {
		for i, point := range points {
			cop := slices.Clone(points)
			cop2 := slices.Clone(points)
			addPoints(point, append(cop[:i], cop[i+1:]...), height, width, visited)
			addPoints2(point, append(cop2[:i], cop2[i+1:]...), height, width, visited2)
		}
	}

	fmt.Printf("Part 1: %d\n", len(visited))
	fmt.Printf("Part 2: %d\n", len(visited2))
}

func addPoints2(point Point, others []Point, height int, width int, visited map[Point]bool) {
	x1, y1 := point.X, point.Y
	visited[point] = true
	for _, otherPoint := range others {
		x2, y2 := otherPoint.X, otherPoint.Y

		slopeY := y2 - y1
		slopeX := x2 - x1

		downX, downY := x1-slopeX, y1-slopeY
		upX, upY := x2+slopeX, y2+slopeY

		for validPoint(Point{downX, downY}, height, width, visited) {
			downX -= slopeX
			downY -= slopeY
		}

		for validPoint(Point{upX, upY}, height, width, visited) {
			upX += slopeX
			upY += slopeY
		}
	}
}

func addPoints(point Point, others []Point, height int, width int, visited map[Point]bool) {
	x1, y1 := point.X, point.Y
	for _, otherPoint := range others {
		x2, y2 := otherPoint.X, otherPoint.Y

		slopeY := y2 - y1
		slopeX := x2 - x1

		validPoint(Point{x1 - slopeX, y1 - slopeY}, height, width, visited)
		validPoint(Point{x2 + slopeX, y2 + slopeY}, height, width, visited)
	}
}

func validPoint(p Point, height int, width int, visited map[Point]bool) bool {
	if p.inGrid(height, width) {
		visited[p] = true
		return true
	} else {
		return false
	}
}
