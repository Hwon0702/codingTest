package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var cards = make([]float64, n+1)
	var dp = make([]float64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%f ", &cards[i])
		dp[i] = 0
	}
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			if dp[i] == 0 {
				dp[i] = dp[i-j] + cards[j]
			} else {
				dp[i] = math.Min(dp[i], dp[i-j]+cards[j])
			}
		}
	}
	fmt.Fprintf(writer, "%0.f\n", dp[n])
}
