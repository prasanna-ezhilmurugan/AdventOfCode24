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
	// fmt.Println("Prasanna Ezhilmurugan")
	input_file, err := os.ReadFile("sample_input.txt")
	check(err)
	input := strings.Split(string(input_file), "\n")
	// for _, line := range input {
	// 	fmt.Println(line)
	// }
	hm := make(map[int][]int)
	for _, line := range input {
		result, _ := strconv.Atoi(strings.Split(line, ": ")[0])

		var operands []int
		for _, str := range strings.Split(strings.Split(line, ": ")[1], " ") {
			operand, _ := strconv.Atoi(str)
			operands = append(operands, operand)
		}

		hm[result] = operands
	}

	totalCalibrationResult := 0
	for result, operands := range hm {
		if isCalibratable(result, operands[1:], operands[0]) {
			totalCalibrationResult += result
			// fmt.Println(totalCalibrationResult)
		}
		fmt.Println()
	}
	fmt.Printf("Part1 Ans: %d\n", totalCalibrationResult)
}

func isCalibratable(result int, operands []int, current int) bool {
	if len(operands) == 0 {
		if result != current {
			return false
		} else {
			return true
		}
	}
	fmt.Printf("%d | %d |\n\t\t\t%d (x|+) %d = %d\n", result, operands, current, operands[0], current*operands[0])
	return isCalibratable(result, operands[1:], current*operands[0]) || isCalibratable(result, operands[1:], current+operands[0])
}
