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
	input := strings.Split(string(input_file), "\n")

	var rules [][2]int
	var testList [][]int
	isRule := true
	for _, line := range input {

		if line == "" {
			isRule = false
			continue
		}

		if isRule {
			rule := strings.Split(line, "|")

			var convRule [2]int
			convRule[0], _ = strconv.Atoi(rule[0])
			convRule[1], _ = strconv.Atoi(rule[1])

			rules = append(rules, convRule)
		} else {

			list := strings.Split(line, ",")
			var convlist []int
			for _, element := range list {
				element, _ := strconv.Atoi(element)
				convlist = append(convlist, element)
			}
			testList = append(testList, convlist)

		}
	}
	// fmt.Println(testList)
	hm := make(map[int][]int)

	//make a hash map with the rules
	for _, rule := range rules {
		hm[rule[0]] = append(hm[rule[0]], rule[1])
	}

	fmt.Printf("Part1 Ans: %d\n", part1(hm, testList))
	fmt.Printf("Part2 Ans: %d\n", part2(hm, testList))
}

func part1(hm map[int][]int, testList [][]int) int {
	result := 0
	for _, list := range testList {
		if isCorrectlyOrdered(hm, list) {
			result += list[len(list)/2]
		}
	}
	return result
}

func part2(hm map[int][]int, testList [][]int) int {
	result := 0
	for _, list := range testList {
		if !isCorrectlyOrdered(hm, list) {
			list = correctOrder(hm, list)
			result += list[len(list)/2]
		}
	}
	return result
}

func isCorrectlyOrdered(hm map[int][]int, list []int) bool {
	isCorrectlyOrdered := true
	for i := 0; i < len(list)-1; i++ {
		// fmt.Println(hm[list[i]])
		// fmt.Println(list[i+1:])

		//check whether list[i+1:] is subset of hm[list[i]]
		if !isSubset(list[i+1:], hm[list[i]]) {
			isCorrectlyOrdered = false
			break
		}
	}

	return isCorrectlyOrdered
}

func correctOrder(hm map[int][]int, list []int) []int{
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-i-1; j++ {
			// check list[j] is in the hm[list[j+1]]
			// if yes swap list[j] with list[j+1]
			if searchElement(hm[list[j+1]], list[j]){
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}
	return list
}

func isSubset(subset []int, superset []int) bool {
	checkset := make(map[int]bool)
	for _, element := range subset {
		checkset[element] = true
	}
	for _, value := range superset {
		if checkset[value] {
			delete(checkset, value)
		}
	}

	return len(checkset) == 0
}

func searchElement(list []int, key int) bool {
	for _, element := range list{
		if element == key{
			return true
		}
	}
	return false
}
