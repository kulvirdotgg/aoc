package day5

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	stl "aoc/stl"
)

// https://adventofcode.com/2025/day/5
func Solution() {
	// ip := stl.ReadFile("example.txt")
	ip := stl.ReadFile("input.txt")
	ranges, ids := splitRanges(ip)
	fresh := freshIds(ranges, ids)

	fmt.Printf("Fresh IDs count: %d\n", fresh)

	totalFresh := totalFreshIds(ranges)
	fmt.Printf("Total fresh IDs in DB: %d\n", totalFresh)
}

func freshIds(ranges [][]int64, ids []int64) (fresh int) {
	// for each food ID, check if it belongs to the fresh IDs
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				fresh++
				break
			}
		}
	}
	return fresh
}

// merge overlapping ranges
// for disjoint ranges calculate the range width
func totalFreshIds(ranges [][]int64) (totalFreshIds int64) {
	slices.SortFunc(ranges, func(a, b []int64) int {
		return cmp.Compare(a[0], b[0])
	})

	rangeStart, rangeEnd := ranges[0][0], ranges[0][1]
	for _, r := range ranges[1:] {
		s, e := r[0], r[1]

		// If the new range is disjoint range
		// [3, 5], [10, 14]
		// 5 < 10, hence [10, 14] is start of new range
		if rangeEnd < s {
			totalFreshIds += rangeEnd - rangeStart + 1
			rangeStart, rangeEnd = s, e
		} else {
			// Overlapping range, start point will remain same
			// end point will be max End of both ranges
			// [12, 18], [14, 16], 18 should be endpoint for this range
			rangeEnd = max(rangeEnd, e)
		}
	}
	totalFreshIds += rangeEnd - rangeStart + 1

	return totalFreshIds
}

// ranges: [[s,e], [s,e]], fresh ids: [1, 2]
func splitRanges(ip []string) (ranges [][]int64, ids []int64) {
	idx := 0
	for len(ip[idx]) != 0 {
		r := strings.Split(ip[idx], "-")
		s, _ := strconv.ParseInt(r[0], 10, 64)
		e, _ := strconv.ParseInt(r[1], 10, 64)

		ranges = append(ranges, []int64{s, e})
		idx++
	}

	// IDX is still at the empty seperator line, move idx forward
	idx++

	for idx < len(ip) {
		id, _ := strconv.ParseInt(ip[idx], 10, 64)
		ids = append(ids, id)

		idx++
	}

	return ranges, ids
}
