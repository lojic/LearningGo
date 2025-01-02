package main

import (
  "strconv"
)

func memoizeBlink() func (int, int) int {
  cache := make(map[[2]int]int)
  var blink func (int, int) int

  blink = func(stone, blinks int) int {
    if result, exists := cache[[2]int{stone,blinks}]; exists {
      return result
    }

    if blinks == 0 {
      cache[[2]int{stone,blinks}] = 1
      return 1
    }

    if stone == 0 {
      val := blink(1, blinks - 1)
      cache[[2]int{stone, blinks}] = val
      return val
    }

    s    := strconv.Itoa(stone)
    slen := len(s)
    l    := slen / 2
    r    := slen % 2

    if r == 0 {
      val1, err := strconv.Atoi(s[:l])
      if err != nil {
        panic(err)
      }
      val2, err := strconv.Atoi(s[l:])
      if err != nil {
        panic(err)
      }

      val := blink(val1, blinks - 1) + blink(val2, blinks - 1)
      cache[[2]int{stone, blinks}] = val
      return val
    } else {
      val := blink(stone * 2024, blinks - 1)
      cache[[2]int{stone, blinks}] = val
      return val
    }
  }

  return blink
}

func solve(blinks int) int {
  workers := 0
  ch      := make(chan int)

  for _, stone := range []int{8069, 87014, 98, 809367, 525, 0, 9494914, 5} {
    workers += 1
    go func(stone int) {
      ch <- memoizeBlink()(stone, blinks)
    }(stone)
  }

  total := 0

  for i := 0; i < workers; i++ {
    total += <-ch
  }

  return total
}

func main() {
  if solve(25) != 183484 {
    panic("Part 1 failed")
  }
  if solve(75) != 218817038947400 {
    panic("Part 2 failed")
  }
}
