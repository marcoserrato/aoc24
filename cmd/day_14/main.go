package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// X is the distance from the left wall (columns)
// Y is the distance from the top wall (rows)

type Point struct {
	X int
	Y int
}

type Velocity Point
type Boundary Point

type Robot struct {
	position Point
	velocity Velocity
}

func (r *Robot) simulate(boundary Boundary) {
	nextX := (r.position.X + r.velocity.X)
	nextY := (r.position.Y + r.velocity.Y)

	if nextX < 0 {
		nextX += boundary.X
	} else {
		nextX = nextX % boundary.X
	}

	if nextY < 0 {
		nextY += boundary.Y
	} else {
		nextY = nextY % boundary.Y
	}

	r.position = Point{nextX, nextY}
}

func printRobots(robots []*Robot, boundary Boundary) {
	grid := make(map[int]map[int]string)

	for i := 0; i < boundary.Y; i++ {
		grid[i] = make(map[int]string)
	}

	for _, robot := range robots {
		grid[robot.position.Y][robot.position.X] = "X"
	}

	for i := 0; i < boundary.Y; i++ {
		for j := 0; j < boundary.X; j++ {
			if v, ok := grid[i][j]; ok {
				fmt.Printf(v)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {
	file, err := os.Open("cmd/day_14/input.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	robots := make([]*Robot, 0)
	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			break
		}

		point, vel := capturePointAndVel(re, string(line[:len(line)-1]))

		robots = append(robots, &Robot{point, vel})
	}

	boundary := Boundary{101, 103}

	for i := 0; i < 100; i++ {
		for _, robot := range robots {
			robot.simulate(boundary)
		}
	}

	fmt.Printf("Part 1: %d\n", calculateQuadrantScore(robots, boundary))

	count := 100
	for !uniq(robots) {
		for _, robot := range robots {
			robot.simulate(boundary)
		}
		count += 1
	}

	printRobots(robots, boundary)

	fmt.Printf("Part 2: %d\n", count)
}

func uniq(robots []*Robot) bool {
	m := make(map[Point]bool)

	for _, robot := range robots {
		if _, ok := m[robot.position]; ok {
			return false
		} else {
			m[robot.position] = true
		}
	}
	return true
}

func calculateQuadrantScore(robots []*Robot, boundary Boundary) int {
	quadrants := make(map[int]int)
	middleX := (boundary.X / 2)
	middleY := (boundary.Y / 2)

	for _, robot := range robots {
		pos := robot.position
		if pos.X == middleX || pos.Y == middleY {
			continue
		}
		quadrantX := pos.X / (middleX + 1)
		quadrantY := pos.Y / (middleY + 1)

		quadrants[quadrantX+quadrantY*2] += 1
	}

	sum := 1
	for _, numRobots := range quadrants {
		sum *= numRobots
	}

	return sum
}

func capturePointAndVel(re *regexp.Regexp, str string) (Point, Velocity) {
	matches := re.FindStringSubmatch(str)

	px, _ := strconv.Atoi(matches[1])
	py, _ := strconv.Atoi(matches[2])
	vx, _ := strconv.Atoi(matches[3])
	vy, _ := strconv.Atoi(matches[4])

	return Point{px, py}, Velocity{vx, vy}
}
