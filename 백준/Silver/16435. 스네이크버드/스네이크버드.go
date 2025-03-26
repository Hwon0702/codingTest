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
	var n, l int
	fmt.Fscanf(reader, "%d %d\n", &n, &l)
	var fruites = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &fruites[i])
	}
	sort.Slice(fruites, func(i, j int) bool {
		return fruites[i] < fruites[j]
	})
	for i := 0; i < n; i++ {
		if fruites[i] <= l {
			l++
		} else {
			break
		}
	}
	fmt.Fprintf(writer, "%d\n", l)
}
