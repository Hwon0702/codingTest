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

func calc(m, n, x, y int) int {
	k := x
	for k <= m*n {
		if (k-x)%m == 0 && (k-y)%n == 0 {
			return k
		}
		k += m
	}
	return -1
}
func main() {
	defer writer.Flush()
	var t int
	var m, n, x, y int
	fmt.Fscanf(reader, "%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(reader, "%d %d %d %d\n", &m, &n, &x, &y)
		fmt.Fprintf(writer, "%d\n", calc(m, n, x, y))
	}
}
