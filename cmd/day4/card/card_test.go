package card

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	expected := Card{
		Id:      3,
		numbers: []int{69, 82, 63, 72, 16, 21, 14, 1},
	}

	expected.winner[1] = true
	expected.winner[21] = true
	expected.winner[53] = true
	expected.winner[59] = true
	expected.winner[44] = true

	assert.Equal(t, expected, FromString("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"))
}

func TestScore(t *testing.T) {
	type testCase struct {
		input string
		score int
	}

	testCases := []testCase{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			c := FromString(tc.input)
			assert.Equal(t, tc.score, c.Score())
		})
	}
}

func TestMatches(t *testing.T) {
	type testCase struct {
		input   string
		matches int
	}

	testCases := []testCase{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 4},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			c := FromString(tc.input)
			assert.Equal(t, tc.matches, c.Matches())
		})
	}
}
