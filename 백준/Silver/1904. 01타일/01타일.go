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
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var res = make([]int, n+1)
	if n <= 2 {
		fmt.Fprintf(writer, "%d\n", n)
	} else {
		res[1] = 1
		res[2] = 2
		for i := 3; i <= n; i++ {
			res[i] = (res[i-1] + res[i-2]) % 15746
		}
		fmt.Fprintf(writer, "%d\n", res[n])
	}
	defer writer.Flush()
}
