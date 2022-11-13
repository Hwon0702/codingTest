package main

import (
	"bufio"
	"fmt"
	"os"
)

func B1003(n int) (zeroArr, oneArr []int) {
	zeroArr = []int{1, 0}
	oneArr = []int{0, 1}

	if n <= 1 {
		return
	}
	for i := 2; i < n+1; i++ {
		zeroArr = append(zeroArr, zeroArr[i-1]+zeroArr[i-2])
		oneArr = append(oneArr, oneArr[i-1]+oneArr[i-2])
	}
	return
}
func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		zeroCount, oneCount := B1003(n)
		fmt.Fprintf(writer, "%d %d\n", zeroCount[n], oneCount[n])
	}
}
