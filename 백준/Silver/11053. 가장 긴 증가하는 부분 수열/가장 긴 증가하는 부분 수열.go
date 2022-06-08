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
	var res = make([]int, n)
	var numbers = make([]int, n)
	var max int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if numbers[i] > numbers[j] && res[i] < res[j] {
				res[i] = res[j]
			}
		}
		res[i] += 1
	}
	for _, v := range res {
		if v > max {
			max = v
		}
	}
	fmt.Fprintf(writer, "%d\n", max)
}
