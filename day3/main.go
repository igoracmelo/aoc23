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

	bakI, bakJ := i, j
	_ = bakI

	djs := []int{-1, 1}
	dis := []int{0, 0}

	for k := range djs {
		di := dis[k]
		dj := djs[k]

		found := false
		for i, j = i+di, j+dj; i >= 0 && j >= 0 && i < len(m) && j < len(m[i]) && unicode.IsDigit(rune(m[i][j])); j += dj {
			found = true
		}
		if !found {
			j = bakJ
			continue
		}

		var sNum string
		if j > bakJ {
			sNum = m[i][bakJ+1 : j]
		} else {
			sNum = m[i][j+1 : bakJ]
		}

		j = bakJ

		num, err := strconv.Atoi(sNum)
		if err != nil {
			panic(err)
		}

		sum += num
	}
	return sum
}
