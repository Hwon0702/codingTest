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
	var n, m, tmp int
	var a, b int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var results = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &tmp)
		results[i] = tmp + results[i-1]
	}
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		fmt.Fprintf(writer, "%d\n", results[b]-results[a-1])
	}

}
