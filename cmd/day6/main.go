package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	raceRecords := parseRaceRecords(string(b))
	fmt.Printf("Part 1: %d\n", part1(raceRecords))
	fmt.Printf("Part 2: %d\n", part2(string(b)))
}

type raceRecord struct {
	time, distance int
}

func parseRaceRecords(s string) []raceRecord {
	result := make([]raceRecord, 0)

	lines := strings.Split(s, "\n")
	times := strings.Fields(lines[0])[1:]
	dists := strings.Fields(lines[1])[1:]

	for i := range times {
		r := raceRecord{}
		r.time, _ = strconv.Atoi(times[i])
		r.distance, _ = strconv.Atoi(dists[i])
		result = append(result, r)
	}

	return result
}

func waysToWin(race raceRecord) int {
	acc := 0
	for i := 1; i < race.time; i++ {
		if i*(race.time-i) <= race.distance {
			continue
		}
		acc++
	}
	return acc
}

func part1(records []raceRecord) int {
	acc := 1
	for _, race := range records {
		acc *= waysToWin(race)
	}
	return acc
}

func part2(s string) int {
	lines := strings.Split(s, "\n")

	timeStr := strings.Replace(lines[0], " ", "", -1)
	timeStr = strings.TrimPrefix(timeStr, "Time:")

	distanceStr := strings.Replace(lines[1], " ", "", -1)
	distanceStr = strings.TrimPrefix(distanceStr, "Distance:")

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		panic(err)
	}
	dist, err := strconv.Atoi(distanceStr)
	if err != nil {
		panic(err)
	}

	minHold := 1
	for ; minHold*(time-minHold) <= dist; minHold++ {
	}

	maxHold := time - 1
	for ; maxHold*(time-maxHold) <= dist; maxHold-- {
	}

	return maxHold - minHold + 1

}
