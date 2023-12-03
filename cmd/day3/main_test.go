package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPart1(t *testing.T) {
	r := strings.NewReader(sampleInput)
	assert.Equal(t, 4361, part1(r))
}
