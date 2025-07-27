package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func find(n int, stock []int) int {
	for i, j := 0, len(stock)-1; i < j; i, j = i+1, j-1 {
		stock[i], stock[j] = stock[j], stock[i]
	}
	max := stock[0]
	sum := 0
	for i := 1; i < n; i++ {
		if max < stock[i] {
			max = stock[i]
		} else {
			sum += max - stock[i]
		}
	}
	return sum
}
func main() {
	defer writer.Flush()
	var T, n int
	fmt.Fscanf(reader, "%d\n", &T)
	for i := 0; i < T; i++ {
		fmt.Fscanf(reader, "%d\n", &n)
		var stock = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscanf(reader, "%d ", &stock[i])
		}
		fmt.Fprintf(writer, "%d\n", find(n, stock))
	}
}
