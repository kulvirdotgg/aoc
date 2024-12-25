package day23

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	locks, keys := keyLocks(ip)
	// fmt.Println(locks)
	// fmt.Println(keys)

	pairs := getPairs(locks, keys)
	fmt.Println(pairs)
}

func getPairs(locks, keys [][5]int) int {
	pairs := 0
	for _, lock := range locks {
		for _, key := range keys {
			overlap := false
			for i := 0; i < 5; i++ {
				if lock[i]+(5-key[i]) > 5 {
					// fmt.Printf("overlap in %d, lock value: %d, key val: %d\n", i, lock[i], 5-key[i])
					// fmt.Printf("overlap lock: %v key: %v\n", lock, key)
					overlap = true
					break
				}
			}
			if !overlap {
				// fmt.Printf("lock: %v\nkey: %v\n", lock, key)
				pairs++
			}
		}
	}
	return pairs
}

func keyLocks(ip []string) ([][5]int, [][5]int) {
	locks := [][5]int{}
	keys := [][5]int{}

	kl := []string{}
	for _, line := range ip {
		if len(line) == 0 {
			if kl[0][0] == '.' {
				key := getHeights(kl)
				keys = append(keys, key)
			} else {
				lock := getHeights(kl)
				locks = append(locks, lock)
			}
			kl = []string{}
		} else {
			kl = append(kl, line)
		}
	}
	if kl[0][0] == '.' {
		key := getHeights(kl)
		keys = append(keys, key)
	} else {
		lock := getHeights(kl)
		locks = append(locks, lock)
	}
	return locks, keys
}

func getHeights(kl []string) [5]int {
	heights := [5]int{}

	for c := 0; c < 5; c++ {
		r := 1
		for r <= 5 {
			if kl[r][c] != kl[r-1][c] {
				break
			}
			r++
		}
		heights[c] = r - 1
	}
	return heights
}
