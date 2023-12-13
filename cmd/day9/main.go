package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) int {
	acc := 0

	for s := bufio.NewScanner(r); s.Scan(); {
		line := s.Text()
		acc += Extrapolate(toSlice(line))
	}

	return acc
}

func part2(r io.Reader) int {
	acc := 0

	for s := bufio.NewScanner(r); s.Scan(); {
		line := s.Text()
		acc += Extrapolate(toSlice(line), InReverse)
	}

	return acc
}

type ExtrapolateConfig struct {
	idx  int
	mult int
}

type ExtrapolateOption func(*ExtrapolateConfig)

func InReverse(e *ExtrapolateConfig) {
	e.idx = 0
	e.mult = -1
}

func Extrapolate(seq []int, opts ...ExtrapolateOption) int {
	cfg := ExtrapolateConfig{
		idx:  len(seq) - 1,
		mult: 1,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	if isZeroes(seq) {
		return 0
	}

	nSeq := make([]int, len(seq)-1)

	for i := 1; i < len(seq); i++ {
		nSeq[i-1] = seq[i] - seq[i-1]
	}

	delta := seq[cfg.idx]

	return (Extrapolate(nSeq, opts...) * cfg.mult) + delta
}

func isZeroes(seq []int) bool {
	for _, val := range seq {
		if val != 0 {
			return false
		}
	}

	return true
}

func toSlice(s string) []int {
	pts := strings.Fields(s)
	result := make([]int, len(pts))

	for i, pt := range pts {
		result[i], _ = strconv.Atoi(pt)
	}

	return result
}
