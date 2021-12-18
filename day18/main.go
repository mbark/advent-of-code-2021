package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
	"strconv"
)

const (
	testData = `
[[[[4,3],4],4],[7,[[8,4],9]]]
[1,1]
`
	testData1 = `
[1,1]
[2,2]
[3,3]
[4,4]
`

	testData2 = `
[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]
`

	testData3 = `
[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
`

	testData4 = `
[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]
`

	homework = `
[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
`

	in = `
[3,[5,[7,[3,9]]]]
[[[[7,0],0],[2,[2,8]]],[[[7,8],1],3]]
[[[[2,7],0],7],4]
[[2,1],[9,0]]
[[[[7,1],[3,2]],[[9,8],5]],[2,7]]
[[[8,9],[[8,7],0]],[[[8,7],[6,3]],[[1,7],[8,9]]]]
[[8,6],[[9,[1,7]],[6,[3,9]]]]
[[2,[[5,6],6]],[[4,[5,9]],[3,[4,5]]]]
[[[[2,0],[1,1]],[6,6]],[[1,9],[[2,7],[6,8]]]]
[[[4,6],[[6,3],[3,9]]],[[[2,6],[6,1]],[[9,9],[1,5]]]]
[[[4,[3,1]],3],6]
[[0,[[5,2],8]],[1,[9,[4,3]]]]
[[[[8,6],[2,1]],[2,[8,6]]],[[[7,1],[3,9]],0]]
[[[[4,7],[2,7]],[[8,9],2]],[[[2,4],[7,2]],[3,7]]]
[[5,[2,2]],[[1,6],[[9,1],[5,0]]]]
[[5,[[1,2],[6,4]]],[6,8]]
[[[5,[1,7]],7],[7,[8,1]]]
[[1,9],[[0,3],[[6,7],[2,4]]]]
[1,[7,[[0,6],0]]]
[[[[5,7],9],[[3,2],7]],[[5,1],[9,9]]]
[[[[0,4],[9,6]],[[8,3],[7,4]]],[7,[6,2]]]
[[[[1,6],0],[[8,0],[3,4]]],[[3,[0,3]],4]]
[4,[[7,8],[4,[9,7]]]]
[[[2,[3,7]],5],[0,[9,9]]]
[[[2,0],[[5,8],[7,6]]],[[9,[6,2]],[3,2]]]
[[[3,1],3],[[[3,7],6],[9,8]]]
[[7,[[2,5],5]],[5,[3,[4,5]]]]
[[[6,7],6],[2,[[9,3],9]]]
[[[[5,6],7],[[3,2],5]],[[9,[4,3]],[3,8]]]
[0,7]
[[[4,6],[2,9]],[[[7,6],[5,1]],7]]
[[0,5],[[1,[4,1]],[[7,3],9]]]
[[[2,[3,8]],5],[[[5,9],8],[7,0]]]
[[[6,[8,6]],[[3,6],7]],[[2,1],[6,[7,5]]]]
[[2,[[6,3],[8,9]]],[[[5,6],4],[[7,0],1]]]
[[[[7,1],[5,6]],8],[[[8,9],4],[8,3]]]
[[[9,2],[1,0]],0]
[[5,[5,[8,5]]],4]
[[3,[5,[4,9]]],3]
[[8,[[7,7],6]],5]
[[4,[[5,1],1]],[1,[1,[9,8]]]]
[[[7,[3,6]],[[2,8],[4,7]]],[[[8,8],[4,0]],[2,4]]]
[[[[3,6],3],[0,9]],2]
[[2,8],[[8,[8,6]],[[1,1],[4,5]]]]
[[2,[1,[1,0]]],[[[6,2],[7,4]],[[7,1],6]]]
[3,[8,[7,[8,6]]]]
[[1,0],[[[0,4],[0,5]],[1,5]]]
[[[[5,0],4],[[7,8],[8,8]]],[[1,7],0]]
[1,[[[4,1],7],[6,[9,0]]]]
[[[1,8],2],[[5,5],[8,5]]]
[[4,[9,[0,6]]],[[[8,9],[4,5]],4]]
[[[[5,4],[1,7]],[[3,1],[7,9]]],[[[0,8],[4,7]],[[5,9],6]]]
[[[[8,0],9],4],[[7,[1,3]],5]]
[[[[5,0],6],[[6,1],8]],[[9,1],7]]
[[9,[6,[8,8]]],[7,[[7,1],6]]]
[[[5,[1,5]],[3,[4,2]]],[[[5,2],7],[[6,9],[2,8]]]]
[[[5,[5,5]],[5,7]],[4,[[2,9],7]]]
[[[[0,4],0],[[0,6],[3,0]]],[0,[[8,1],2]]]
[[[7,[4,6]],[[7,2],[4,6]]],[[[9,3],[4,9]],6]]
[[6,7],7]
[[[4,1],[8,[1,5]]],[[4,6],0]]
[[[4,[5,5]],5],[[0,[2,7]],[1,1]]]
[[[[0,1],3],[6,7]],[4,7]]
[[4,[6,4]],[[[9,8],1],[9,3]]]
[[[4,9],0],[[[7,0],[0,9]],[1,[1,0]]]]
[[[7,9],[[9,5],[6,9]]],[[0,[3,0]],[0,[5,9]]]]
[9,[[0,0],[[1,9],9]]]
[[[5,[0,5]],[[9,8],[9,5]]],[[0,[2,5]],7]]
[[[[5,8],6],9],[[[2,7],7],[[7,8],5]]]
[[8,[[4,7],6]],2]
[[[[7,1],[9,0]],[9,[1,7]]],[[8,[6,7]],[2,5]]]
[[4,[2,9]],8]
[[[[7,6],[5,3]],[5,[9,7]]],[[6,[8,1]],[[6,4],9]]]
[[7,[[7,8],4]],[[1,3],[4,[9,7]]]]
[[[6,[6,7]],[[2,8],3]],[7,[6,[0,3]]]]
[[9,8],[[0,[4,8]],[[9,1],1]]]
[[[[4,0],[5,9]],7],[6,[[5,9],[9,6]]]]
[[8,1],[1,[9,[8,3]]]]
[[[1,[5,1]],[6,7]],[[5,9],[2,[6,7]]]]
[[[3,7],[[7,8],1]],[[0,[6,3]],[8,0]]]
[[5,[[9,3],[1,2]]],7]
[[[1,[9,9]],3],[[6,4],[4,1]]]
[[6,[1,[3,6]]],[2,9]]
[[2,[0,2]],[5,[[9,4],[5,0]]]]
[[4,[[3,1],[7,0]]],[[9,1],[[5,5],[6,7]]]]
[[3,[[7,1],[3,4]]],[7,[9,[9,4]]]]
[[9,9],[[5,4],[[9,7],4]]]
[[[5,1],8],[[6,7],9]]
[[[0,[9,5]],[4,3]],[3,2]]
[[[6,[4,1]],[[8,7],[5,3]]],[[[1,2],5],[[9,2],5]]]
[[[[7,4],[9,0]],[[1,8],[2,9]]],[[5,[1,9]],[4,0]]]
[[[4,[3,8]],[[3,3],[2,8]]],[[[1,3],9],[[8,5],6]]]
[[[[6,4],[7,9]],[[7,6],8]],[7,[9,8]]]
[[7,[3,5]],7]
[[[[5,0],[2,3]],[3,7]],[[4,[6,3]],[7,[4,4]]]]
[[6,[3,[7,6]]],[[[5,8],[8,1]],[3,[1,5]]]]
[[8,[9,[5,2]]],2]
[[1,[5,4]],[[7,[8,0]],8]]
[[[[2,7],4],3],[[1,4],[8,4]]]
[3,[9,2]]
`

	twoTest = `
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
`
)

