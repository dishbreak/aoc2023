package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/dishbreak/aoc2023/cmd/day4/card"
)

func main() {
	f, err := os.Open("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
}

func part1(r io.Reader) int {
	s := bufio.NewScanner(r)
	acc := 0

	for s.Scan() {
		c := card.FromString(s.Text())
		acc += c.Score()
	}

	return acc
}
