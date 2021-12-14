package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
	"strings"
)

var in = `
PPFCHPFNCKOKOSBVCFPP

VC -> N
SC -> H
CK -> P
OK -> O
KV -> O
HS -> B
OH -> O
VN -> F
FS -> S
ON -> B
OS -> H
PC -> B
BP -> O
OO -> N
BF -> K
CN -> B
FK -> F
NP -> K
KK -> H
CB -> S
CV -> K
VS -> F
SF -> N
KB -> H
KN -> F
CP -> V
BO -> N
SS -> O
HF -> H
NN -> F
PP -> O
VP -> H
BB -> K
VB -> N
OF -> N
SH -> S
PO -> F
OC -> S
NS -> C
FH -> N
FP -> C
SO -> P
VK -> C
HP -> O
PV -> S
HN -> K
NB -> C
NV -> K
NK -> B
FN -> C
VV -> N
BN -> N
BH -> S
FO -> V
PK -> N
PS -> O
CO -> K
NO -> K
SV -> C
KO -> V
HC -> B
BC -> N
PB -> C
SK -> S
FV -> K
HO -> O
CF -> O
HB -> P
SP -> N
VH -> P
NC -> K
KC -> B
OV -> P
BK -> F
FB -> F
FF -> V
CS -> F
CC -> H
SB -> C
VO -> V
VF -> O
KP -> N
HV -> H
PF -> H
KH -> P
KS -> S
BS -> H
PH -> S
SN -> K
HK -> P
FC -> N
PN -> S
HH -> N
OB -> P
BV -> S
KF -> N
OP -> H
NF -> V
CH -> K
NH -> P
`

var testData = `
NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`

func main() {
	s := util.ReadInput(in, "\n\n")
	start := s[0]

	m := make(Mapping)
	for _, l := range strings.Split(s[1], "\n") {
		s := strings.Split(l, " -> ")
		m[Pair{A: s[0][0], B: s[0][1]}] = To(s[1][0])
	}

	fmt.Printf("first %d\n", polymerize(start, m, 10))
	fmt.Printf("second %d\n", polymerize(start, m, 40))
}

type Pair struct {
	A uint8
	B uint8
}

func (p Pair) String() string {
	return fmt.Sprintf("%s", string([]uint8{p.A, p.B}))
}

type To uint8

func (t To) String() string {
	return string(t)
}

type Mapping map[Pair]To

type Pairs map[Pair]int

func polymerize(start string, mapping Mapping, loops int) int {
	starting, ending := To(start[0]), To(start[len(start)-1])
	pairs := make(Pairs)
	for i := 1; i < len(start); i++ {
		a, b := start[i-1], start[i]
		pairs[Pair{A: a, B: b}] += 1
	}

	for i := 0; i < loops; i++ {
		next := make(Pairs)
		for p, c := range pairs {
			to, ok := mapping[Pair{A: p.A, B: p.B}]
			if ok {
				t := uint8(to)
				next[Pair{A: p.A, B: t}] += c
				next[Pair{A: t, B: p.B}] += c
			} else {
				next[p] += c
			}
		}

		pairs = next
	}

	var total int
	counts := make(map[To]int)
	for p, c := range pairs {
		counts[To(p.A)] += c
		counts[To(p.B)] += c
		total += c
	}

	var most, least To
	for a, c := range counts {
		if most == 0 || c > counts[most] {
			most = a
		}
		if least == 0 || c < counts[least] {
			least = a
		}
	}

	highest := counts[most]
	lowest := counts[least]

	if most == starting {
		highest += 1
	}
	if most == ending {
		highest += 1
	}
	if least == starting {
		lowest += 1
	}
	if least == ending {
		lowest += 1
	}

	return highest/2 - lowest/2
}
