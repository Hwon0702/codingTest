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
	var dp [][]int
	var nums []int
	fmt.Fscanf(reader, "%d\n", &n)

	dp = make([][]int, n)
	nums = make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 21)
	}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &nums[i])
	}
	dp[0][nums[0]] = 1
	for i := 1; i < n-1; i++ {
		for j := 0; j < 21; j++ {
			if 0 <= j+nums[i] && j+nums[i] <= 20 {
				dp[i][j+nums[i]] += dp[i-1][j]
			}
			if 0 <= j-nums[i] && j-nums[i] <= 20 {
				dp[i][j-nums[i]] += dp[i-1][j]
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", dp[n-2][nums[len(nums)-1]])
}
