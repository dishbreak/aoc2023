package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestPart1(t *testing.T) {
	r := strings.NewReader(sampleInput)
	assert.Equal(t, 8, part1(r))
}

func TestParser(t *testing.T) {
	gs := parseGameSummary("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")
	expected := gameSummary{
		id: 4,
		rounds: []map[string]int{
			{
				"green": 1,
				"red":   3,
				"blue":  6,
			},
			{
				"green": 3,
				"red":   6,
			},
			{
				"green": 3,
				"blue":  15,
				"red":   14,
			},
		},
	}

	assert.Equal(t, expected, gs)
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(sampleInput)
	assert.Equal(t, 2286, part2(r))
}
