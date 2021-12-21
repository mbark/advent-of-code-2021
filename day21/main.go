package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
)

var (
	testData = `
Player 1 starting position: 4
Player 2 starting position: 8
`

	in = `
Player 1 starting position: 7
Player 2 starting position: 5
`
)

func main() {
	lines := util.ReadInput(in, "\n")
	l1 := lines[0]
	l2 := lines[1]

	start1 := util.Str2Int(string(l1[len(l1)-1]))
	start2 := util.Str2Int(string(l2[len(l2)-1]))

	fmt.Printf("first: %d\n", first(start1, start2))
	fmt.Printf("second: %d\n", second(start1, start2))
}

func first(start1, start2 int) int {
	// 0-9 instead of 1-10
	start1 -= 1
	start2 -= 1

	var score1, score2 int
	roll := 1

	var rolls int
	var loser int
	for ; ; {
		roll1 := roll + roll + roll + 3
		roll = roll + 3
		rolls += 3

		start1 = (start1 + roll1) % 10
		score1 += start1 + 1
		if score1 >= 1000 {
			loser = score2
			break
		}

		roll2 := roll + roll + roll + 3
		roll = roll + 3
		rolls += 3

		start2 = (start2 + roll2) % 10
		score2 += start2 + 1
		if score2 >= 1000 {
			loser = score1
			break
		}
	}

	return rolls * loser
}

var allRolls map[int]int

func init() {
	allRolls = make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				allRolls[i+j+k] += 1
			}
		}
	}
}

type State struct {
	Position1 int
	Position2 int
	Score1    int
	Score2    int
}

type Positions map[State]int

func second(start1, start2 int) int {
	endAt := 21
	scores := make([]Positions, endAt+10)
	for i := 0; i < endAt+10; i++ {
		scores[i] = Positions{}
	}
	scores[0][State{Position1: start1 - 1, Position2: start2 - 1}] = 1

	for i := 0; i < endAt; i++ {
		for state, count := range scores[i] {
			if state.Score1 >= 21 && state.Score2 >= 21 {
				continue
			}

			for r1, times1 := range allRolls {
				news := State{Position1: (state.Position1 + r1) % 10}
				news.Score1 = state.Score1 + news.Position1 + 1

				if news.Score1 >= 21 {
					scores[news.Score1][news] += count * times1
				} else {
					for r2, times2 := range allRolls {
						news.Position2 = (state.Position2 + r2) % 10
						news.Score2 = state.Score2 + news.Position2 + 1

						scores[util.MaxInt(news.Score1, news.Score2)][news] += count * times1 * times2
					}
				}
			}
		}
	}

	var winner1, winner2 int
	for _, score := range scores {
		for s, count := range score {
			if s.Score1 >= 21 {
				winner1 += count
			} else if s.Score2 >= 21 {
				winner2 += count
			}
		}
	}

	return util.MaxInt(winner1, winner2)
}
