package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
  "math"
)

func main() {
  file, err := os.Open("cmd/day_02/input.txt")
  if err != nil {
    panic(err)
  }

  reader := bufio.NewReader(file)

  safe := 0
  safe2 := 0

  for {
    line, err := reader.ReadBytes('\n')
    if err != nil {
      break
    }
    str_line := string(line[:len(line)-1])
    values := strings.Split(str_line, " ")

    converted := make([]int, 0)

    for _, value := range(values) {
      nint, err := strconv.Atoi(value)
      if err != nil {
        panic(err)
      }
      converted = append(converted, nint)
    }

    if safe_level(converted) {
      safe += 1
    }
    if safe_level_2(converted) {
      safe2 += 1
      fmt.Println(converted)
    } else {
    }
  }

  fmt.Printf("Part 1 %d\n", safe)
  fmt.Printf("Part 2 %d\n", safe2)
}

func safe_without(lvls []int, ind int) bool {
  f := make([]int, len(lvls))
  b := make([]int, len(lvls))
  t := make([]int, len(lvls))

  var remove_before bool
  var remove_me bool
  var remove_after bool

  // Appending to slices mutates the underlying slice, so we need to
  // copy these.
  copy(f, lvls)
  copy(b, lvls)
  copy(t, lvls)

  if ind >= 2 {
    remove_before = safe_level(append(t[:ind-2], t[ind - 1:]...))
  }
  remove_me = safe_level(append(f[:ind-1], f[ind:]...))
  remove_after = safe_level(append(b[:ind], b[ind+1:]...))

  return remove_before || remove_me || remove_after
}

func safe_level_2(lvls []int) bool {
  previous := lvls[1] > lvls[0]
  index := 1
  for index < len(lvls) {
    next := lvls[index] > lvls[index-1]
    if previous != next { 
      return safe_without(lvls, index)
    } else if k := int(math.Abs(float64(lvls[index]) - float64(lvls[index-1]))); k < 1 || k > 3 {
      return safe_without(lvls, index)
    } 
    next = previous
    index+=1
  }

  return true
}

func safe_level(lvls []int) bool {
  previous := lvls[0] > lvls[1]
  index := 0

  for index < (len(lvls) - 1) {
    next := lvls[index] > lvls[index+1]
    if previous != next {
      return false
    } else if k := int(math.Abs(float64(lvls[index]) - float64(lvls[index+1]))); k < 1 || k > 3 {
      return false
    } 
    next = previous
    index+=1
  }
  return true
}
