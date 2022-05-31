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
	var n, k int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	var monies = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &monies[n-1-i])
	}
	var cnt int
	for _, v := range monies {
		if k == 0 {
			break
		}
		cnt += k / v
		k = k % v
	}
	fmt.Fprintf(writer, "%d", cnt)
}
