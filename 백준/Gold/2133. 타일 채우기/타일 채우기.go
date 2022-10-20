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
	var res = make([]int, 31)
	if n%2 == 1 {
		fmt.Fprintf(writer, "%d\n", 0)
	} else {
		res[0] = 1
		res[2] = 3
		for i := 4; i <= n; i++ {
			res[i] = res[i-2] * 3
			for j := 0; j < i-2; j += 2 {
				res[i] += res[j] * 2
			}
		}
		fmt.Fprintf(writer, "%d\n", res[n])
	}
}
