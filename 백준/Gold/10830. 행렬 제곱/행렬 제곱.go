package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	n, m   int
)

func multiple(mat1, mat2 [][]int) [][]int {
	var res = make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for z := 0; z < n; z++ {
				res[i][j] += mat1[i][z] * mat2[z][j] % 1000
			}
		}
	}
	return res
}
func power(a [][]int, b int) [][]int {
	if b == 1 {
		return a
	} else {
		div := power(a, b/2)
		if b%2 == 0 {
			return multiple(div, div)
		} else {
			return multiple(multiple(div, div), a)
		}
	}
}
func main() {
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &matrix[i][j])
		}
	}
	res := power(matrix, m)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprintf(writer, "%d ", res[i][j]%1000)
		}
		fmt.Fprint(writer, "\n")
	}
}
