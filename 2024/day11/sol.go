package day11

import (
	"aoc/stl"
	"fmt"
	"strconv"
)

func Solution() {
	ip := stl.ReadFile("input.txt")
	arrangement := stl.IntsFromString(ip[0])
	small := lessBlinks(arrangement, 25)
	large := moreBlinks(arrangement, 75)
	fmt.Println(small, large)
}

func lessBlinks(ip []int, blinks int) int {
	arrange := ip

	for range blinks {
		temp := []int{}
		for _, stone := range arrange {
			switch {
			case stone == 0:
				temp = append(temp, 1)
			case len(strconv.Itoa(stone))%2 == 0:
				f, s := splitStone(stone)
				temp = append(temp, f)
				temp = append(temp, s)
			default:
				temp = append(temp, 2024*stone)
			}
		}
		arrange = temp
	}
	return len(arrange)
}

func moreBlinks(ip []int, blinks int) uint64 {
	dp := make(map[int]uint64)
	for _, stone := range ip {
		dp[stone] += 1
	}

	for range blinks {
		updatedDp := dp
		for stone := range dp {
			cnt := dp[stone]
			switch {
			case stone == 0:
				updatedDp[1] += cnt
			case len(strconv.Itoa(stone))%2 == 0:
				f, s := splitStone(stone)
				updatedDp[f] += cnt
				updatedDp[s] += cnt
			default:
				newStone := 2024 * stone
				updatedDp[newStone] += cnt
			}
		}
		dp = updatedDp
	}
	var stones uint64 = 0
	for _, count := range dp {
		stones += count
	}
	return stones
}

func splitStone(stone int) (int, int) {
	str := strconv.Itoa(stone)
	first, _ := strconv.Atoi(str[:len(str)/2])
	second, _ := strconv.Atoi(str[len(str)/2:])
	return first, second
}
