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
	var n, tc int
	var start, end int
	fmt.Fscanf(reader, "%d\n", &n)
	var numbers = make([]int, n+1)
	var numbersSum = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	numbersSum[1] = numbers[1]
	for i := 1; i <= n; i++ {
		numbersSum[i] = numbers[i] + numbersSum[i-1]
	}
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d %d\n", &start, &end)
		fmt.Fprintf(writer, "%d\n", numbersSum[end]-numbersSum[start-1])
	}

}