func main() {
	var pairs []*Pair
	var pairs2 []*Pair
	for _, l := range util.ReadInput(in, "\n") {
		active := new(Pair)

		for _, c := range l {
			switch c {
			case '[':
				active.Left = &Pair{Up: active}
				active = active.Left
			case ']':
				active = active.Up
			case ',':
				active.Up.Right = &Pair{Up: active.Up}
				active = active.Up.Right
			default:
				active.Value = util.Str2Int(string(c))
			}
		}

		pairs = append(pairs, active)
		pairs2 = append(pairs2, active.Copy(nil))
	}

	fmt.Printf("first: %d\n", first(pairs).Magnitude())
	p := second(pairs2)
	fmt.Printf("second: %d\n", p.Magnitude())
}

func first(pairs []*Pair) *Pair {
	p := pairs[0]

	for _, pn := range pairs[1:] {
		p = add(p, pn)
		reduce(p)
	}

	return p
}

func second(pairs []*Pair) *Pair {
	var pmax int
	var maxPair *Pair

	for i, p1 := range pairs {
		for j, p2 := range pairs {
			if i == j {
				continue
			}

			a := add(p1.Copy(nil), p2.Copy(nil))
			reduce(a)
			if mag := a.Magnitude(); mag > pmax {
				pmax = mag
				maxPair = a
			}
		}
	}

	return maxPair
}

