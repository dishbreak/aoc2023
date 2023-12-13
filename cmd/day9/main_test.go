package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtrapolate(t *testing.T) {
	type testCase struct {
		input  []int
		result int
		opts   []ExtrapolateOption
	}

	testCases := []testCase{
		{[]int{0, 3, 6, 9, 12, 15}, 18, []ExtrapolateOption{}},
		{[]int{1, 3, 6, 10, 15, 21}, 28, []ExtrapolateOption{}},
		{[]int{10, 13, 16, 21, 30, 45}, 68, []ExtrapolateOption{}},
		{[]int{0, 3, 6, 9, 12, 15}, -3, []ExtrapolateOption{InReverse}},
		{[]int{1, 3, 6, 10, 15, 21}, 0, []ExtrapolateOption{InReverse}},
		{[]int{10, 13, 16, 21, 30, 45}, 5, []ExtrapolateOption{InReverse}},
	}

	for i, tc := range testCases {
		name := fmt.Sprint("test case ", i)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, Extrapolate(tc.input, tc.opts...))
		})
	}
}
