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

func GCD(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
func main() {
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var nums = make([]int, n)
	var dist = make([]int, n-1)
	var gcds = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &nums[i])
	}
	for i := 0; i < n-1; i++ {
		dist[i] = nums[i+1] - nums[i]
	}
	for i := 0; i < len(dist)-1; i++ {
		gcds[i] = GCD(dist[i], dist[i+1])
	}
	gcd := GCD(gcds[0], gcds[1])
	for i := 1; i < n-1; i++ {
		gcd = GCD(gcd, gcds[i])
	}
	var cnt int
	for i := 0; i < len(dist); i++ {
		cnt += dist[i] / gcd
	}
	fmt.Fprintf(writer, "%d\n", cnt-n+1)

}
