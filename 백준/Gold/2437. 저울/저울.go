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
	var n, compare int
	fmt.Fscanf(reader, "%d\n", &n)
	var nums = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &nums[i])
	}
	compare = 1
	sort.Ints(nums)
	for _, v := range nums {
		if compare < v {
			break
		}
		compare += v
	}
	fmt.Fprintf(writer, "%d\n", compare)
}
