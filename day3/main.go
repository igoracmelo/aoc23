package main

import (
	"strconv"
	"strings"
	"unicode"
)

func main() {

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

func sumAdjascentTo(m []string, i, j int) int {
	sum := 0

	dj := -1
	for ; j+dj >= 0 && unicode.IsDigit(rune(m[i][j+dj])); dj-- {
	}

	if dj != -1 {
		sNum := m[i][j+dj+1 : j]
		num, err := strconv.Atoi(sNum)
		if err != nil {
			panic(err)
		}

		sum += num
	}

	dj = 1
	for ; j+dj < len(m[i]) && unicode.IsDigit(rune(m[i][j+dj])); dj++ {
	}

	if dj != 1 {
		sNum := m[i][j+1 : j+dj]
		num, err := strconv.Atoi(sNum)
		if err != nil {
			panic(err)
		}
		sum += num
	}

	return sum
}
