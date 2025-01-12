package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parse(string_values []string) []int {
	var values []int
	for _, string_value := range string_values {
		value, _ := strconv.Atoi(string_value)
		values = append(values, value)
	}
	return values
}

func main() {
	input_file, _ := os.ReadFile("input.txt")
	input := strings.Split(string(input_file), "\n")
	fmt.Printf("Part1 Ans: %d\n", part1(input))
	fmt.Printf("Part2 Ans: %d\n", part2(input))
}

func part1(input []string) int {
	result := 0
	for _, line := range input {
		target, _ := strconv.Atoi(strings.Split(line, ": ")[0])
		values := parse(strings.Split(strings.Split(line, ": ")[1], " "))

		if is_calibratable(target, values[1:], values[0]) {
			result += target
		}
	}
	return result
}

func is_calibratable(target int, values []int, current int) bool {
	if len(values) == 1 {
		if values[0]*current == target || values[0]+current == target {
			return true
		} else {
			return false
		}
	}
	return is_calibratable(target, values[1:], current*values[0]) || is_calibratable(target, values[1:], current+values[0])
}

func part2(input []string) int {
	result := 0
	for _, line := range input {
		target, _ := strconv.Atoi(strings.Split(line, ": ")[0])
		values := parse(strings.Split(strings.Split(line, ": ")[1], " "))

		if is_concat_calibratable(target, values[1:], values[0]) {
			result += target
		}
	}
	return result
}

func concat(operand1 int, operand2 int) int {
	return operand1*int(math.Pow(10, float64(len(strconv.Itoa(operand2))))) + operand2
}
func is_concat_calibratable(target int, values []int, current int) bool {

	if len(values) == 1 {
		if values[0]*current == target || values[0]+current == target || concat(current, values[0]) == target {
			return true
		} else {
			return false
		}
	}
	return is_concat_calibratable(target, values[1:], concat(current, values[0])) || is_concat_calibratable(target, values[1:], current*values[0]) || is_concat_calibratable(target, values[1:], current+values[0])
}
