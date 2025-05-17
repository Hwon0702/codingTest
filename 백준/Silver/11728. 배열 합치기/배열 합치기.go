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
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var list = make([]int, n+m)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &list[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &list[n+i])
	}
	sort.Ints(list)
	for i := 0; i < len(list); i++ {
		fmt.Fprintf(writer, "%d ", list[i])
	}
}
