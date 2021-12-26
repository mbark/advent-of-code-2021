package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("first: %d\n", solve(largeToSmall))
	fmt.Printf("first: %d\n", solve(smallToLarge))
}

var (
	largeToSmall = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	smallToLarge = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func solve(digitOrder []int) int {
	var min, max int
	for i := 0; i < 14; i++ {
		min += util.PowInt(10, i)
		max += 9 * util.PowInt(10, i)
	}

	found := findDigits(digitOrder, zBits{}, steps)

	var sb strings.Builder
	for i := len(found) - 1; i >= 0; i-- {
		sb.WriteString(strconv.Itoa(found[i]))
	}
	num, _ := strconv.Atoi(sb.String())

	return num
}

func findDigits(digitOrder []int, z zBits, steps []vars) []int {
	if len(steps) == 0 {
		return []int{}
	}

	for _, d := range digitOrder {
		step, nextSteps := steps[0], steps[1:]
		newz := doStep(step, d, z)
		if step.div == 26 && len(newz) >= len(z) {
			// we should always pop a value at div = 26
			continue
		}

		val := findDigits(digitOrder, newz, nextSteps)
		if val != nil {
			return append(val, d)
		}
	}

	return nil
}

type vars struct {
	chk int
	div int
	add int
}

func doStep(v vars, inp int, fromz zBits) zBits {
	z := make(zBits, len(fromz))
	copy(z, fromz)

	if v.div == 1 {
		if z.Last()+v.chk != inp {
			z.Push(inp + v.add)
		}
	} else {
		if z.Last()+v.chk != inp {
			z.ReplaceLast(inp + v.add)
		} else {
			z.Pop()
		}
	}

	return z
}

type zBits []int

func (z zBits) Last() int {
	if len(z) == 0 {
		return 0
	}

	return z[len(z)-1]
}

func (z *zBits) Pop() int {
	oldz := *z
	newz, last := oldz[:len(oldz)-1], oldz[len(oldz)-1]
	*z = newz
	return last
}

func (z *zBits) Push(v int) {
	*z = append(*z, v)
}

func (z *zBits) ReplaceLast(v int) {
	(*z)[len(*z)-1] = v
}

var steps = []vars{
	{chk: 12, add: 7, div: 1},
	{chk: 12, add: 8, div: 1},
	{chk: 13, add: 2, div: 1},
	{chk: 12, add: 11, div: 1},
	{chk: -3, add: 6, div: 26},
	{chk: 10, add: 12, div: 1},
	{chk: 14, add: 14, div: 1},
	{chk: -16, add: 13, div: 26},
	{chk: 12, add: 15, div: 1},
	{chk: -8, add: 10, div: 26},
	{chk: -12, add: 6, div: 26},
	{chk: -7, add: 10, div: 26},
	{chk: -6, add: 8, div: 26},
	{chk: -11, add: 5, div: 26},
}

/*
	Manually verified.

   PUSH input[0] + 7
   PUSH input[1] + 8
   PUSH input[2] + 2
   PUSH input[3] + 11
   POP. Must have input[4] == popped_value - 3
   PUSH input[5] + 12
   PUSH input[6] + 14
   POP. Must have input[7] == popped_value - 16
   PUSH input[8] + 15
   POP. Must have input[9] == popped_value - 8
   POP. Must have input[10] == popped_value - 12
   POP. Must have input[11] == popped_value - 7
   POP. Must have input[12] == popped_value - 6
   POP. Must have input[13] == popped_value - 11

    input[4] = input[3] + 8
    input[7] = input[6] - 2
    input[9] = input[8] + 7
    input[10] = input[5]
    input[11] = input[2] - 5
    input[12] = input[1] + 2
    input[13] = input[0] - 4

	input[3] = 1
	input[4] = 9
	input[6] = 3, 9
	input[7] = 1, 7
	input[8] = 1, 2
	input[9] = 8, 9
	input[5] = 1, 9
	input[10] = 1, 9
	input[2] = 6, 9
	input[11] = 1, 4
	input[1] = 1, 7
	input[12] = 3, 9
	input[0] = 5, 9
	input[13] = 1, 5

	sorted:
	input[0] = 5, 9
	input[1] = 1, 7
	input[2] = 6, 9
	input[3] = 1
	input[4] = 9
	input[5] = 1, 9
	input[6] = 3, 9
	input[7] = 1, 7
	input[8] = 1, 2
	input[9] = 8, 9
	input[10] = 1, 9
	input[11] = 1, 4
	input[12] = 3, 9
	input[13] = 1, 5


	max: 97919997299495
	min: 51619131181131
*/
