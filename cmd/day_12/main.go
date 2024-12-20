package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Point struct {
	X int
	Y int
}

type DirectionalNum struct {
	direction string
	num       int
}

func main() {
	file, err := os.Open("cmd/day_12/input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	grid := make(map[int]map[int]string)
	row := 0

	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			break
		}
		grid[row] = make(map[int]string)

		for col, letter := range strings.Split(string(line[:len(line)-1]), "") {
			grid[row][col] = letter
		}
		row += 1
	}

	visited := make(map[Point]bool)
	visited2 := make(map[Point]bool)
	sum := 0
	sum2 := 0

	for row = 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			p, a := explore(row, col, grid, visited)
			p2, a2 := explore2(row, col, grid, visited2)
			sum += p * a
			sum2 += p2 * a2
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", sum2)
}

func explore(row int, col int, grid map[int]map[int]string, visited map[Point]bool) (int, int) {
	if _, ok := visited[Point{row, col}]; !ok {
		return _explore(grid[row][col], row, col, grid, visited, make(map[Point]bool))
	} else {
		return 0, 0
	}
}

func explore2(row int, col int, grid map[int]map[int]string, visited map[Point]bool) (int, int) {
	if _, ok := visited[Point{row, col}]; !ok {
		points := make(map[Point]bool)
		actual := make(map[Point]bool)
		findPoints(grid[row][col], row, col, grid, visited, points, actual)

		area, perim := calculateSides(grid[row][col], actual, grid), len(points)
		return area, perim
	} else {
		return 0, 0
	}
}

func calculateSides(current string, points map[Point]bool, grid map[int]map[int]string) int {
	sidesMap := make(map[string]map[int][]DirectionalNum)
	sidesMap["row"] = make(map[int][]DirectionalNum)
	sidesMap["col"] = make(map[int][]DirectionalNum)

	directions := []Point{Point{1, 0}, Point{0, 1}, Point{-1, 0}, Point{0, -1}} // U, R, D, L
	for point := range points {
		for i, direction := range directions {
			nextX := point.X + direction.X
			nextY := point.Y + direction.Y
			if (nextX < 0 || nextY < 0 || nextX > len(grid) || nextY > len(grid[0])) || grid[nextX][nextY] != current {
				switch i {
				case 0:
					if v, ok := sidesMap["row"][point.X+1]; ok {
						sidesMap["row"][point.X+1] = append(v, DirectionalNum{"up", point.Y})
					} else {
						t := make([]DirectionalNum, 0)
						sidesMap["row"][point.X+1] = append(t, DirectionalNum{"up", point.Y})
					}
				case 1:
					if v, ok := sidesMap["col"][point.Y+1]; ok {
						sidesMap["col"][point.Y+1] = append(v, DirectionalNum{"right", point.X})
					} else {
						t := make([]DirectionalNum, 0)
						sidesMap["col"][point.Y+1] = append(t, DirectionalNum{"right", point.X})
					}
				case 2:
					if v, ok := sidesMap["row"][point.X-1]; ok {
						sidesMap["row"][point.X-1] = append(v, DirectionalNum{"down", point.Y})
					} else {
						t := make([]DirectionalNum, 0)
						sidesMap["row"][point.X-1] = append(t, DirectionalNum{"down", point.Y})
					}
				case 3:
					if v, ok := sidesMap["col"][point.Y-1]; ok {
						sidesMap["col"][point.Y-1] = append(v, DirectionalNum{"left", point.X})
					} else {
						t := make([]DirectionalNum, 0)
						sidesMap["col"][point.Y-1] = append(t, DirectionalNum{"left", point.X})
					}
				}
			}
		}
	}

	sum := 0
	for _, innerMap := range sidesMap {
		for _, ints := range innerMap {
			imap := groupList(ints)
			for _, is := range imap {
				sort.Ints(is)
				sum += 1
				for i := 1; i < len(is); i++ {
					if is[i]-is[i-1] != 1 {
						sum += 1
					}
				}
			}
		}
	}

	return sum
}

func groupList(list []DirectionalNum) map[string][]int {
	m := make(map[string][]int)

	for _, el := range list {
		if n, k := m[el.direction]; k {
			m[el.direction] = append(n, el.num)
		} else {
			t := make([]int, 0)
			m[el.direction] = append(t, el.num)
		}
	}

	return m

}

func _explore(c string, row int, col int, grid map[int]map[int]string, visited map[Point]bool, n map[Point]bool) (int, int) {
	if _, ok := n[Point{row, col}]; ok {
		return 0, 0
	}

	if row < 0 || col < 0 || row > len(grid) || col > len(grid[0]) {
		return 1, 0
	}

	if grid[row][col] != c {
		return 1, 0
	}

	perim := 0
	area := 1

	visited[Point{row, col}] = true
	n[Point{row, col}] = true

	directions := []Point{Point{1, 0}, Point{0, 1}, Point{-1, 0}, Point{0, -1}}
	for _, direction := range directions {
		nextRow := row + direction.X
		nextCol := col + direction.Y

		nPerim, nArea := _explore(c, nextRow, nextCol, grid, visited, n)
		perim += nPerim
		area += nArea
	}

	return perim, area
}

func findPoints(c string, row int, col int, grid map[int]map[int]string, visited map[Point]bool, points map[Point]bool, actual map[Point]bool) int {
	if _, ok := points[Point{row, col}]; ok {
		return 0
	}

	if row < 0 || col < 0 || row > len(grid) || col > len(grid[0]) {
		return 1
	}

	if grid[row][col] != c {
		return 1
	}

	visited[Point{row, col}] = true
	points[Point{row, col}] = true
	allFour := 0

	directions := []Point{Point{1, 0}, Point{0, 1}, Point{-1, 0}, Point{0, -1}}
	for _, direction := range directions {
		nextRow := row + direction.X
		nextCol := col + direction.Y

		allFour += findPoints(c, nextRow, nextCol, grid, visited, points, actual)
	}

	if allFour > 0 {
		actual[Point{row, col}] = true
	}

	return 0
}
