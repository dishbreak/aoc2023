package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("inputs/day10.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
}

func part1(r io.Reader) int {
	space, start := parse(r)
	return longestPt(space, start)

}
