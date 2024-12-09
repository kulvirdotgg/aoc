package day9

import (
	"aoc/stl"
	"fmt"
	"strconv"
	"strings"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	dkmap := diskMap(ip[0])
	disk := generateDisk(dkmap)
	checksum := updateChecksum(disk)
	fmt.Println(checksum)
}

func updateChecksum(disk []int) int {
	size := len(disk)

	writeIdx := size - 1
	for disk[writeIdx] == -1 {
		writeIdx--
	}

	checksum := 0
	i := 0
	for i <= writeIdx {
		if disk[i] == -1 {
			checksum += (i * disk[writeIdx])
			writeIdx--
			for disk[writeIdx] == -1 {
				writeIdx--
			}
		} else {
			checksum += (i * disk[i])
		}
		i++
	}
	return checksum
}

func generateDisk(diskmap []int64) []int {
	disk := []int{}
	isFile := true
	id := 0
	for _, size := range diskmap {
		if isFile {
			for range size {
				disk = append(disk, id)
			}
			id++
		} else {
			for range size {
				disk = append(disk, -1)
			}
		}
		isFile = !isFile
	}
	return disk
}

func diskMap(ip string) []int64 {
	dkmap := []int64{}

	var builder strings.Builder
	for _, data := range ip {
		builder.WriteRune(data)
		num, _ := strconv.ParseInt(builder.String(), 10, 64)
		dkmap = append(dkmap, num)
		builder.Reset()
	}
	return dkmap
}
