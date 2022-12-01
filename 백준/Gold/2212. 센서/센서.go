package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n, k, res int
	fmt.Fscanf(reader, "%d\n%d\n", &n, &k)
	var nums = make([]int, n)
	var dist = make([]int, n-1)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &nums[i])
	}
	sort.Ints(nums)
	for i := 1; i < n; i++ {
		dist[i-1] = nums[i] - nums[i-1]
	}
	sort.Ints(dist)
	for i := 0; i < n-k; i++ {
		res += dist[i]
	}
	fmt.Fprintf(writer, "%d\n", res)
}
