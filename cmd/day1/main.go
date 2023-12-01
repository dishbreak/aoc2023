package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) int {
	s := bufio.NewScanner(r)

	acc := 0
	for s.Scan() {
		line := s.Text()
		acc += getCalibrationNumber(line)
	}
	return acc
}

func getCalibrationNumber(s string) int {

	findDigit := func(start, inc int) int {
		for i := start; i != -1 || i != len(s); i += inc {
			if c := s[i]; c >= '0' && c <= '9' {
				return int(s[i] - '0')
			}
		}

		return -1
	}

	firstDigit := findDigit(0, 1)          // left-hand side
	secondDigit := findDigit(len(s)-1, -1) // right-hand side

	return firstDigit*10 + secondDigit
}

func getCalibrationNumberV2(s string) int {

	lut := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	type token struct {
		start, end, value int
	}

	getTokens := func(s, sub string, value int) []token {
		acc := make([]token, 0)
		j := 0

		// repeatedly use the strings.Index() method until there's nothing found for the given token.
		for {
			i := strings.Index(s[j:], sub)
			if i == -1 {
				break
			}
			acc = append(acc, token{
				value: value,
				start: j + i,
				end:   j + i + len(sub) - 1,
			})
			j += i + 1
		}
		return acc
	}

	tokens := make([]token, 0)
	for k, v := range lut {
		tokens = append(tokens, getTokens(s, k, v)...)
	}

	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i].start < tokens[j].start
	})

	firstDigit := tokens[0].value

	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i].end > tokens[j].end
	})

	secondDigit := tokens[0].value

	return firstDigit*10 + secondDigit
}

func part2(r io.Reader) int {
	s := bufio.NewScanner(r)

	acc := 0
	for s.Scan() {
		line := s.Text()
		num := getCalibrationNumberV2(line)
		fmt.Println(line, num)
		acc += num
	}
	return acc
}
