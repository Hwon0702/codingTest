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
	var tc int
	var results = make([]int, 101)
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 1; i <= 3; i++ {
		results[i] = 1
	}
	for i := 0; i < tc; i++ {
		var n int
		fmt.Fscanf(reader, "%d\n", &n)
		if n <= 3 {
			fmt.Fprintf(writer, "%d\n", 1)
		} else {
			if results[n] > 1 {
				fmt.Fprintf(writer, "%d\n", results[n])
			} else {
				for i := 4; i <= n; i++ {
					results[i] = results[i-3] + results[i-2]
				}
				fmt.Fprintf(writer, "%d\n", results[n])
			}
		}
	}
}
