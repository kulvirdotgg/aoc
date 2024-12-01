package day1

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func Solution() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	now := time.Now()
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
	taken := time.Since(now).Microseconds()
	fmt.Println(taken)
}

// NOTE: some cool ass feature to read lines from file
// copied from my 1brc sol, but its cool so putting it here
func ReadChunk() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}
	defer file.Close()

	const chunkSize = 32

	buf := make([]byte, chunkSize)
	left := make([]byte, 0, chunkSize)
	for {
		nbytes, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nfinished reading exiting!!")
				return
			}
			fmt.Println("error reading file", err)
		}

		buf = buf[:nbytes]
		eol := bytes.LastIndex(buf, []byte{'\n'})

		// the actual data to be consumed
		// appened after left slice, because left contains newline read in previous iter
		// but that data belongs to this line
		data := append(left, buf[:eol]...)

		left = make([]byte, len(buf[eol+1:]))
		left = append(left, buf[eol+1:]...)

		fmt.Println(string(data))
	}
}
