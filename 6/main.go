package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPos(matrix [][]byte) (byte, int, int) {
	var direction byte

	direction = 'N'
	x := -1
	y := -1

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			switch matrix[i][j] {
			case '^':
				direction = '^'
				x = i
				y = j
			case 'v':
				direction = 'v'
				x = i
				y = j
			case '<':
				direction = '<'
				x = i
				y = j
			case '>':
				direction = '>'
				x = i
				y = j
			}
		}
	}

	return direction, x, y
}

func move(matrix [][]byte) bool {
	direction, x, y := getPos(matrix)
	switch direction {
	case '^':
		matrix[x][y] = 'X'
		if x-1 < 0 {
			return false
		}
		if (matrix[x-1][y]) != '#' {
			matrix[x-1][y] = '^'
		} else {
			matrix[x][y+1] = '>'
		}
	case 'v':
		matrix[x][y] = 'X'
		if x+1 >= len(matrix) {
			return false
		}
		if (matrix[x+1][y]) != '#' {
			matrix[x+1][y] = 'v'
		} else {
			matrix[x][y-1] = '<'
		}
	case '<':
		matrix[x][y] = 'X'
		if y-1 < 0 {
			return false
		}
		if (matrix[x][y-1]) != '#' {
			matrix[x][y-1] = '<'
		} else {
			matrix[x-1][y] = '^'
		}
	case '>':
		matrix[x][y] = 'X'
		if y+1 >= len(matrix) {
			return false
		}
		if (matrix[x][y+1]) != '#' {
			matrix[x][y+1] = '>'
		} else {
			matrix[x+1][y] = 'v'
		}
	}
	return true
}

func countDistinctPos(matrix [][]byte) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'X' {
				count++
			}
		}
	}
	return count
}

// func print(matrix [][]byte) {

// 	for i := 0; i < len(matrix); i++ {
// 		for j := 0; j < len(matrix[i]); j++ {
// 			fmt.Printf("%c ", matrix[i][j])
// 		}
// 		fmt.Println()
// 	}

// }

func part1(matrix [][]byte) int {
	is_inside := true
	// print(matrix)
	for is_inside {
		// fmt.Println()
		is_inside = move(matrix)
		// print(matrix)
	}
	return countDistinctPos(matrix)

}

func main() {
	input_file, err := os.ReadFile("input.txt")
	check(err)
	input := strings.Split(string(input_file), "\n")

	var matrix [][]byte
	for i := 0; i < len(input); i++ {
		matrix = append(matrix, []byte(input[i]))
	}

	fmt.Printf("Part1 Ans: %d\n", part1(matrix))
}
