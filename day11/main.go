package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/maps"
)

var testData = `
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

var testData1 = `
11111
19991
19191
19991
11111
`

var in = `
4836484555
4663841772
3512484556
1481547572
7741183422
8683222882
4215244233
1544712171
5725855786
1717382281
`

func main() {
	m := maps.NewIntMap(testData)
	first, second := both(m)
	fmt.Printf("first: %d\n", first)
	fmt.Printf("second: %d\n", second)
}

func both(m maps.IntMap) (int, int) {
	var total int
	step := 1
	for ; ; step++ {
		for _, c := range m.Coordinates() {
			m.Inc(c)
		}

		flashed := make(map[maps.Coordinate]struct{})
		for {
			var count int
			for _, c := range m.Coordinates() {
				v := m.At(c)
				if v <= 9 {
					continue
				}
				if _, ok := flashed[c]; ok {
					continue
				}

				// not flashed and >= 9
				count += 1
				flashed[c] = struct{}{}
				for _, ac := range m.Surrounding(c) {
					m.Set(ac, m.At(ac)+1)
				}
			}
			if count == 0 {
				break
			}

			if step <= 100 {
				total += count
			}
		}

		if len(flashed) == m.Length() {
			break
		}
		for c := range flashed {
			m.Set(c, 0)
		}
	}

	return total, step
}
