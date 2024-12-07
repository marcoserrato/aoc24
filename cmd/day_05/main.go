package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("cmd/day_05/input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)

	dependencies := make(map[int][]int)
	for {
		line, _ := reader.ReadBytes('\n')

		if string(line) == "\n" {
			break
		}
		numbers := strings.Split(string(line[:len(line)-1]), "|")
		left_number, _ := strconv.Atoi(numbers[0])
		right_number, _ := strconv.Atoi(numbers[1])

		if value, ok := dependencies[right_number]; ok {
			dependencies[right_number] = append(value, left_number)
		} else {
			temp := make([]int, 0)
			dependencies[right_number] = append(temp, left_number)
		}
	}

	pages := make([][]int, 0)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		current_page := make([]int, 0)
		numbers := strings.Split(string(line[:len(line)-1]), ",")
		for _, num := range numbers {
			inum, _ := strconv.Atoi(num)
			current_page = append(current_page, inum)
		}
		pages = append(pages, current_page)
	}

	sum := 0
	sum2 := 0
	for _, p := range pages {
		if valid_page(p, dependencies) {
			sum += p[len(p)/2]
		} else {
			sum2 += reordered_page(p, dependencies)
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", sum2)
}

func to_set(input []int) map[int]bool {
	hashset := make(map[int]bool)
	for _, v := range input {
		hashset[v] = true
	}
	return hashset
}

func union(self map[int]bool, other map[int]bool) []int {
	results := make([]int, 0)

	for k, _ := range self {
		if _, ok := other[k]; ok {
			results = append(results, k)
		}
	}

	return results
}

func reordered_page(page []int, dependencies map[int][]int) int {
	pageSet := to_set(page)
	numberOfDependent := make(map[int]int)
	for _, number := range page {
		if dependents, ok := dependencies[number]; ok {
			dependentsSet := to_set(dependents)
			dependenciesPresent := len(union(dependentsSet, pageSet))
			numberOfDependent[dependenciesPresent] = number
		} else {
			numberOfDependent[0] = number
		}
	}

	keys := make([]int, 0)
	for k := range numberOfDependent {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	result := make([]int, 0)
	for _, k := range keys {
		result = append(result, numberOfDependent[k])
	}

	return result[len(result)/2]
}

// Order maintains a map of int to array of int. The key is the page,
// and the array is the numbers that must appear before it.

// Pages is the page line. A number would be valid, if it's dependents have
// all occurred.

// eg
// Rule: 2|1 => 2 must appear before 1
// Order: 2, 1
// visited [2]
// order would have 1 => [2]
func valid_page(pages []int, order map[int][]int) bool {
	visited := make([]int, 0)

	for _, page := range pages {
		for _, visit := range visited {
			if dependencies, ok := order[visit]; ok {
				for _, dependency := range dependencies {
					if dependency == page {
						return false
					}
				}
			}
		}
		visited = append(visited, page)
	}
	return true
}
