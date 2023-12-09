package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
)

func main() {
	f, err := os.Open("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))

}

func part1(r io.Reader) (steps int) {
	directions, nodes := parse(r)

	cond := nodes["AAA"]

	for ; cond.name != "ZZZ"; steps++ {
		switch directions[steps%len(directions)] {
		case 'L':
			cond = cond.left
		case 'R':
			cond = cond.right
		}
	}

	return
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part2(r io.Reader) int {
	directions, nodes := parse(r)

	var conds []*node
	for name, node := range nodes {
		if strings.HasSuffix(name, "A") {
			conds = append(conds, node)
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(conds))

	results := make(chan int)
	go func() {
		wg.Wait()
		close(results)
	}()

	travel := func(n *node) {
		defer wg.Done()
		cond := n
		seen := make(map[string]bool)

		steps := 0
		for ; ; steps++ {
			if seen[cond.name] && strings.HasSuffix(cond.name, "Z") {
				break
			}
			seen[cond.name] = true
			switch directions[steps%len(directions)] {
			case 'L':
				cond = cond.left
			case 'R':
				cond = cond.right
			}
		}
		results <- steps
	}

	for _, cond := range conds {
		go travel(cond)
	}

	var lengths []int
	for val := range results {
		lengths = append(lengths, val)
	}

	return LCM(lengths[0], lengths[1], lengths[2:]...)
}

type node struct {
	name        string
	left, right *node
}

var matchNode = regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)$`)

func parse(r io.Reader) (directions []byte, nodes map[string]*node) {
	s := bufio.NewScanner(r)

	s.Scan()
	b := s.Bytes()

	directions = append(directions, b...)

	nodes = make(map[string]*node)

	getOrCreate := func(name string) *node {
		n, ok := nodes[name]
		if !ok {
			n = &node{name: name}
			nodes[name] = n
		}
		return n
	}

	for s.Scan() {
		line := s.Text()
		matches := matchNode.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		root := getOrCreate(matches[1])
		left := getOrCreate(matches[2])
		right := getOrCreate(matches[3])

		root.left, root.right = left, right
	}

	return
}
