package main

import ("slices")

func part1(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	total := 0

	for i, n1 := range left {
		total += AbsInt(n1 - right[i])
	}

	return total
}

func part2(left, right []int) int {
	counter := make(map[int]int, len(right))

	for _, n := range right {
		counter[n] += 1
	}

	total := 0

	for _, n := range left {
		total += n * counter[n]
	}

	return total
}

func main() {
	var left, right []int

	for _, pair := range Parse[int](1, Ints, "\n") {
		left  = append(left, pair[0])
		right = append(right, pair[1])
	}

	if part1(left, right) != 1189304 {
		panic("Part 1 failed")
	}

	if part2(left, right) != 24349736 {
		panic("Part 2 failed")
	}
}
