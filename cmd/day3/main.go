package main

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
}

var matchPartNum = regexp.MustCompile(`\d+`)
var matchSymbol = regexp.MustCompile(`[#*\/%@$&\+\-\=]`)

func part1(r io.Reader) int {
	s := bufio.NewScanner(r)

	symbols := make(map[image.Point]byte)
	parts := make(map[image.Point]int)

	for lineNo := 0; s.Scan(); lineNo++ {
		line := s.Text()

		// read all symbols
		idxs := matchSymbol.FindAllStringIndex(line, -1)
		for _, loc := range idxs {
			symbols[image.Pt(loc[0], lineNo)] = byte(line[loc[0]])
		}

		// read all numbers
		idxs = matchPartNum.FindAllStringIndex(line, -1)
		for _, loc := range idxs {
			partNum, _ := strconv.Atoi(line[loc[0]:loc[1]])
			for i := loc[0]; i < loc[1]; i++ {
				parts[image.Pt(i, lineNo)] = partNum
			}
		}
	}

	ptNorth := image.Pt(0, -1)
	ptSouth := image.Pt(0, 1)
	ptEast := image.Pt(1, 0)
	ptWest := image.Pt(-1, 0)
	ptNorthEast := ptNorth.Add(ptEast)
	ptNorthWest := ptNorth.Add(ptWest)
	ptSouthEast := ptSouth.Add(ptEast)
	ptSouthWest := ptSouth.Add(ptWest)

	acc := 0

	// iterate over all the positions of symbols
	for p := range symbols {
		// look east and west first -- these are always seen once by the symbol
		acc += parts[p.Add(ptEast)]
		acc += parts[p.Add(ptWest)]

		/*
			with north and south, this gets messy. here are all the valid cases for north
			...
			.&.

			.1.
			.&.

			12.
			.&.

			.12
			.&.

			123
			.&.

			1.1
			.&.

			1..
			.&.

			..1
			.&.

			if there's a number directly above, that's all that can be seen by the symbol
			if there's *NOT* a number directly above, we might see numbers at both corners
		*/

		num, directlyAbove := parts[p.Add(ptNorth)]
		acc += num
		if !directlyAbove {
			acc += parts[p.Add(ptNorthEast)]
			acc += parts[p.Add(ptNorthWest)]
		}

		// same logic holds for south
		num, directlyBelow := parts[p.Add(ptSouth)]
		acc += num
		if !directlyBelow {
			acc += parts[p.Add(ptSouthEast)]
			acc += parts[p.Add(ptSouthWest)]
		}

	}
	return acc
}
