package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
)

func main() {
  file, err := os.Open("cmd/day_04/input.txt")

  if err != nil {
    panic(err)
  }

  crossword := make(map[int]map[int]string)

  reader := bufio.NewReader(file)
  row := 0

  for {
    line, err := reader.ReadBytes('\n')
    if err != nil {
      break
    }
    crossword[row] = make(map[int]string)

    for col, char := range(line) {
      crossword[row][col] = string(char)
    }

    row+= 1
  }

  found := 0
  found2 := 0

  iter_row := 0
  for iter_row < len(crossword) {
    iter_col  := 0
    for iter_col < len(crossword[iter_row]) {
      found += search(iter_row, iter_col, crossword)

      if search_mas(iter_row, iter_col, crossword) {
        found2 += 1
      }
      iter_col +=1
    }
    iter_row += 1
  }
  fmt.Printf("Part 1: %d\n", found)
  fmt.Printf("Part 2: %d\n", found2)
}

func search_direction(row int, col int, drow int, dcol int, cross map[int]map[int]string) bool {
  word := []string{"X", "M", "A", "S"}
  index := 0
  for {
    if index > 3 {
      break
    }
    if x, ok := cross[row]; ok {
      if y, ok := x[col]; ok {
        if y == word[index] {
          row += drow
          col += dcol
          index +=1
        } else {
          return false
        }
      } else {
        return false
      }
    } else {
      return false
    }
  }

  return true
}

func search_mas(row int, col int, cross map[int]map[int]string) bool {
  if cross[row][col] == "A" {
    diag1 := [][]int{{-1, -1}, {0, 0}, {1, 1}}
    d1w, ok := build_word(row, col, diag1, cross)

    if !ok {
      return false
    } 

    diag2 := [][]int{{1, -1}, {0, 0}, {-1, 1}}
    d2w, ok := build_word(row, col, diag2, cross)

    if !ok {
      return false
   }

    return (*d1w == "SAM" || *d1w == "MAS") && (*d2w == "SAM" || *d2w == "MAS")
  } else {
    return false
  }
}

func build_word(row int, col int, coords [][]int, cross map[int]map[int]string) (*string, bool) {
  var word strings.Builder
  for _, coor := range(coords) {
    nrow := row + coor[0]
    ncol := col + coor[1]

    if v, ok := cross[nrow]; ok {
      if letter, ok := v[ncol]; ok {
        word.WriteString(letter)
      } else {
        return nil, false
      }
    } else {
      return nil, false
    }
  }

  str := word.String()

  return &str, true
}

func search(row int, col int, cross map[int]map[int]string) int {
  if cross[row][col] != "X" {
    return 0
  }

  found := 0
  directions := [][]int{{-1, 0 },{1, 0},{0, 1},{0, -1},{1, 1},{-1, -1},{1, -1},{-1, 1}}

  for _, dir := range(directions) {
    if search_direction(row, col, dir[0], dir[1], cross) {
      found += 1
    }
  }

  return found
}
