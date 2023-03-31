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
	if n%5 == 2 || n%5 == 0 {
		fmt.Fprintf(writer, "%s\n", "CY")
	} else {
		fmt.Fprintf(writer, "%s\n", "SK")
	}
}
