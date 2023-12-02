package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) int {
	s := bufio.NewScanner(r)
	acc := 0
	for s.Scan() {
		gs := parseGameSummary(s.Text())
		if gs.isImpossible(map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		}) {
			continue
		}
		acc += gs.id
	}
	return acc
}

func part2(r io.Reader) int {
	s := bufio.NewScanner(r)
	acc := 0
	for s.Scan() {
		gs := parseGameSummary(s.Text())
		acc += gs.getPowerNeed()
	}

	return acc
}

type gameSummary struct {
	id     int
	rounds []map[string]int
}

var gameRegexp = regexp.MustCompile(`^Game (\d+): (.+)$`)

func parseGameSummary(s string) (gs gameSummary) {
	matches := gameRegexp.FindStringSubmatch(s)
	gs.id, _ = strconv.Atoi(matches[1])

	rounds := strings.Split(matches[2], ";")
	gs.rounds = make([]map[string]int, len(rounds))

	for i, round := range rounds {
		gs.rounds[i] = make(map[string]int)
		colors := strings.Split(round, ", ")
		for _, color := range colors {
			pts := strings.Fields(color)
			ct, _ := strconv.Atoi(pts[0])
			gs.rounds[i][pts[1]] = ct
		}
	}

	return
}

func (gs gameSummary) isImpossible(constraints map[string]int) bool {
	for _, round := range gs.rounds {
		for color, ct := range round {
			if ct > constraints[color] {
				return true
			}
		}
	}
	return false
}

func (gs gameSummary) getMinReqs() map[string]int {
	result := map[string]int{
		"red":   -1,
		"green": -1,
		"blue":  -1,
	}

	for _, round := range gs.rounds {
		for color, ct := range round {
			if result[color] < ct {
				result[color] = ct
			}
		}
	}

	return result
}

func (gs gameSummary) getPowerNeed() int {
	minCubes := gs.getMinReqs()

	acc := 1 // always 1 with products, 0 with sums
	for _, ct := range minCubes {
		acc *= ct
	}

	return acc
}