type Pair struct {
	Value int
	Left  *Pair
	Right *Pair
	Up    *Pair
}

func (p *Pair) Copy(up *Pair) *Pair {
	if p == nil {
		return nil
	}
	if p.IsRegular() {
		return &Pair{Value: p.Value, Up: up}
	}

	pnew := &Pair{Up: up}
	pnew.Left = p.Left.Copy(pnew)
	pnew.Right = p.Right.Copy(pnew)
	return pnew
}

func (p *Pair) IsRegular() bool {
	if p == nil {
		return false
	}

	return p.Left == nil && p.Right == nil
}

func (p *Pair) Magnitude() int {
	if p.IsRegular() {
		return p.Value
	}

	return 3*p.Left.Magnitude() + 2*p.Right.Magnitude()
}

func (p Pair) String() string {
	if p.IsRegular() {
		return strconv.Itoa(p.Value)
	}

	return fmt.Sprintf("[%s,%s]", p.Left.String(), p.Right.String())
}

func add(p1, p2 *Pair) *Pair {
	top := &Pair{Left: p1, Right: p2}
	p1.Up, p2.Up = top, top
	return top
}

func reduce(pair *Pair) {
	for {
		reduced := reduceStep(pair, 1, true) || reduceStep(pair, 1, false)
		if !reduced {
			break
		}
	}
}

func reduceStep(pair *Pair, depth int, explode bool) bool {
	if pair == nil {
		return false
	}

	switch {
	case explode && depth > 4 && pair.Left.IsRegular() && pair.Right.IsRegular():
		pair.explode()
		return true
	case !explode && pair.Value >= 10 && pair.IsRegular():
		pair.split()
		return true
	}

	if ok := reduceStep(pair.Left, depth+1, explode); ok {
		return true
	}

	if ok := reduceStep(pair.Right, depth+1, explode); ok {
		return true
	}

	return false
}

func (e *Pair) explode() {
	left := findLeft(e.Up, e)
	right := findRight(e.Up, e)

	if left != nil {
		left.Value += e.Left.Value
	}
	if right != nil {
		right.Value += e.Right.Value
	}

	e.Left, e.Right = nil, nil
	e.Value = 0
}

func findLeft(p *Pair, from *Pair) *Pair {
	if p == nil || p.IsRegular() {
		return p
	}

	if p.Left == from {
		return findLeft(p.Up, p)
	}
	if p.Right == from {
		return findLeft(p.Left, p)
	}

	return findLeft(p.Right, p)
}

func findRight(p *Pair, from *Pair) *Pair {
	if p == nil || p.IsRegular() {
		return p
	}

	if p.Right == from {
		return findRight(p.Up, p)
	}
	if p.Left == from {
		return findRight(p.Right, p)
	}

	return findRight(p.Left, p)
}

func (e *Pair) split() {
	vall, valr := e.Value/2, e.Value/2+e.Value%2
	e.Left = &Pair{Value: vall, Up: e}
	e.Right = &Pair{Value: valr, Up: e}
}
