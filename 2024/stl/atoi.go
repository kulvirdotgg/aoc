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

		if (ch == ' ' || ch == ',' || ch == ';') && builder.Len() != 0 {
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
