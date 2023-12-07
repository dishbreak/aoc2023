package almanac

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestToLocation(t *testing.T) {
	type testCase struct {
		input, result int
	}

	testCases := []testCase{
		{79, 82},
		{14, 43},
		{55, 86},
		{13, 35},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			a := FromString(testInput)
			assert.Equal(t, tc.result, a.ToLocation(tc.input))
		})
	}
}
