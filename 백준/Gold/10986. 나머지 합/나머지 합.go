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
	var n, m, cnt int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var nums = make([]int, n+1)
	var mod = make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &nums[i])
	}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	for i := 0; i < n; i++ {
		mod[nums[i]%m] += 1
	}

	cnt = mod[0]
	for _, v := range mod {
		cnt += v * (v - 1) / 2
	}
	fmt.Fprintf(writer, "%d\n", cnt)
}
