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
	fmt.Fscanf(reader, "%d\n", &n)
	var dp = make([]uint64, 10)
	for i := 0; i < 10; i++ {
		dp[i] = 1
	}
	var sum uint64
	for i := 1; i < n; i++ {
		for j := 1; j < 10; j++ {
			dp[j] += (dp[j-1]) % 10007
		}
	}
	for _, v := range dp {
		sum += v
	}
	fmt.Fprintf(writer, "%d\n", sum%10007)
}
