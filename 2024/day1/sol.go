package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := []int{}
	rightFreq := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()

		list := strings.Split(line, " ")
		one, _ := strconv.Atoi(list[0])
		two, _ := strconv.Atoi(list[3])
		left = append(left, one)
		rightFreq[two]++
	}

	sop := 0
	for i := 0; i < len(left); i++ {
		if count, ok := rightFreq[left[i]]; ok {
			sop += left[i] * count
		}
	}
	fmt.Println(sop)
}
