package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func (p Point) eql(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) over(other Point) bool {
	return p.X > other.X || p.Y > other.Y
}

func (p Point) add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

type Game struct {
	buttonA Point
	buttonB Point
	prize   Point
}

func main() {
	file, err := os.Open("cmd/day_13/input.txt")
	if err != nil {
		panic(err)
	}

	buttonsRegex := regexp.MustCompile(`Button [A|B]: X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	games := make([]Game, 0)
	reader := bufio.NewReader(file)

	for {
		buttonA, err := reader.ReadBytes('\n')
		buttonB, _ := reader.ReadBytes('\n')
		prize, _ := reader.ReadBytes('\n')

		games = append(games, Game{
			capturePoint(buttonsRegex, string(buttonA[:len(buttonA)-1])),
			capturePoint(buttonsRegex, string(buttonB[:len(buttonB)-1])),
			capturePoint(prizeRegex, string(prize[:len(prize)-1])),
		})

		// Read empty line between games
		_, err = reader.ReadBytes('\n')
		if err != nil {
			break
		}
	}

	sum := 0
	sum2 := 0
	part2Augment := 10000000000000
	for _, game := range games {
		sum += solveGame(game)
		augmentedGame := Game{game.buttonA, game.buttonB, Point{game.prize.X + part2Augment, game.prize.Y + part2Augment}}
		sum2 += solveGame(augmentedGame)
	}

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", sum2)
}

func solveGame(game Game) int {
	// If the shift in x and y position is denoted as <button><direction> then the equations describing
	// the number of button presses for both A and B needed to arrive at our desired location would be:
	// Equation 1 for (X): ax*x + bx*x = px
	// Equation 2 for (Y): ay*y + by*y = py
	//
	// The intersection of these two lines (when happening at whole numbers, cause you can't 'partially' press
	// a button) would tell us the number of times to press each button to arrive at our answer.
	//
	// By solving an equation for one variable and plugging the result in the other (and vice-versa) we arrive
	// and the two equations below. If both numbers are whole, then we know there's a solution and we can calculate
	// the cost using 3a + b. I think the minimum number of tokens in the problem statement is a read herring. Since
	// these equations are equal (not less than or greater than) and are straight lines, there's no region or other
	// intersection points we can consider. There's either a solution or not!
	buttonBHits := float64(((game.buttonA.X * game.prize.Y) - (game.buttonA.Y * game.prize.X))) / float64(((game.buttonA.X * game.buttonB.Y) - (game.buttonA.Y * game.buttonB.X)))

	buttonAHits := float64(((game.buttonB.X * game.prize.Y) - (game.buttonB.Y * game.prize.X))) / float64(((game.buttonB.X * game.buttonA.Y) - (game.buttonB.Y * game.buttonA.X)))

	// Covert our float to int and back to see if we lose precision, which tells us if the result
	// is a whole number or not.
	if buttonBHits == float64(int64(buttonBHits)) && buttonAHits == float64(int64(buttonAHits)) {
		return 3*int(buttonAHits) + int(buttonBHits)
	} else {
		return 0
	}
}

type State struct {
	position Point
	acc      int
}

// My gut reaction to this problem was to use BFS and DFS instead of realizing
// this is a simple math problem :) These functions basically *never* return for
// the sample input. I'll keep them around as a reminder though.
func minTokens2(game Game) int {
	m := make([]State, 0)
	m = append(m, State{Point{0, 0}, 0})
	minimum := math.MaxInt

	for len(m) > 0 {
		f := m[0]
		m = m[1:]

		if f.position.eql(game.prize) {
			if f.acc < minimum {
				minimum = f.acc
			}
			continue
		}
		if f.position.over(game.prize) {
			continue
		}

		m = append(m, State{f.position.add(game.buttonB), f.acc + 1})
		m = append(m, State{f.position.add(game.buttonA), f.acc + 3})
	}
	return minimum
}

func minTokens(game Game, position Point, acc int, a int, b int) int {
	if position.eql(game.prize) {
		return acc
	}

	if position.over(game.prize) {
		return math.MaxInt
	}

	// Button A costs 3 | Bustton B costs 1
	left := minTokens(game, position.add(game.buttonA), acc+3, a+1, b)
	right := minTokens(game, position.add(game.buttonB), acc+1, a, b+1)

	return int(math.Min(float64(left), float64(right)))
}

func capturePoint(regex *regexp.Regexp, str string) Point {
	matches := regex.FindStringSubmatch(str)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])

	return Point{x, y}
}
