package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Equation struct {
	result int
	values []int
}

func main() {
	file, err := os.Open("cmd/day_07/input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	equations := make([]Equation, 0)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		resultAndValues := strings.Split(string(line[:len(line)-1]), ":")
		result, _ := strconv.Atoi(resultAndValues[0])
		values := stringToNums(resultAndValues[1])
		equations = append(equations, Equation{result: result, values: values})
	}

	validCount := 0
	validCount2 := 0
	partOneOps := []string{"add", "mul"}
	partTwoOps := []string{"add", "mul", "cat"}

	for _, equation := range equations {
		if canSolve(0, slices.Clone(equation.values), equation.result, partOneOps) {
			validCount += equation.result
		}
		if canSolve(0, slices.Clone(equation.values), equation.result, partTwoOps) {
			validCount2 += equation.result
		}
	}

	fmt.Printf("Part 1: %d\n", validCount)
	fmt.Printf("Part 2: %d\n", validCount2)
}

func canSolve(runningSum int, nextValues []int, target int, ops []string) bool {
	if runningSum > target {
		return false
	}
	if len(nextValues) == 0 {
		return runningSum == target
	} else {
		value := nextValues[0]
		nextVals := nextValues[1:]
		solve := false
		for _, op := range ops {
			switch op {
			case "add":
				solve = solve || canSolve(runningSum+value, slices.Clone(nextVals), target, ops)
			case "mul":
				solve = solve || canSolve(runningSum*value, slices.Clone(nextVals), target, ops)
			case "cat":
				catted := fmt.Sprintf("%s%s", strconv.Itoa(runningSum), strconv.Itoa(value))
				cattedNum, _ := strconv.Atoi(catted)
				solve = solve || canSolve(cattedNum, slices.Clone(nextVals), target, ops)
			}
		}
		return solve
	}
}

func stringToNums(strNums string) []int {
	stringNums := strings.Split(strNums, " ")
	nums := make([]int, len(stringNums)-1)
	for i, num := range stringNums[1:] {
		v, _ := strconv.Atoi(num)
		nums[i] = v
	}
	return nums
}
