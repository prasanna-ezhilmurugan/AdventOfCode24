package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input_file, err := os.ReadFile("input.txt")
	check(err)
	input_lines := strings.Split(string(input_file), "\n")
	var input [][]int
	for _, line := range input_lines {
		line := strings.Split(line, " ")
		var convertedLine []int
		for _, element := range line {
			converted_element, _ := strconv.Atoi(element)
			convertedLine = append(convertedLine, converted_element)
		}
		input = append(input, convertedLine)
	}

	fmt.Printf("Part1 Ans: %d\n", part1(input))
	fmt.Printf("Part2 Ans: %d\n", part2(input))
}

func part1(input [][]int) int {
	count := 0
	for _, list := range input {
		if isCorrectlyOrdered(list) {
			count++
		}
	}
	return count
}

func part2(input [][]int) int {
	count := 0
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			var temp []int
			temp = append(temp, line[:i]...)
			temp = append(temp, line[i+1:]...)
			if isCorrectlyOrdered(temp) {
				count++
				break
			}
		}
	}
	return count
}

func isCorrectlyOrdered(list []int) bool {

	var increasingList bool
	if list[0] > list[1] {
		increasingList = false
	} else {
		increasingList = true
	}

	for i := 0; i < len(list)-1; i++ {
		if increasingList {
			if !(list[i] < list[i+1] && list[i+1]-list[i] <= 3 && list[i+1]-list[i] >= 1) {
				return false
			}
		} else {
			if !(list[i] > list[i+1] && list[i]-list[i+1] <= 3 && list[i]-list[i+1] >= 1) {
				return false
			}
		}
	}
	return true
}
