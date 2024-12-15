package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	file, err := os.Open("cmd/day_10/input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)

	starts := make([]Point, 0)
	hikeMap := make(map[int]map[int]int)
	row := 0

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		hikeMap[row] = make(map[int]int)

		for col, height := range strings.Split(string(line[:len(line)-1]), "") {
			heightInt, err := strconv.Atoi(height)

			if err != nil {
				panic(err)
			}
			hikeMap[row][col] = heightInt

			if heightInt == 0 {
				starts = append(starts, Point{row, col})
			}
		}
		row += 1
	}
	scores := 0
	scores2 := 0
	for _, start := range starts {
		scores += dfs(start, hikeMap, nil, make(map[Point]bool))
		scores2 += dfs2(start, hikeMap, nil)
	}

	fmt.Printf("Part 1: %d\n", scores)
	fmt.Printf("Part 2: %d\n", scores2)

}

func dfs2(point Point, grid map[int]map[int]int, prev *Point) int {
	// Only consider routes for which the gradient increases by one.
	if prev != nil && grid[prev.X][prev.Y]+1 != grid[point.X][point.Y] {
		return 0
	}
	if grid[point.X][point.Y] == 9 {
		return 1
	}
	if outOfBound(point, grid) {
		return 0
	}

	directions := []Point{Point{-1, 0}, Point{1, 0}, Point{0, -1}, Point{0, 1}}
	sum := 0
	for _, direction := range directions {
		newPoint := Point{point.X + direction.X, point.Y + direction.Y}
		sum += dfs2(newPoint, grid, &point)
	}

	return sum
}

func dfs(point Point, grid map[int]map[int]int, prev *Point, visited map[Point]bool) int {
	// Only consider routes for which the gradient increases by one.
	if prev != nil && grid[prev.X][prev.Y]+1 != grid[point.X][point.Y] {
		return 0
	}

	if grid[point.X][point.Y] == 9 {
		if _, ok := visited[point]; ok {
			return 0
		} else {
			visited[point] = true
			return 1
		}
	}

	if outOfBound(point, grid) {
		return 0
	}

	directions := []Point{Point{-1, 0}, Point{1, 0}, Point{0, -1}, Point{0, 1}}
	sum := 0
	for _, direction := range directions {
		newPoint := Point{point.X + direction.X, point.Y + direction.Y}
		sum += dfs(newPoint, grid, &point, visited)
	}

	return sum
}

func outOfBound(point Point, grid map[int]map[int]int) bool {
	return point.X < 0 || point.Y < 0 || point.X > len(grid) || point.Y > len(grid[0])
}
