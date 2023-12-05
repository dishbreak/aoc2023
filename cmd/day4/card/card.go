package card

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	Id      int
	winner  [100]bool
	numbers []int
}

var parser = regexp.MustCompile(`^Card \s*(\d+): ([\d ]+)\|([\d ]+)$`)

func FromString(s string) (c Card) {
	matches := parser.FindStringSubmatch(s)
	if len(matches) != 4 {
		panic(fmt.Errorf("input string '%s' is invalid", s))
	}

	c.Id, _ = strconv.Atoi(matches[1])

	winners := strings.Fields(matches[2])
	for _, winner := range winners {
		num, _ := strconv.Atoi(winner)
		c.winner[num] = true
	}

	numbers := strings.Fields(matches[3])
	c.numbers = make([]int, len(numbers))
	for i, number := range numbers {
		c.numbers[i], _ = strconv.Atoi(number)
	}

	return
}

func (c Card) Score() (score int) {
	for _, number := range c.numbers {
		if !c.winner[number] {
			continue
		}
		if score == 0 {
			score = 1
			continue
		}
		score = score << 1
	}

	return
}

func (c Card) Matches() (matches int) {
	for _, number := range c.numbers {
		if !c.winner[number] {
			continue
		}
		matches++
	}
	return
}
