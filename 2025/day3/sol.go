package day3

import (
	"fmt"

	stl "aoc/stl"
)

// https://adventofcode.com/2025/day/3
func Solution() {
	// ip := stl.ReadFile("example.txt")
	ip := stl.ReadFile("input.txt")
	joltage := getJoltage(ip)
	largeJoltage := getLargeJoltage(ip)

	fmt.Println(joltage, largeJoltage)
}

func getJoltage(powerBanks []string) (joltage int) {
	n := len(powerBanks[0])

	for _, batteries := range powerBanks {
		maxJoltage := 0

		for i := range batteries {
			for j := i + 1; j < n; j++ {
				tens := batteries[i] - '0'
				ones := batteries[j] - '0'

				joltage := tens*10 + ones
				maxJoltage = max(maxJoltage, int(joltage))
			}
		}
		joltage += maxJoltage
	}
	return joltage
}

func getLargeJoltage(banks []string) (joltage int) {
	n := len(banks[0])

	for _, bank := range banks {
		stack := []int{}
		for i := range n {
			currDigit := int(bank[i] - '0')

			// standard monotonic stack question
			// sucks to be me, it took soo much time to solve this
			for len(stack) > 0 && stack[len(stack)-1] < currDigit && len(stack)+(n-i) > 12 {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, currDigit)
		}

		if len(stack) > 12 {
			stack = stack[:12]
		}

		joltage += numFromArr(stack)
	}
	return joltage
}

func numFromArr(s []int) (num int) {
	n := len(s)

	multiplier := 1
	for i := n - 1; i >= 0; i-- {
		num += (s[i] * multiplier)
		multiplier *= 10
	}
	return num
}
