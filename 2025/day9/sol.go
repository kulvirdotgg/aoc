package day9

import (
	"fmt"
	"math"

	"aoc/stl"
)

func Solution() {
	ip := stl.ReadFile("input.txt")
	// ip := stl.ReadFile("example.txt")

	// initializing the capacity of slide to size of inputs
	coords := make([][]int, 0, 500)
	for _, line := range ip {
		coords = append(coords, stl.IntsFromString(line))
	}

	area := largestRectangle(coords)
	fmt.Printf("Larges area formed by red tiles is: %d\n", area)
}

func largestRectangle(coords [][]int) (area int) {
	n := len(coords)
	for i := range n {
		for j := i + 1; j < n; j++ {
			first, second := coords[i], coords[j]

			l := math.Abs(float64(first[0] - second[0] + 1))
			b := math.Abs(float64(first[1] - second[1] + 1))

			area = max(area, int(l*b))
		}
	}
	return area
}
