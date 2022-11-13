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
	if n == 1 {
		fmt.Fprintf(writer, "%d\n", 1)
	} else if n == 2 {
		fmt.Fprintf(writer, "%d\n", 3)
	} else {
		var results = make([]int, n+1)
		results[1] = 1
		results[2] = 3
		for i := 3; i <= n; i++ {
			results[i] = (results[i-1] + results[i-2]*2) % 10007
		}
		fmt.Fprintf(writer, "%d\n", results[n]%10007)
	}
}
