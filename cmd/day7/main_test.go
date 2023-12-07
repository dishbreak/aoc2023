package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandType(t *testing.T) {
	type testCase struct {
		input  string
		result Hand
	}

	/*
			Five of a kind, where all five cards have the same label: AAAAA
		Four of a kind, where four cards have the same label and one card has a different label: AA8AA
		Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
		Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
		Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
		One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
		High card, where all cards' labels are distinct: 23456

	*/
	testCases := []testCase{
		{"AAAAA", FiveOfAKind},
		{"AA8AA", FourOfAKind},
		{"23332", FullHouse},
		{"TTT98", ThreeOfAKind},
		{"23432", TwoPair},
		{"A23A4", OnePair},
		{"23456", HighCard},
	}

	for i, tc := range testCases {
		name := fmt.Sprint("test case ", i)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, getHandType(tc.input))
		})
	}
}

func TestPart1(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	assert.Equal(t, 6440, part1(strings.NewReader(input)))
}

func TestPart2(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	assert.Equal(t, 5905, part2(strings.NewReader(input)))
}
