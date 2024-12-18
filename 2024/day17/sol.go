package day17

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("extc.txt")

	for _, line := range ip {
		if len(line) == 0 {
			continue
		}
		fmt.Println(line)
	}
}
