package util

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func ReadInput(in, splitBy string) []string {
	trimmed := strings.Trim(in, "\n")
	return strings.Split(trimmed, splitBy)
}

func NumberList(in string, separator string) []int {
	var list []int
	for _, s := range strings.Split(in, separator) {
		if s == "" {
			continue
		}
		list = append(list, Str2Int(s))
	}

	return list
}

func Str2Int(in string) int {
	i, _ := strconv.Atoi(in)
	return i
}

func NewBoolMatrix(width, height int) map[int]map[int]bool {
	m := make(map[int]map[int]bool, width)
	for i := 0; i < width; i++ {
		m[i] = make(map[int]bool, height)
		for j := 0; j < height; j++ {
			m[i][j] = false
		}
	}

	return m
}

func WithTime() func() {
	now := time.Now()

	return func() { fmt.Printf("time taken: %v\n", time.Now().Sub(now)) }
}

func WithProfiling() func() {
	flag.Parse()
	if *cpuprofile == "" {
		return func() {}
	}

	f, err := os.Create(*cpuprofile)
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(f)
	return pprof.StopCPUProfile
}

func MaxInt(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func MinInt(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func Btoi(s string) int64 {
	i, _ := strconv.ParseInt(s, 2, 64)
	return i
}

