package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	records := []raceRecord{
		{7, 9},
		{15, 40},
		{30, 200},
	}

	assert.Equal(t, 288, part1(records))
}

func TestPart2(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`

	assert.Equal(t, 71503, part2(input))
}
