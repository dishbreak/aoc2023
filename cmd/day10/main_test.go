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
		{
			input: `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`,
			result: 4,
		},
		{
			input: `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`,
			result: 8,
		},
	}

	for i, tc := range testCases {
		name := fmt.Sprint("test case ", i)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, part1(strings.NewReader(tc.input)))
		})
	}
}
