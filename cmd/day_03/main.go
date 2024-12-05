package main

import (
	"bufio"
	"fmt"
	"os"
  "strconv"
)

func is_digit(num string) (bool) {
  _, err := strconv.Atoi(num)

  return err == nil
}

func parse_do(reader *bufio.Reader) bool {
  byte, err := reader.Peek(3)

  if err != nil {
    return false
  }
  if string(byte[:]) == "o()" {
    reader.Discard(3)
    return true
  }
  return false
}

func parse_dont(reader *bufio.Reader) bool {
  byte, err := reader.Peek(6)
  if err != nil {
    return false
  }
  if string(byte[:]) == "on't()" {
    reader.Discard(6)
    return true
  }
  return false
}

func main() {
  file, err := os.Open("cmd/day_03/input.txt")
  if err != nil {
    panic(err)
  }

  reader := bufio.NewReader(file)
  sum := 0
  sum2 := 0
  previous_char, _ := reader.ReadByte()
  in_sequence := false
  direction := false
  left := make([]byte, 0)
  right := make([]byte, 0)
  enabled := true

  for {
    char, err := reader.ReadByte()

    if err != nil {
      break
    }

    switch  {
    case string(char) == "d":
      if parse_do(reader) {
        enabled = true
      } else if parse_dont(reader) {
        enabled = false
      }
    case string(char) == "m":
      left = make([]byte, 0)
      right = make([]byte, 0)
      in_sequence = true
      direction = false
    case string(char) == "u" && string(previous_char) == "m" && in_sequence:
      break
    case string(char) == "l" && string(previous_char) == "u" && in_sequence:
      break
    case string(char) == "(" && string(previous_char) == "l" && in_sequence:
      break
    case (string(previous_char) == "(" || is_digit(string(previous_char)) || string(previous_char) == ",") && is_digit(string(char)) && in_sequence:
      if direction {
        right = append(right, char)
      } else {
        left = append(left, char)
      }
    case is_digit(string(previous_char)) && in_sequence && string(char) == ",":
      direction = true
    case is_digit(string(previous_char)) && in_sequence && string(char) == ")" && direction:
      nleft, _ := strconv.Atoi(string(left[:]))
      nright, _ := strconv.Atoi(string(right[:]))
      sum += nleft * nright
      if enabled {
        sum2 += nleft * nright
      }
    default:
      left = make([]byte, 0)
      right = make([]byte, 0)
      direction = false
      in_sequence = false
    }

    previous_char = char
  }

  fmt.Printf("Part 1: %d\n", sum)
  fmt.Printf("Part 2: %d\n", sum2)
}
