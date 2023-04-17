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
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n/4; i++ {
		fmt.Fprintf(writer, "long ")
	}
	fmt.Fprintf(writer, "int")
}
