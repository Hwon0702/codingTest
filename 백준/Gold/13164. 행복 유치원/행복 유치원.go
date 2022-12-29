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
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	var person = make([]int, n)
	var diff = make([]int, n-1)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &person[i])
	}
	for i := 1; i < n; i++ {
		diff[i-1] = person[i] - person[i-1]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(diff)))
	if n == k {
		fmt.Fprintf(writer, "%d\n", 0)
	} else {
		for i := k - 1; i < len(diff); i++ {
			res += diff[i]
		}
		fmt.Fprintf(writer, "%d\n", res)
	}
}
