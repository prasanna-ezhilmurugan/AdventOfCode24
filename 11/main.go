package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input_file, _ := os.ReadFile("sample_input.txt")
	input := strings.Split(string(input_file), " ")

	var arr []int
	for _, value := range input {
		element, _ := strconv.Atoi(value)
		arr = append(arr, element)
	}

	fmt.Printf("Part1 Ans: %d\n", part1(arr, 75))
}

func part1(arr []int, blinks int) int {

	// fmt.Println("Initial Arrangement")
	// fmt.Println(arr)
	// fmt.Println()

	for count := 0; count < blinks; count++ {
		size := len(arr)
		itr := 0
		for itr < size {
			if arr[itr] == 0 {
				arr[itr] = 1
			} else if findLength(arr[itr])%2 == 0 {
				operand := int(math.Pow(10, float64(findLength(arr[itr])/2)))
				arr = slices.Insert(arr, itr+1, arr[itr]%operand)
				arr[itr] /= operand
				itr++
				size++

			} else {
				arr[itr] *= 2024
			}
			itr++
		}

		// fmt.Printf("After %d blinks\n", count+1)
		// fmt.Println(arr)
		// fmt.Println()

	}

	return len(arr)
}


func findLength(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}
