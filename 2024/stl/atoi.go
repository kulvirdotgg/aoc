package stl

import (
	"strconv"
	"strings"
	"unicode"
)

// NOTE: Cool ass way to get integers from a string
// obv there was other ways too, but this is cool
func IntsFromString(line string) []int {
	array := []int{}

	neg := false
	var builder strings.Builder
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			builder.WriteRune(ch)
		}

		if ch == '-' {
			neg = true
		}

		if (ch == ' ' || ch == ',' || ch == ';' || ch == '|') && builder.Len() != 0 {
			num, err := strconv.ParseInt(builder.String(), 10, 64)
			if err != nil {
				panic(err)
			}

			if neg {
				num *= -1
			}
			builder.Reset()
			neg = false

			array = append(array, int(num))
		}
	}

	if builder.Len() != 0 {
		num, err := strconv.ParseInt(builder.String(), 10, 64)
		if err != nil {
			panic(err)
		}

		if neg {
			num *= -1
		}
		builder.Reset()
		neg = false

		array = append(array, int(num))
	}

	return array
}

func MatrixFromString(ip []string) [][]int64 {
	matrix := [][]int64{}

	var builder strings.Builder
	for _, line := range ip {
		arr := []int64{}
		for _, data := range line {
			builder.WriteRune(data)
			num, _ := strconv.ParseInt(builder.String(), 10, 64)
			arr = append(arr, num)
			builder.Reset()
		}
		matrix = append(matrix, arr)
	}
	return matrix
}
