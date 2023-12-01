package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	type testCase struct {
		input  string
		result int
	}

	testCases := []testCase{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			assert.Equal(t, tc.result, getCalibrationNumber(tc.input))
		})
	}
}

func TestPart2(t *testing.T) {
	type testCase struct {
		input  string
		result int
	}

	testCases := []testCase{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"ninetwonine7ninetwonend", 91},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			assert.Equal(t, tc.result, getCalibrationNumberV2(tc.input))
		})
	}
}
