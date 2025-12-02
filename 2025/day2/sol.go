package day2

import (
	"fmt"
	"iter"
	"strconv"
	"strings"

	stl "aoc/stl"
)

// https://adventofcode.com/2025/day/2
func Solution() {
	// ip := stl.ReadFile("example.txt")
	ip := stl.ReadFile("input.txt")

	// comma seperated ranges from the input line
	ranges := strings.SplitSeq(ip[0], ",")
	oneStarSum := identical(ranges)
	twoStarSum := identicalSplits(ranges)
	fmt.Println(oneStarSum, twoStarSum)
}

func identical(ranges iter.Seq[string]) (invalidSum int64) {
	for r := range ranges {
		// 11 - 22
		// start = 11 & end = 22
		startEnd := strings.Split(r, "-")
		s, e := startEnd[0], startEnd[1]

		start, end := getInt(s), getInt(e)
		for i := start; i <= end; i++ {
			numStr := strconv.FormatInt(i, 10)

			n := len(numStr)

			if n&1 == 1 {
				continue
			}

			firstHalf, secondHalf := numStr[:n/2], numStr[n/2:]
			if firstHalf == secondHalf {
				invalidSum += i
			}
		}
	}
	return invalidSum
}

func identicalSplits(ranges iter.Seq[string]) (invalidSum int64) {
	for r := range ranges {
		// 11 - 22
		// start = 11 & end = 22
		startEnd := strings.Split(r, "-")
		s, e := startEnd[0], startEnd[1]

		start, end := getInt(s), getInt(e)
		for i := start; i <= end; i++ {
			numStr := strconv.FormatInt(i, 10)
			n := len(numStr)

			for numOfSplits := 2; numOfSplits <= 10; numOfSplits++ {
				// not enough characters to split into `x` parts
				if n < numOfSplits {
					break
				}
				// if cannot split into equal `x` parts
				if n%numOfSplits != 0 {
					continue
				}

				splitSize := n / numOfSplits
				chunks := splitString(numStr, splitSize)

				// validate if all the chunks of `x` size are same
				init := chunks[0]
				sameChunks := true
				for _, c := range chunks {
					if init != c {
						sameChunks = false
						break
					}
				}

				// if all the chunks are same then its invalid ID, add it to sum
				// also stop checking further split sizes for same ID
				if sameChunks {
					invalidSum += i
					break
				}
			}
		}
	}
	return invalidSum
}

// split the string into equal sized chunks of `sz` size
func splitString(s string, sz int) (chunks []string) {
	for i := 0; i < len(s); i += sz {
		chunks = append(chunks, s[i:i+sz])
	}
	return chunks
}

func getInt(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return num
}
