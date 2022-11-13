package main

import (
	"bufio"
	"fmt"
	"os"
)

func getGcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m, result int
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &m)
		var numbers = make([]int, m)
		result = 0
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d ", &numbers[j])
		}

		for k := 0; k < m; k++ {
			for l := k + 1; l < m; l++ {
				result += getGcd(numbers[k], numbers[l])
			}
		}
		fmt.Fprintf(writer, "%d\n", result)
	}
}
