package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	Start, End int
}

func parseInput(filename string) ([]Interval, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()

	var ranges []Interval
	var ids []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "-") { // first part
			parts := strings.Split(scanner.Text(), "-")
			var intParts [2]int
			intParts[0], _ = strconv.Atoi(parts[0])
			intParts[1], _ = strconv.Atoi(parts[1])
			ranges = append(ranges, Interval{intParts[0], intParts[1]})
		} else if scanner.Text() != "" { // second part
			id, _ := strconv.Atoi(scanner.Text())
			ids = append(ids, id)
		}
	}
	return ranges, ids
}

func checkId(id int, merged []Interval) bool {
	i := sort.Search(len(merged), func(i int) bool {
		return merged[i].Start > id
	})

	if i > 0 {
		candidate := merged[i-1]
		// fmt.Println("id", id, "is checked in this range:", candidate)
		if id <= candidate.End && id >= candidate.Start {
			fmt.Println("id", id, "is valid in range", candidate)
			return true
		}
	}
	return false
}

func mergeRanges(ranges []Interval) []Interval {
	merged := []Interval{ranges[0]}
	for i, r := range ranges {
		if i == 0 {
			continue
		}
		if r.Start < ranges[i-1].End {
			merged[len(merged)-1].End = max(r.End, ranges[i-1].End)
			continue
		}
		merged = append(merged, r)
	}
	return merged
}

func main() {
	ranges, ids := parseInput("../input/day05.txt")
	// ranges, ids := parseInput("../input/test05.txt")

	fmt.Println("how they are", ranges)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
	fmt.Println("sorted", ranges)
	merged := mergeRanges(ranges)
	fmt.Println("merged", merged)
	res := 0
	for _, id := range ids {
		if checkId(id, merged) {
			res++
		}
	}
	fmt.Println(res)
}
