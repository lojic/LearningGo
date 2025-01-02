package main

import (
  "aoc/advent"
)

func plus(x, y int) int { return x + y }
func mult(x, y int) int { return x * y }

func conc(x, y int) int {
  power  := pow10(ilog10(y))
  return x * power + y
}

func ilog10(n int) int {
  var count int
  for n >= 10 {
    n /= 10
    count++
  }
  return count
}

func pow10(n int) int {
  result := 10

  for n > 0 {
    result *= 10
    n -= 1
  }

  return result
}

func isValid(ops []func (int, int) int, answer, result int, operands []int) bool {
  if result > answer {
    return false
  }
  if len(operands) == 0 {
    return result == answer
  }

  for _, op := range ops {
    if isValid(ops, answer, op(result, operands[0]), operands[1:]) {
      return true
    }
  }

  return false
}

func solve(operators ...func (int, int) int) int {
  input := advent.Parse[int](7, advent.Ints, "\n")
  total := 0

  for _, lst := range input {
    if isValid(operators, lst[0], lst[1], lst[2:]) {
      total += lst[0]
    }
  }

  return total
}

func main() {
  if solve(mult, plus) != 2664460013123 {
    panic("Part 1 failed")
  }

  if solve(mult, plus, conc) != 426214131924213 {
    panic("Part 2 failed")
  }
}
