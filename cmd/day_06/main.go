package main

import (
	"bufio"
	"fmt"
	"os"
)

type Guard struct {
	position  Point
	direction string
}

type Point struct {
	row int
	col int
}

var directionSequence = map[string]string{"up": "right", "left": "up", "right": "down", "down": "left"}

func main() {
	file, err := os.Open("cmd/day_06/input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)

	grid := make(map[int]map[int]string)
	var startPoint Point

	row := 0
	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			break
		}

		col := 0
		grid[row] = make(map[int]string)

		for _, character := range line[:len(line)-1] {
			if string(character) == "^" {
				startPoint = Point{row, col}
				grid[row][col] = "."
			} else {
				grid[row][col] = string(character)
			}
			col += 1
		}
		row += 1
	}

	fmt.Printf("Part 1: %d\n", partOne(startPoint, grid))
	fmt.Printf("Part 2 (Wrong): %d\n", partTwo(startPoint, grid))
	fmt.Printf("Part 2: %d\n", partTwoMod(startPoint, grid))
}

func partOne(start Point, grid map[int]map[int]string) int {
	visited := make(map[Point]bool)
	guard := Guard{position: start, direction: "up"}
	for {
		if ok := guard.step(grid); ok {
			visited[Point{guard.position.row, guard.position.col}] = true
		} else {
			break
		}
	}
	return len(visited)
}

func partTwoMod(start Point, grid map[int]map[int]string) int {
	cycles := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "#" || (i == start.row && j == start.col) {
				continue
			}
			grid[i][j] = "#"
			if causesCycle(grid, Guard{position: start, direction: "up"}) {
				cycles += 1
			}

			grid[i][j] = "."
		}
	}
	return cycles
}

// Tried being clever by only considering points that occur in front of the guard
// but I think I must be making a wrong assumption since this ends up returning a
// higher number (so somehow counting duplicates maybe?)
func partTwo(start Point, grid map[int]map[int]string) int {
	cycles := 0
	visited := make(map[Point]bool)
	guard := Guard{position: start, direction: "up"}
	for {
		prow, pcol := guard.position.row, guard.position.col
		pdirection := guard.direction

		if ok := guard.step(grid); !ok {
			break
		}

		nrow, ncol := guard.position.row, guard.position.col

		if ok := visited[Point{nrow, ncol}]; ok {
			continue
		}

		grid[nrow][ncol] = "#"

		if causesCycle(grid, Guard{position: Point{prow, pcol}, direction: pdirection}) {
			cycles += 1
			visited[Point{nrow, ncol}] = true
		}
		grid[nrow][ncol] = "."
	}

	return cycles
}

func (g *Guard) step(grid map[int]map[int]string) bool {
	nextRow, nextCol := g.position.row, g.position.col

	// Calculate what the next position the Guard will be at.
	switch g.direction {
	case "up":
		nextRow -= 1
	case "left":
		nextCol -= 1
	case "right":
		nextCol += 1
	case "down":
		nextRow += 1
	}

	// If it falls outside the grid, we can no longer step
	if exited(Point{nextRow, nextCol}, len(grid), len(grid[0])) {
		return false
	}

	// If we can move, move, if not, turn and try and move again.
	if canMoveTo(grid, Point{nextRow, nextCol}) {
		g.position.row = nextRow
		g.position.col = nextCol
		return true
	} else {
		g.direction = directionSequence[g.direction]
		return g.step(grid)
	}
}

// A cycle would happen if direction and position occur twice.
func causesCycle(grid map[int]map[int]string, guard Guard) bool {
	slow := Guard{guard.position, guard.direction}
	fast := Guard{guard.position, guard.direction}

	if ok := fast.step(grid); !ok {
		return false
	}

	if ok := fast.step(grid); !ok {
		return false
	}

	for {
		if slow == fast {
			return true
		}

		slow.step(grid)
		if ok := fast.step(grid); !ok {
			return false
		}
		if ok := fast.step(grid); !ok {
			return false
		}
	}
}

func canMoveTo(grid map[int]map[int]string, point Point) bool {
	return grid[point.row][point.col] == "." || grid[point.row][point.col] == "X"
}

func exited(position Point, r int, c int) bool {
	return position.row >= r || position.row < 0 || position.col >= c || position.col < 0
}
