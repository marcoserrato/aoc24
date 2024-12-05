package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
  "sort"
  "math"
  "fmt"
)


func main() {
  file, err := os.Open("cmd/day_01/input.txt")

  if err != nil {
    panic(err)
  }

  reader := bufio.NewReader(file)

  left := make([]int, 0)
  right := make([]int, 0)

  for {
    line, err := reader.ReadBytes('\n')

    if err != nil {
      break
    }

    str_line := string(line[:len(line)-1])
    numbers := strings.Split(str_line, "   ")

    left_number, err := strconv.Atoi(numbers[0])
    if err != nil {
      panic(err)
    }
    right_number, err := strconv.Atoi(numbers[1])
    if err != nil {
      panic(err)
    }
    left = append(left, left_number)
    right = append(right, right_number)
  }

  sort.Ints(left)
  sort.Ints(right)


  fmt.Printf("First Part: %d\n", first_part(left, right))
  fmt.Printf("Second Part: %d\n", second_part(left, right))
}

func first_part(left []int, right []int) int {
  var sum float64 = 0

  for i, num := range(left) {
    sum += math.Abs(float64(num - right[i]))
  }
  return int(sum)
}

func second_part(left []int, right []int) int {
  rightm := make(map[int]int)

  for _, num := range(right) {
    if _, ok := rightm[num]; ok {
      rightm[num] += 1
    } else {
      rightm[num] = 1
    }
  }
  
  sum := 0
  for _, num := range(left) {
    if v, ok := rightm[num]; ok {
      sum += num * v
    }
  }

  return sum
}
