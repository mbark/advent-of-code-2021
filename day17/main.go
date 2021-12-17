package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/maps"
	"github.com/mbark/advent-of-code-2021/util"
	"strings"
)

const (
	testData = `target area: x=20..30, y=-10..-5`
	in       = `target area: x=201..230, y=-99..-65`
)

func main() {
	s := strings.TrimPrefix(in, "target area: ")
	split := strings.Split(s, ", ")
	xs := strings.Split(strings.TrimPrefix(split[0], "x="), "..")
	ys := strings.Split(strings.TrimPrefix(split[1], "y="), "..")

	xstart, xend := util.Str2Int(xs[0]), util.Str2Int(xs[1])
	ystart, yend := util.Str2Int(ys[0]), util.Str2Int(ys[1])

	upperLeft := maps.Coordinate{X: xstart, Y: ystart}
	lowerLeft := maps.Coordinate{X: xend, Y: yend}
	first, second := first(upperLeft, lowerLeft)
	fmt.Printf("first %d\n", first)
	fmt.Printf("second %d\n", second)
}

func first(upperLeft, lowerRight maps.Coordinate) (int, int) {
	var successCount int
	var successfulMaxY int

	for y := -100; y < 100; y++ {
		for x := 1; x <= lowerRight.X; x++ {
			velx := x
			vely := y
			pos := maps.Coordinate{}

			var success, fail bool
			var maxY int
			for step := 0; !(success || fail); step++ {
				pos.X += velx
				pos.Y += vely

				if velx > 0 {
					velx -= 1
				} else if velx < 0 {
					velx += 1
				}
				vely -= 1

				if pos.Y > maxY {
					maxY = pos.Y
				}

				xWithin := upperLeft.X <= pos.X && pos.X <= lowerRight.X
				yWithin := upperLeft.Y <= pos.Y && pos.Y <= lowerRight.Y

				if xWithin && yWithin {
					success = true
					successCount += 1
					if maxY > successfulMaxY {
						successfulMaxY = maxY
					}
				} else if pos.X > lowerRight.X || pos.Y < upperLeft.Y {
					fail = true
				}
			}
		}
	}

	return successfulMaxY, successCount
}
