package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	arr  = []int{}
	n, m int

	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func dfs(idx int) {
	if len(arr) == m {
		for _, v := range arr {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintf(writer, "\n")
		return
	}
	for i := idx; i <= n; i++ {
		arr = append(arr, i)
		dfs(i)
		if len(arr) > 0 {
			arr = arr[:len(arr)-1]
		}
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	dfs(1)
}
