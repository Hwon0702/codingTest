package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanln(reader, &n)
	numArr := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscanln(reader, &num)
		numArr[i] = num
	}

	sort.Ints(numArr)
	for _, num := range numArr {
		fmt.Fprintln(writer, num)

	}
	writer.Flush()
}
