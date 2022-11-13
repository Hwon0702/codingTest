package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/*
https://www.acmicpc.net/problem/1463
*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var DP = make([]float64, 1000001)

	for i := 2; i <= n; i++ {
		DP[i] = DP[i-1] + 1

		if i%2 == 0 {
			DP[i] = math.Min(DP[i], DP[i/2]+1)

		}
		if i%3 == 0 {
			DP[i] = math.Min(DP[i], DP[i/3]+1)
		}
	}
	fmt.Fprintf(writer, "%d", int(DP[n]))
}
