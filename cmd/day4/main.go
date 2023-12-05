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
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
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

func part2(r io.Reader) int {
	cards := make([]card.Card, 0)

	for s := bufio.NewScanner(r); s.Scan(); {
		cards = append(cards, card.FromString(s.Text()))
	}

	numCards := len(cards)

	counts := make([]int, numCards+1)
	for i := 1; i < len(counts); i++ {
		counts[i] = 1
	}

	for _, card := range cards {
		m := card.Matches()

		for j := 0; j < m; j++ {
			wonCardId := card.Id + j + 1
			counts[wonCardId] += counts[card.Id]
		}
	}

	acc := 0
	for _, ct := range counts {
		acc += ct
	}

	return acc
}
