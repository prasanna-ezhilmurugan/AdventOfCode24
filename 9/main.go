package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func rearrageDisk(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == -1 {
			src := i
			for j := len(arr) - 1; j > 0 && j > i; j-- {
				if arr[j] != -1 {
					dest := j
					temp := arr[src]
					arr[src] = arr[dest]
					arr[dest] = temp
					break
				}
			}
		}
	}
	return arr
}

func createDisk(arr []int) []int {
	var disk []int
	count := 0
	for k, element := range arr {
		for i := 0; i < element; i++ {
			if k%2 == 0 {
				disk = append(disk, count)
			} else {
				disk = append(disk, -1)
			}
		}
		if k%2 == 0 {
			count++
		}
	}
	return disk
}

func part1(arr []int) int {
	result := 0
	for i, element := range arr {
		if element != -1 {
			result += element * i
		} else {
			break
		}
	}
	return result
}

func main() {
	input_file, err := os.ReadFile("input.txt")
	check(err)

	input := string(input_file)
	arr := make([]int, len(input))
	for i, r := range input {
		arr[i], _ = strconv.Atoi(string(r))
	}
	// fmt.Println(createDisk(arr))
	// fmt.Println(rearrageDisk(createDisk(arr)))
	fmt.Println(part1(rearrageDisk(createDisk(arr))))
}
