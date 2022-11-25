package main

import (
	"bufio"
	"fmt"
	"os"
)

var target = make([][]int, 9)
var zeros = [][]int{}
var numbers = make([]int, 9)
var writer = bufio.NewWriter(os.Stdout)

func rowSearch(row, targetNum int) bool {
	for c := 0; c < 9; c++ {
		if target[row][c] == targetNum {
			return false
		}
	}
	return true
}

func colSearch(col, targetNum int) bool {
	for r := 0; r < 9; r++ {
		if target[r][col] == targetNum {
			return false
		}
	}
	return true
}

func squareSearch(row, col, targetNum int) bool {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if target[row/3*3+i][col/3*3+j] == targetNum {
				return false
			}
		}
	}
	return true
}

func dfs(depth int) {
	if depth == len(zeros) {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Fprintf(writer, "%d ", target[i][j])
			}
			fmt.Fprintf(writer, "\n")
		}
		writer.Flush()
		os.Exit(0)
	}
	v := zeros[depth]
	for num := 1; num <= 9; num++ {
		if rowSearch(v[0], num) && colSearch(v[1], num) && squareSearch(v[0], v[1], num) {
			target[v[0]][v[1]] = num
			dfs(depth + 1)
			target[v[0]][v[1]] = 0

		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	target = make([][]int, 9)
	for i := 0; i < 9; i++ {
		target[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			fmt.Fscanf(reader, "%d ", &target[i][j])
			if target[i][j] == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
		numbers[i] = (i + 1)
	}
	dfs(0)
}
