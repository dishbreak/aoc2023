package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sync"

	"github.com/dishbreak/aoc2023/cmd/day5/almanac"
)

func main() {
	f, err := os.Open("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	contents, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	input := string(contents)

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input string) int {
	a := almanac.FromString(input)

	minLoc := math.MaxInt
	for _, seed := range a.Seeds {
		loc := a.ToLocation(seed)
		if loc < minLoc {
			minLoc = loc
		}
	}

	return minLoc
}

func part2(input string) int {
	a := almanac.FromString(input)

	result := make(chan int)

	var wg sync.WaitGroup
	for _, seedRange := range a.SeedRanges {
		wg.Add(seedRange.End - seedRange.Start)
	}

	for _, seedRange := range a.SeedRanges {
		for s := seedRange.Start; s < seedRange.End; s++ {
			go func(seed int) {
				defer wg.Done()
				result <- a.ToLocation(seed)
			}(s)
		}
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	minLoc := math.MaxInt
	for loc := range result {
		if loc < minLoc {
			minLoc = loc
		}
	}

	return minLoc
}
