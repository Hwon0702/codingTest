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
	if n <= 2 {
		fmt.Fprintf(writer, "%d", 1)
	} else {
		var results = make([]int, n+1)
		results[1] = 1
		results[2] = 1
		for i := 3; i <= n; i++ {
			results[i] = results[i-2] + results[i-1]
		}
		fmt.Fprintf(writer, "%d\n", results[n])
	}
}
