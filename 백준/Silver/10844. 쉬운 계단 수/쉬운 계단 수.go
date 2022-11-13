package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	var sum uint64
	fmt.Fscanf(reader, "%d\n", &n)
	var res = make([][]uint64, n+1)
	for i := 0; i < n+1; i++ {
		res[i] = make([]uint64, 10)
	}
	for i := 1; i < 10; i++ {
		res[1][i] = 1
	}
	for i := 2; i < n+1; i++ {
		res[i][0] = res[i-1][1]
		res[i][9] = res[i-1][8]
		for j := 1; j < 9; j++ {
			res[i][j] = (res[i-1][j-1] + res[i-1][j+1]) % 1000000000
		}
	}
	for _, v := range res[n] {
		sum += v
	}
	fmt.Fprintf(writer, "%d\n", sum%1000000000)
}
