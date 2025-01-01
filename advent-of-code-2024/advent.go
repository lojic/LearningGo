package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Parse text into a single float or int or string
func Atom(text string) any {
	if f, err := strconv.ParseFloat(strings.Trim(text, " \n"), 64); err == nil {
		i := int(f)

		if f == float64(i) {
			return i
		}

		return f
	}

	return text
}

// Return a slice of all the atoms (numbers or symbol names) in text
func Atoms(text string) []any {
	return Map(regexp.MustCompile(`[a-zA-Z_0-9.+-]+`).FindAllString(text, -1), Atom)
}

// Return a slice of all the digits in text (as ints 0 - 9), ignoring non-digit characters
func Digits(text string) []int {
	return Map(regexp.MustCompile(`[0-9]`).FindAllString(text, -1), mustAtoi)
}

// Return a slice of integers
func Ints(text string) []int {
	return Map(regexp.MustCompile(`[0-9]+`).FindAllString(text, -1), mustAtoi)
}

func Parse[T any](day int, parser func(string) []T, sep string) [][]T {
	fname := fmt.Sprintf("day%02d.txt", day)

	content, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	return Map(strings.Split(strings.Trim(string(content), " \t\n"), sep), parser)
}

// --------------------------------------------------------------------------------------------
// Private Functions
// --------------------------------------------------------------------------------------------

// Convert a string to an int, and panic on failure
func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}
