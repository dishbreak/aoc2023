package almanac

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Almanac struct {
	Seeds      []int
	SeedRanges []Range
	tables     [][]rule
}

type rule struct {
	start, end, offset int
}

type Range struct {
	Start, End int
}

func FromString(s string) (a Almanac) {
	sections := strings.Split(s, "\n\n")

	seedStr := strings.Fields(strings.TrimPrefix(sections[0], "seeds: "))
	a.Seeds = make([]int, len(seedStr))

	a.SeedRanges = make([]Range, 0)

	for i := range seedStr {
		a.Seeds[i], _ = strconv.Atoi(seedStr[i])
	}

	for i := 0; i < len(a.Seeds); i += 2 {
		a.SeedRanges = append(a.SeedRanges, Range{a.Seeds[0], a.Seeds[0] + a.Seeds[1]})
	}

	sections = sections[1:]
	a.tables = make([][]rule, len(sections))

	for i, section := range sections {
		lines := strings.Split(section, "\n")
		lines = lines[1:]

		a.tables[i] = make([]rule, len(lines))
		for j, line := range lines {
			if line == "" {
				continue
			}
			pts := strings.Fields(line)
			if len(pts) == 0 {
				panic(fmt.Errorf("can't parse line '%s'", line))
			}

			r := rule{}

			r.start, _ = strconv.Atoi(pts[1])

			transform, _ := strconv.Atoi(pts[0])
			extent, _ := strconv.Atoi(pts[2])

			r.end = r.start + extent
			r.offset = transform - r.start

			a.tables[i][j] = r
		}

		sort.Slice(a.tables[i], func(j, k int) bool {
			return a.tables[i][j].start < a.tables[i][k].start
		})
	}

	return
}

func (a Almanac) String() string {
	var sb strings.Builder
	sb.WriteString("seeds: ")
	sb.WriteString(fmt.Sprint(a.Seeds))
	sb.WriteString("\n\n")

	sb.WriteString("rules:\n")
	for _, table := range a.tables {
		for _, r := range table {
			sb.WriteString(fmt.Sprintf("(%d, %d) -> %d\n", r.start, r.end, r.offset))
		}
		sb.WriteString("----\n")
	}
	return sb.String()
}

func (a Almanac) ToLocation(seed int) int {
	result := seed

	for _, table := range a.tables {
		result = remap(result, table)
	}

	return result
}

func remap(val int, rules []rule) int {
	// if it's lower than the lowest interval, spit it back
	if val < rules[0].start {
		return val
	}

	// if it's higher than highest interval end, spit it back
	if rules[len(rules)-1].end <= val {
		return val
	}

	for _, r := range rules {
		if r.start < val && r.end <= val {
			continue
		}
		if r.start <= val && r.end > val {
			return r.offset + val
		}
		if r.start > val {
			break
		}
	}
	return val
}
