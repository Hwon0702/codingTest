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
	var tc, n, m int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d\n", &n)
		var coins = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &coins[j])
		}
		fmt.Fscanf(reader, "%d\n", &m)
		var dp = make([]int, m+1)
		dp[0] = 1
		for _, v := range coins {
			for j := 0; j <= m; j++ {
				if j >= v {
					dp[j] += dp[j-v]
				}
			}
		}
		fmt.Fprintf(writer, "%d\n", dp[m])
	}
}
