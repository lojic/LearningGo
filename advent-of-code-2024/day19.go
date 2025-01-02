package main

import ("strings")

func memoizeCheckDesign(pats []string) func (string) int {
  cache := make(map[string]int)
  var checkDesign func (string) int

  checkDesign = func (design string) int {
    if result, exists := cache[design]; exists {
      return result
    }

    if design == "" {
      return 1
    }

    total := 0

    for _, pat := range pats {
      after, found := strings.CutPrefix(design, pat)

      if found {
        total += checkDesign(after)
      }
    }

    cache[design] = total
    return total
  }

  return checkDesign
}

func main() {
  results := Parse[string](19, Words, "\n\n")
  pats    := results[0]
  designs := results[1]
  part1   := 0
  part2   := 0

  for _, design := range designs {
    cnt := memoizeCheckDesign(pats)(design)

    if cnt > 0 {
      part1 += 1
      part2 += cnt
    }
  }

  if part1 != 353 {
    panic("Part 1 failed")
  }

  if part2 != 880877787214477 {
    panic("Part 2 failed")
  }
}
