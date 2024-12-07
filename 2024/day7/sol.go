package day7

import (
	"aoc/stl"
	"fmt"
	"strconv"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	callibratedVals := 0
	for _, line := range ip {
		callibratedVals += betterCallibration(stl.IntsFromString(line))
	}
	fmt.Println(callibratedVals)
}

func getCallibrated(line []int) int {
	total := line[0]
	ops := line[1:]
	var ok func(ops []int, res int, idx int) bool
	ok = func(ops []int, res, idx int) bool {
		if idx == len(ops) {
			return (res == total)
		}
		if res > total {
			return false
		}
		return ok(ops, res+ops[idx], idx+1) || ok(ops, res*ops[idx], idx+1)
	}

	// if we start with zero multiply with zero wouldn't yeild good results
	res := ops[0]
	if ok(ops, res, 1) {
		return total
	}
	return 0
}

func betterCallibration(line []int) int {
	total := line[0]
	ops := line[1:]
	var ok func(ops []int, res int, idx int) bool
	ok = func(ops []int, res, idx int) bool {
		if idx == len(ops) {
			return (res == total)
		}
		if res > total {
			return false
		}

		concatString := fmt.Sprintf("%d%d", res, ops[idx])
		concat, _ := strconv.Atoi(concatString)
		return ok(ops, res+ops[idx], idx+1) || ok(ops, res*ops[idx], idx+1) || ok(ops, concat, idx+1)
	}

	// if we start with zero multiply with zero wouldn't yeild good results
	res := ops[0]
	if ok(ops, res, 1) {
		return total
	}
	return 0
}
