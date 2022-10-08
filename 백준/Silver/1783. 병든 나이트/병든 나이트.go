package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m float64
	fmt.Fscanf(reader, "%f %f\n", &n, &m)
	if n == 1 {
		fmt.Fprintf(writer, "%d\n", 1)
	} else if n == 2 {
		fmt.Fprintf(writer, "%0.f\n", math.Min(4, math.Floor((m+1)/2)))
	} else {
		if m <= 6 {
			fmt.Fprintf(writer, "%0.f\n", math.Min(4, m))
		} else {
			fmt.Fprintf(writer, "%0.f\n", m-2)
		}
	}
}
