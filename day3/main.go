package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed "input"
var input string

func main() {
	fmt.Println(compute(input))
}

func compute(in string) int {
	in = strings.TrimSpace(in)
	m := strings.Split(in, "\n")
	for i := range m {
		m[i] = strings.TrimSpace(m[i])
	}

	sum := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] != '.' && !unicode.IsDigit(rune(m[i][j])) {
				sum += sumAdjascentTo(m, i, j)
			}
		}
	}
	return sum
}

// sumAdjascentTo sum the numbers adjascent to current position.
// current position must be a symbol and != "."
func sumAdjascentTo(m []string, i, j int) int {

	// delta i's and delta j's
	dis := []int{+0, +0, -1, +1, +1, -1, +1, -1}
	djs := []int{-1, +1, +0, +0, +1, +1, -1, -1}

	bakI, bakJ := i, j
	sum := 0

	for k := range dis {
		i = bakI + dis[k]
		j = bakJ + djs[k]

		if i < 0 || i >= len(m) {
			continue
		}
		if j < 0 || j >= len(m[i]) {
			continue
		}

		// if the caracter is a digit, it finds all digits and parses and sums it
		if unicode.IsDigit(rune(m[i][j])) {
			sum += parseNumberInPosition(m, i, j)
		}
	}

	return sum
}

// parseNumberInPosition finds and parses the number and replace it with "..." to avoid the problem of double counting
func parseNumberInPosition(m []string, i, j int) int {
	bakJ := j
	ini := j
	end := j

	l := m[i]
	for ; j >= 0 && unicode.IsDigit(rune(l[j])); j-- {
		ini = j
	}

	j = bakJ

	for ; j < len(l) && unicode.IsDigit(rune(l[j])); j++ {
		end = j
	}

	sNum := l[ini : end+1]
	num, err := strconv.Atoi(sNum)
	if err != nil {
		panic(err)
	}

	m[i] = l[:ini] + strings.Repeat(".", len(sNum)) + l[end+1:]
	return num
}
