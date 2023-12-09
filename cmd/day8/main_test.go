package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	type testCase struct {
		input  string
		result int
	}

	testCases := []testCase{
		{`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`, 2},
		{`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`, 6},
	}

	for i, tc := range testCases {
		name := fmt.Sprint("test case ", i)
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			assert.Equal(t, tc.result, part1(r))
		})
	}

}

func TestPart2(t *testing.T) {
	testInput := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	assert.Equal(t, 6, part2(strings.NewReader(testInput)))
}
