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

	var n, k, r int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)

	var ryan = []int{}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &r)
		if r == 1 {
			ryan = append(ryan, i)
		}
	}
	if len(ryan) < k {
		fmt.Fprintf(writer, "%d\n", -1)
	} else {
		var start = 0
		var end = k - 1
		var ryanCount = 0
		var min = 1000001
		for start != end {
			ryanCount = ryan[end] - ryan[start] + 1
			if ryanCount < min {
				min = ryanCount
			}
			if end == len(ryan)-1 {
				break
			}
			start++
			end++
		}
		fmt.Fprintf(writer, "%d\n", min)
	}
}
