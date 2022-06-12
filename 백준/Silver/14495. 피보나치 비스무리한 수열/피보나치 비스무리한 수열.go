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

	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var res = make([]int, n+1)
	if n <= 3 {
		fmt.Fprintf(writer, "%d", 1)
	} else {
		for i := 0; i <= 3; i++ {
			res[i] = 1
		}
		for i := 4; i <= n; i++ {
			res[i] = res[i-1] + res[i-3]
		}
		fmt.Fprintf(writer, "%d", res[n])
	}
}
