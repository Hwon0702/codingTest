package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var x, y int
	fmt.Fscanf(reader, "%d %d\n", &x, &y)
	if x == y {
		fmt.Fprintf(writer, "%d\n", 0)
	} else {
		dist := y - x
		var n = 0
		for {
			if dist <= n*(n+1) {
				break
			}
			n += 1
		}
		if dist <= n*n {
			fmt.Fprintf(writer, "%d\n", n*2-1)
		} else {
			fmt.Fprintf(writer, "%d\n", n*2)
		}
	}
}
