package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
/*
A 12
K 12
Q 11
J 10
T 9
9 8
8 7
7 6
6 5
5 4
4 3
3 2
2 1
*/

func main() {
	f, err := os.Open("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

type Hand int

const (
	FiveOfAKind = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func getHandType(cards string) Hand {
	pairs := make(map[rune]int)
	var pairCount [6]int

	for _, c := range cards {
		pairs[c]++
	}

	for _, v := range pairs {
		pairCount[v]++
	}

	if pairCount[5] == 1 {
		return FiveOfAKind
	}

	if pairCount[4] == 1 {
		return FourOfAKind
	}

	if pairCount[3] == 1 && pairCount[2] == 1 {
		return FullHouse
	}

	if pairCount[3] == 1 {
		return ThreeOfAKind
	}

	if pairCount[2] == 2 {
		return TwoPair
	}

	if pairCount[2] == 1 {
		return OnePair
	}

	return HighCard
}

func getHandTypeWithJokers(cards string) Hand {
	pairs := make(map[rune]int)
	var pairCount [6]int

	for _, c := range cards {
		pairs[c]++
	}

	jokerCt := pairs['J']
	delete(pairs, 'J')

	highCard := ' '
	maxVal := -1

	for card, ct := range pairs {
		if ct > maxVal {
			maxVal = ct
			highCard = card
		}
	}

	pairs[highCard] += jokerCt

	for _, v := range pairs {
		pairCount[v]++
	}

	if pairCount[5] == 1 {
		return FiveOfAKind
	}

	if pairCount[4] == 1 {
		return FourOfAKind
	}

	if pairCount[3] == 1 && pairCount[2] == 1 {
		return FullHouse
	}

	if pairCount[3] == 1 {
		return ThreeOfAKind
	}

	if pairCount[2] == 2 {
		return TwoPair
	}

	if pairCount[2] == 1 {
		return OnePair
	}

	return HighCard
}

type CamelCardHard struct {
	Cards string
	Bid   int
}

func parse(r io.Reader) []CamelCardHard {
	result := make([]CamelCardHard, 0)

	s := bufio.NewScanner(r)
	for s.Scan() {
		pts := strings.Fields(s.Text())
		ch := CamelCardHard{Cards: pts[0]}
		ch.Bid, _ = strconv.Atoi(pts[1])
		result = append(result, ch)
	}

	return result
}

func part1(r io.Reader) int {
	hands := parse(r)

	var strength [255]int
	strength['A'] = 13
	strength['K'] = 12
	strength['Q'] = 11
	strength['J'] = 10
	strength['T'] = 9
	strength['9'] = 8
	strength['8'] = 7
	strength['7'] = 6
	strength['6'] = 5
	strength['5'] = 4
	strength['4'] = 3
	strength['3'] = 2
	strength['2'] = 1

	sort.Slice(hands, func(i, j int) bool {
		// is one less than the other?
		one := hands[i]
		other := hands[j]

		oneHand := getHandType(one.Cards)
		otherHand := getHandType(other.Cards)

		if oneHand < otherHand {
			return false
		}

		if oneHand > otherHand {
			return true
		}

		for k := range one.Cards {
			if one.Cards[k] == other.Cards[k] {
				continue
			}
			if strength[one.Cards[k]] < strength[other.Cards[k]] {
				return true
			}
			return false
		}
		return true
	})

	acc := 0
	for i, hand := range hands {
		acc += (i + 1) * hand.Bid
	}

	return acc
}

func part2(r io.Reader) int {
	hands := parse(r)

	var strength [255]int
	strength['A'] = 13
	strength['K'] = 12
	strength['Q'] = 11
	strength['J'] = -1
	strength['T'] = 9
	strength['9'] = 8
	strength['8'] = 7
	strength['7'] = 6
	strength['6'] = 5
	strength['5'] = 4
	strength['4'] = 3
	strength['3'] = 2
	strength['2'] = 1

	sort.Slice(hands, func(i, j int) bool {
		// is one less than the other?
		one := hands[i]
		other := hands[j]

		oneHand := getHandTypeWithJokers(one.Cards)
		otherHand := getHandTypeWithJokers(other.Cards)

		if oneHand < otherHand {
			return false
		}

		if oneHand > otherHand {
			return true
		}

		for k := range one.Cards {
			if one.Cards[k] == other.Cards[k] {
				continue
			}
			if strength[one.Cards[k]] < strength[other.Cards[k]] {
				return true
			}
			return false
		}
		return true
	})

	acc := 0
	for i, hand := range hands {
		acc += (i + 1) * hand.Bid
	}

	return acc
}
