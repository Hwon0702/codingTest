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

func check() {
	var n, m, num int
	fmt.Fscanf(reader, "%d\n", &n)
	var book1 = make(map[int]bool)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &num)
		book1[num] = true
	}
	fmt.Fscanf(reader, "%d\n", &m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &num)
		if _, f := book1[num]; !f {
			fmt.Fprintf(writer, "%d\n", 0)
		} else {
			fmt.Fprintf(writer, "%d\n", 1)
		}
	}
}

func main() {
	defer writer.Flush()
	var tc int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		check()
	}
}
