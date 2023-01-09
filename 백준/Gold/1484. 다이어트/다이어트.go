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

func cal(a, b int) int {
	return a*a - b*b
}

func main() {
	defer writer.Flush()
	var g, v int
	var ans bool
	fmt.Fscanf(reader, "%d\n", &g)
	var l, r = 1, 1
	for l+r <= g {
		v = cal(l, r)
		if v == g {
			fmt.Fprintf(writer, "%d\n", l)
			l += 1
			ans = true
		} else if v > g {
			r += 1
		} else if v < g {
			l += 1
		}
	}
	if !ans {
		fmt.Fprintf(writer, "%d\n", -1)
	}
}
