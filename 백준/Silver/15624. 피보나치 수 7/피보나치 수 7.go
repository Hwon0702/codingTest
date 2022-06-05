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
	if n == 0 {
		fmt.Fprintf(writer, "%d", 0)
	} else if n == 1 {
		fmt.Fprintf(writer, "%d", 1)
	} else {
		var fibArr = make([]int, n+1)
		fibArr[0] = 0
		fibArr[1] = 1
		for i := 2; i <= n; i++ {
			fibArr[i] = fibArr[i-2]%1000000007 + fibArr[i-1]%1000000007
		}
		fmt.Fprintf(writer, "%d", fibArr[n]%1000000007)
	}
}
