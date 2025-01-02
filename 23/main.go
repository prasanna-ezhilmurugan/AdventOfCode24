package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input_file, _ := os.ReadFile("input.txt")
	input := strings.Split(string(input_file), "\n")
	hm := make(map[string][]string)

	for _, line := range input {
		arg1 := strings.Split(line, "-")[0]
		arg2 := strings.Split(line, "-")[1]
		hm[arg1] = append(hm[arg1], arg2)
		hm[arg2] = append(hm[arg2], arg1)
	}

	// fmt.Println(hm)
	fmt.Printf("Part1 Ans: %d", part1(hm))
}

func part1(hm map[string][]string) int {
	var uniqueSets [][]string
	for i := range hm {
		for j := range hm {
			for k := range hm {
				if isMutuallyConnected(hm[i], j, k) && isMutuallyConnected(hm[j], i, k) && isMutuallyConnected(hm[k], i, j) && isUnique(uniqueSets, i, j, k) {
					uniqueSets = append(uniqueSets, []string{i, j, k})
				}
			}
		}
	}
	// fmt.Println(uniqueSets)
	count := 0
	for _, line := range uniqueSets {
		for _, comp := range line {
			if strings.HasPrefix(comp, "t") {
				count++
				break
			}
		}
	}
	return count
}

func isUnique(uniqueSets [][]string, i string, j string, k string) bool {
	for _, arr := range uniqueSets {
		if slices.Contains(arr, i) && slices.Contains(arr, j) && slices.Contains(arr, k) {
			return false
		}
	}
	return true
}

func isMutuallyConnected(list []string, x string, y string) bool {
	return slices.Contains(list, x) && slices.Contains(list, y)
}
