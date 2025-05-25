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
	var n int
	var cnt = 0
	fmt.Fscanf(reader, "%d\n", &n)
	var grades = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &grades[i])
	}
	for i := n - 2; i > -1; i-- {
		if grades[i] >= grades[i+1] {
			cnt += grades[i] - (grades[i+1] - 1)
			grades[i] = grades[i+1] - 1
		}
	}
	fmt.Fprintf(writer, "%d\n", cnt)
}
