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
	var a, b, c, d, e, f int
	fmt.Fscanf(reader, "%d %d %d %d %d %d\n", &a, &b, &c, &d, &e, &f)

	x := (c*e - b*f) / (a*e - b*d)
	y := (c*d - a*f) / (b*d - a*e)
	fmt.Fprintf(writer, "%d %d\n", x, y)

}
