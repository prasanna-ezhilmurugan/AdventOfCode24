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

func part1(matrix [][]byte) int {
	count := 0
	directions := [8][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] == 'X' {

				for k := 0; k < 8; k++ {
					x := directions[k][0]
					y := directions[k][1]

					if i+x >= 0 && j+y >= 0 && i+x < len(matrix[i]) && j+y < len(matrix[i]) && matrix[i+x][j+y] == 'M' {
						if i+x*2 >= 0 && j+y*2 >= 0 && i+x*3 >= 0 && j+y*3 >= 0 && i+x*3 < len(matrix[i]) && j+y*3 < len(matrix[i]) &&
							matrix[i+x*2][j+y*2] == 'A' &&
							matrix[i+x*3][j+y*3] == 'S' {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

func part2(matrix [][]byte) int {
	count := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] == 'A' {
				if i-1 >= 0 && j-1 >= 0 && j+1 < len(matrix) && i+1 < len(matrix) &&
					((matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S') || (matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M')) &&
					((matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] == 'S') || (matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] == 'M')) {
					count++
				}
			}
		}
	}
	return count
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
	fmt.Printf("Part2 Ans: %d\n", part2(matrix))

}
