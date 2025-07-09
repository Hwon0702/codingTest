package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	var n, cnt float64
	fmt.Fscanf(reader, "%f\n", &n)
	cnt = math.Sqrt(n)
	fmt.Fprintf(writer, "%0.f\n", math.Floor(cnt))
}
