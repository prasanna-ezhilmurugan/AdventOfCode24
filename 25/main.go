package main

import (
	"fmt"
	"os"
	"strings"
)

const MAXHEIGHT = 6

func getMatrix(input string) [][]byte {
	var matrix [][]byte
	for _, line := range strings.Split(input, "\n") {
		matrix = append(matrix, []byte(line))
	}
	return matrix
}

func isLock(matrix [][]byte) bool {
	length := len(matrix[0])
	topRow := string(matrix[0])
	var checkRow string
	for i := 0; i < length; i++ {
		checkRow += "#"
	}
	return checkRow == topRow
}

func getLockHeight(schematic [][]byte) []int {
	result := make([]int, len(schematic[0]))
	for i := 1; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			if schematic[i][j] == '#' {
				result[j]++
			}
		}
	}
	return result
}

func getKeyHeight(schematic [][]byte) []int {
	result := make([]int, len(schematic[0]))
	for i := 0; i < len(schematic)-1; i++ {
		for j := 0; j < len(schematic[i]); j++ {
			if schematic[i][j] == '#' {
				result[j]++
			}
		}
	}
	return result
}

func isOverlapping(lock []int, key []int) bool {
	for i := range lock {
		if lock[i]+key[i] >= MAXHEIGHT {
			return true
		}
	}
	return false
}

func main() {
	input_file, _ := os.ReadFile("input.txt")
	input := strings.Split(string(input_file), "\n\n")

	var keys [][][]byte
	var locks [][][]byte
	for _, schematic := range input {
		schematic := getMatrix(schematic)
		if isLock(schematic) {
			locks = append(locks, schematic)
		} else {
			keys = append(keys, schematic)
		}
	}

	var keyHeights [][]int
	var lockHeights [][]int

	for _, lock := range locks {
		lockHeights = append(lockHeights, getLockHeight(lock))
	}
	for _, key := range keys {
		keyHeights = append(keyHeights, getKeyHeight(key))
	}
	fmt.Println(lockHeights)
	fmt.Println(keyHeights)

	result := 0
	for _, key := range keyHeights {
		for _, lock := range lockHeights {
			if !isOverlapping(lock, key) {
				result++
			}
		}
	}
	fmt.Println(result)
}
