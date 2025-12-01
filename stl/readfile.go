package stl

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	op := []string{}
	for scanner.Scan() {
		op = append(op, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return op
}
