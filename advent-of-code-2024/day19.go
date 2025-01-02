package main

import (
  "io"
  "os"
  "strings"
)

type prefixSuffix struct {
  prefix string
  suffix string
}

func parse() ([]string, []string) {
  ifs, err := os.Open("day19.txt")

  if err != nil {
    panic(err)
  }

  b, err := io.ReadAll(ifs)

  if err != nil {
    panic(err)
  }

  sections := strings.Split(string(b), "\n\n")
  towels   := strings.Split(sections[0], ", ")
  designs  := strings.Split(sections[1], "\n")

  return towels, designs[:len(designs)-1] // Last one empty due to trailing \n in file
}

func prefixesSuffixes(design string, maxlen int) []prefixSuffix {
  limit  := min(maxlen, len(design)) + 1
  result := make([]prefixSuffix, limit-1)

  for i, prefixlen := 0, 1; prefixlen < limit; i, prefixlen = i+1, prefixlen+1 {
    result[i] = prefixSuffix{ prefix: design[0:prefixlen], suffix: design[prefixlen:] }
  }

  return result
}

func memoizeCheckDesign(towels map[string]bool, maxlen int) func (string) int {
  cache := make(map[string]int,64)
  var checkDesign func (string) int

  checkDesign = func (design string) int {
    if result, exists := cache[design]; exists {
      return result
    }

    arrangements := 0

    if _, ok := towels[design]; ok {
      arrangements = 1
    }

    for _, ps := range prefixesSuffixes(design, maxlen) {
      if _, ok := towels[ps.prefix]; ok {
        arrangements += checkDesign(ps.suffix)
      }
    }

    cache[design] = arrangements
    return arrangements
  }

  return checkDesign
}

func Solve() (int, int) {
  pats, designs := parse()
  towels := make(map[string]bool,500)
  maxlen := 0

  for _, towel := range pats {
    l := len(towel)

    if l > maxlen {
      maxlen = l
    }

    towels[towel] = true
  }

  part1   := 0
  part2   := 0

  for _, design := range designs {
    cnt := memoizeCheckDesign(towels, maxlen)(design)

    if cnt > 0 {
      part1 += 1
      part2 += cnt
    }
  }

  return part1, part2
}

func main() {
  part1, part2 := Solve()

  if part1 != 353 {
    panic("Part 1 failed")
  }

  if part2 != 880877787214477 {
    panic("Part 2 failed")
  }
}
