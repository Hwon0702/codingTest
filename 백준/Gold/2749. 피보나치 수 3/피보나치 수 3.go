package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	M      = 1000000
	P      = 1500000
)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	a, b := 0, 1
	for i := 0; i < n%P; i++ {
		a, b = b%M, (a+b)%M
	}
	fmt.Fprintf(writer, "%d\n", a)
}
