package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("cmd/day_11/input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	line, err := reader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	numbers := make([]int, 0)
	for _, number := range strings.Split(string(line[:len(line)-1]), " ") {
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, numberInt)

	}

	fmt.Printf("Part 1: %d\n", blink(numbers, 25))
	fmt.Printf("Part 2: %d\n", blinkDP(numbers, 75))
}

func blink(input []int, times int) int {
	for blink := 0; blink < times; blink++ {
		temp := make([]int, 0)
		for _, number := range input {
			if number == 0 {
				temp = append(temp, 1)
			} else if numStr := strconv.Itoa(number); len(numStr)%2 == 0 {
				left := numStr[:len(numStr)/2]
				right := numStr[(len(numStr) / 2):]

				leftInt, _ := strconv.Atoi(left)
				rightInt, _ := strconv.Atoi(right)

				temp = append(temp, leftInt)
				temp = append(temp, rightInt)
			} else {
				temp = append(temp, number*2024)
			}
		}
		input = temp
	}
	return len(input)
}

func blinkDP(input []int, times int) int {
	dp := make(map[int]map[int]int)
	sum := 0
	for _, number := range input {
		sum += _blinkDP2(number, times, dp)
	}
	return sum
}

func _blinkDP2(num int, blink int, dp map[int]map[int]int) int {
	if blink == 0 {
		return 1
	}

	if _, ok := dp[num]; !ok {
		dp[num] = make(map[int]int)
	}

	if stones, ok := dp[num][blink]; ok {
		// We've already calculated this number
		return stones
	}

	if num == 0 {
		dp[num][blink] = _blinkDP2(1, blink-1, dp)
	} else if numStr := strconv.Itoa(num); len(numStr)%2 == 0 {
		left := numStr[:len(numStr)/2]
		right := numStr[(len(numStr) / 2):]

		leftInt, _ := strconv.Atoi(left)
		rightInt, _ := strconv.Atoi(right)

		dp[num][blink] = _blinkDP2(leftInt, blink-1, dp) + _blinkDP2(rightInt, blink-1, dp)
	} else {
		dp[num][blink] = _blinkDP2(num*2024, blink-1, dp)
	}

	return dp[num][blink]
}
