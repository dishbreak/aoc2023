package main

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"io"
)

var (
	ptNorth = image.Pt(0, -1)
	ptSouth = image.Pt(0, 1)
	ptWest  = image.Pt(-1, 0)
	ptEast  = image.Pt(1, 0)
)

type pipeTile struct {
	next map[image.Point]image.Point
}

var (
	// | is a vertical pipe connecting north and south.
	pipeVertical = &pipeTile{
		next: map[image.Point]image.Point{
			ptNorth: ptNorth,
			ptSouth: ptSouth,
		},
	}
	// - is a horizontal pipe connecting east and west.
	pipeHorizontal = &pipeTile{
		next: map[image.Point]image.Point{
			ptEast: ptEast,
			ptWest: ptWest,
		},
	}
	// L is a 90-degree bend connecting north and east.
	pipeNEEll = &pipeTile{
		next: map[image.Point]image.Point{
			ptWest:  ptNorth,
			ptSouth: ptEast,
		},
	}
	// J is a 90-degree bend connecting north and west.
	pipeNWEll = &pipeTile{
		next: map[image.Point]image.Point{
			ptEast:  ptNorth,
			ptSouth: ptWest,
		},
	}
	// 7 is a 90-degree bend connecting south and west.
	pipeSWEll = &pipeTile{
		next: map[image.Point]image.Point{
			ptNorth: ptWest,
			ptEast:  ptSouth,
		},
	}
	// F is a 90-degree bend connecting south and east.
	pipeSEEll = &pipeTile{
		next: map[image.Point]image.Point{
			ptNorth: ptEast,
			ptWest:  ptSouth,
		},
	}
)

var pipeLookup = map[byte]*pipeTile{
	// | is a vertical pipe connecting north and south.
	'|': pipeVertical,
	// - is a horizontal pipe connecting east and west.
	'-': pipeHorizontal,
	// L is a 90-degree bend connecting north and east.
	'L': pipeNEEll,
	// J is a 90-degree bend connecting north and west.
	'J': pipeNWEll,
	// 7 is a 90-degree bend connecting south and west.
	'7': pipeSWEll,
	// F is a 90-degree bend connecting south and east.
	'F': pipeSEEll,
	// . is ground; there is no pipe in this tile.
	'.': nil,
}

func parse(r io.Reader) (map[image.Point]*pipeTile, image.Point) {

	space := make(map[image.Point]*pipeTile)
	s := bufio.NewScanner(r)

	var start image.Point
	for y := 0; s.Scan(); y++ {
		for x, c := range s.Text() {
			if c == 'S' {
				start = image.Pt(x, y)
				continue
			}
			tile, ok := pipeLookup[byte(c)]
			if !ok {
				panic(fmt.Errorf("unrecognized token '%s'", string(c)))
			}
			space[image.Pt(x, y)] = tile
		}
	}

	return space, start
}

type path struct {
	dir, pos image.Point
}

func longestPt(space map[image.Point]*pipeTile, ptStart image.Point) int {
	steps := 1 // A journey of a thousand miles begins with a single step

	var legs []path

	for _, dir := range []image.Point{ptNorth, ptSouth, ptEast, ptWest} {
		n := ptStart.Add(dir)

		// if the space doesn't exist or h
		tile, ok := space[n]
		if !ok || tile == nil {
			continue
		}

		// if we can't find a next direction, this pipe doesn't connect
		if _, ok := tile.next[dir]; !ok {
			continue
		}

		legs = append(legs, path{dir, ptStart})
	}

	if len(legs) != 2 {
		panic(fmt.Errorf("expected 2 starting points, got %d", len(legs)))
	}

	advance := func() {
		for i := range legs {
			legs[i].pos = legs[i].pos.Add(legs[i].dir)
			tile, ok := space[legs[i].pos]
			if !ok && tile == nil {
				panic(errors.New("pipe dream, found a hole"))
			}
			legs[i].dir, ok = tile.next[legs[i].dir]
			if !ok {
				panic(errors.New("dead end"))
			}
		}
	}
	advance()

	for ; legs[0].pos != legs[1].pos; steps++ {
		advance()
	}

	return steps
}
