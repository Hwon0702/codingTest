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

func main() {
	defer writer.Flush()
	var n int
	var dp = make([][]uint64, 31)
	for i := 0; i < 31; i++ {
		dp[i] = make([]uint64, 31)
		dp[0][i] = 1
	}
	for i := 1; i < 31; i++ {
		for j := i; j < 31; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	for {
		fmt.Fscanf(reader, "%d\n", &n)
		if n == 0 {
			break
		}
		fmt.Fprintf(writer, "%d\n", dp[n][n])
	}
}
