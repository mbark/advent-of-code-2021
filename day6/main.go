package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
	"strings"
)

var in = `
4,1,1,1,5,1,3,1,5,3,4,3,3,1,3,3,1,5,3,2,4,4,3,4,1,4,2,2,1,3,5,1,1,3,2,5,1,1,4,2,5,4,3,2,5,3,3,4,5,4,3,5,4,2,5,5,2,2,2,3,5,5,4,2,1,1,5,1,4,3,2,2,1,2,1,5,3,3,3,5,1,5,4,2,2,2,1,4,2,5,2,3,3,2,3,4,4,1,4,4,3,1,1,1,1,1,4,4,5,4,2,5,1,5,4,4,5,2,3,5,4,1,4,5,2,1,1,2,5,4,5,5,1,1,1,1,1,4,5,3,1,3,4,3,3,1,5,4,2,1,4,4,4,1,1,3,1,3,5,3,1,4,5,3,5,1,1,2,2,4,4,1,4,1,3,1,1,3,1,3,3,5,4,2,1,1,2,1,2,3,3,5,4,1,1,2,1,2,5,3,1,5,4,3,1,5,2,3,4,4,3,1,1,1,2,1,1,2,1,5,4,2,2,1,4,3,1,1,1,1,3,1,5,2,4,1,3,2,3,4,3,4,2,1,2,1,2,4,2,1,5,2,2,5,5,1,1,2,3,1,1,1,3,5,1,3,5,1,3,3,2,4,5,5,3,1,4,1,5,2,4,5,5,5,2,4,2,2,5,2,4,1,3,2,1,1,4,4,1,5
`

var testData = `3,4,3,1,2`

func main() {
	numbers := util.NumberList(strings.Trim(in, "\n"), ",")
	fmt.Printf("first: %d\n", simulate(numbers, 80))
	fmt.Printf("second: %d\n", simulate(numbers, 256))
}

func simulate(numbers []int, generations int) int {
	m := make([]int, 9)
	for _, n := range numbers {
		m[n] += 1
	}

	for g := 1; g <= generations; g += 1 {
		next := make([]int, 9)
		for n, c := range m {
			if n-1 < 0 {
				next[6] += c
				next[8] = c
			} else {
				next[n-1] += c
			}
		}

		m = next
	}

	var count int
	for _, c := range m {
		count += c
	}

	return count
}
