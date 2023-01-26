package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	matrix = [][]int{}
)

func mulMatrix(matrix1, matrix2 [][]int) [][]int {
	res := [][]int{}
	res = append(res, []int{0, 0})
	res = append(res, []int{0, 0})
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				res[i][j] += matrix1[i][k] * matrix2[k][j] % 1000000007
			}
		}
	}
	return res
}

func power(a [][]int, b uint64) [][]int {
	if b == 1 {
		return a
	}
	tmp := power(a, b/2)
	if b%2 == 0 {
		return mulMatrix(tmp, tmp)
	} else {
		return mulMatrix(mulMatrix(tmp, tmp), a)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n uint64
	matrix = append(matrix, []int{1, 1})
	matrix = append(matrix, []int{1, 0})
	fmt.Fscanf(reader, "%d\n", &n)
	result := power(matrix, n)
	fmt.Fprintf(writer, "%d", result[0][1]%1000000007)
}
