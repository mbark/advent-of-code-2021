package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
	"strings"
)

var testData = `
start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

var testData1 = `
fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`

var in = `
vn-DD
qm-DD
MV-xy
end-xy
KG-end
end-kw
qm-xy
start-vn
MV-vn
vn-ko
lj-KG
DD-xy
lj-kh
lj-MV
ko-MV
kw-qm
qm-MV
lj-kw
VH-lj
ko-qm
ko-start
MV-start
DD-ko
`

func main() {
	g := make(Graph)
	nodeMap := make(map[string]int)
	n := 1

	var start Node
	for _, l := range util.ReadInput(in, "\n") {
		s := strings.Split(l, "-")
		n1, n2 := nodeMap[s[0]], nodeMap[s[1]]
		if n1 == 0 {
			n1 = n
			n++
			nodeMap[s[0]] = n1
		}
		if n2 == 0 {
			n2 = n
			n++
			nodeMap[s[1]] = n2
		}

		node1 := Node{n: n1, s: s[0]}
		node2 := Node{n: n2, s: s[1]}

		g[node1] = append(g[node1], node2)
		g[node2] = append(g[node2], node1)

		if node1.IsStart() {
			start = node1
		}
		if node2.IsStart() {
			start = node2
		}
	}

	fmt.Printf("first: %d\n", solve(g, start, 1))
	fmt.Printf("first: %d\n", solve(g, start, 2))
}

type Graph map[Node][]Node

func (g Graph) Next(n Node) []Node {
	return g[n]
}

type Node struct {
	n int
	s string
}

func (n Node) IsBig() bool {
	return strings.ToUpper(n.s) == n.s
}

func (n Node) IsStart() bool {
	return n.s == "start"
}

func (n Node) IsEnd() bool {
	return n.s == "end"
}

type Visited struct {
	max int
	m   []int
}

func NewVisited(max int, graph Graph) Visited {
	var maxVal int
	for n := range graph {
		if n.n > maxVal {
			maxVal = n.n
		}
	}

	return Visited{max: max, m: make([]int, maxVal)}
}

func (v Visited) Add(n Node) {
	v.m[n.n] += 1
}

func (v Visited) Visited(n Node) bool {
	if n.IsBig() {
		return false
	}

	val := v.m[n.n]
	if val == 0 {
		return false
	}

	if val > 0 && (n.IsStart() || n.IsEnd()) {
		return true
	}

	if v.max == 1 {
		return val >= v.max
	}

	var visitedTwice bool
	for _, val := range v.m {
		if val > 1 {
			visitedTwice = true
		}
	}

	return visitedTwice
}

func (v Visited) Copy() Visited {
	cp := make([]int, len(v.m))
	for k, v := range v.m {
		cp[k] = v
	}

	return Visited{m: cp, max: v.max}
}

func solve(graph Graph, start Node, max int) int {
	paths := recurse(NewVisited(max, graph), start, graph)
	return paths
}

func recurse(visited Visited, node Node, graph Graph) int {
	if !node.IsBig() {
		visited = visited.Copy()
		visited.Add(node)
	}

	var count int
	for _, next := range graph.Next(node) {
		if next.IsEnd() {
			count += 1
			continue
		}

		if visited.Visited(next) {
			continue
		}

		count += recurse(visited, next, graph)
	}

	return count
}
