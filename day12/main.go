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

var (
	startNode = Node("start")
	endNode   = Node("end")
)

func main() {
	g := make(Graph)
	for _, l := range util.ReadInput(in, "\n") {
		s := strings.Split(l, "-")
		g[Node(s[0])] = append(g[Node(s[0])], Node(s[1]))
		g[Node(s[1])] = append(g[Node(s[1])], Node(s[0]))
	}

	fmt.Println(g)
	fmt.Printf("first: %d\n", solve(g, 1))
	fmt.Printf("first: %d\n", solve(g, 2))
}

type Graph map[Node][]Node

func (g Graph) Next(n Node) []Node {
	return g[n]
}

type Node string

func (n Node) IsBig() bool {
	return strings.ToUpper(string(n)) == string(n)
}

type Visited struct {
	max int
	m   map[Node]int
}

func NewVisited(max int) Visited {
	return Visited{max: max, m: make(map[Node]int)}
}

func (v Visited) Add(n Node) {
	v.m[n] += 1
}

func (v Visited) Visited(n Node) bool {
	if n.IsBig() {
		return false
	}

	val := v.m[n]
	if val == 0 {
		return false
	}

	if val > 0 && (n == startNode || n == endNode) {
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
	cp := make(map[Node]int, len(v.m))
	for k, v := range v.m {
		cp[k] = v
	}

	return Visited{m: cp, max: v.max}
}

func solve(graph Graph, max int) int {
	paths := recurse(NewVisited(max), startNode, graph)
	for _, p := range paths {
		for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
			p[i], p[j] = p[j], p[i]
		}
	}
	return len(paths)
}

func recurse(visited Visited, node Node, graph Graph) [][]Node {
	if node == endNode {
		return [][]Node{{node}}
	}

	if visited.Visited(node) {
		return nil
	}

	if !node.IsBig() {
		visited.Add(node)
	}
	var paths [][]Node
	for _, next := range graph.Next(node) {
		ps := recurse(visited.Copy(), next, graph)
		if ps != nil {
			for _, p := range ps {
				paths = append(paths, append(p, node))
			}
		}
	}

	return paths
}
