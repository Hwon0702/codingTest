package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func sub(a int, b int) uint64 {
	res := a - b
	if res < 0 {
		return uint64(res * -1)
	}

	return uint64(res)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var rate = make([]int, n)
	var diffSum uint64
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &rate[i])
	}
	sort.Ints(rate)
	for i := 0; i < n; i++ {
		diffSum += sub(rate[i], i+1)
	}
	fmt.Fprintf(writer, "%d\n", diffSum)
}
