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
	var x, y, c float64
	fmt.Fscanf(reader, "%f %f %f", &x, &y, &c)
	var high = float64(0)
	var low = float64(1)
	var ans = float64(0)

	if x > y {
		high = x
	} else {
		high = y
	}
	for low+0.001 <= high {
		w := (low + high) / 2
		h1 := math.Sqrt(x*x - w*w)
		h2 := math.Sqrt(y*y - w*w)
		test := (h1 * h2) / (h1 + h2)
		if test >= c {
			ans = w
			low = w
		} else {
			high = w
		}
	}
	fmt.Fprintf(writer, "%.3f", ans)
}
