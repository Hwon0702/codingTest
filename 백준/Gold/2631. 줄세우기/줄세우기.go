package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var res = make([]float64, n)
	var numbers = make([]float64, n)
	var max float64
	for i := 0; i < n; i++ {
		res[i] = 1
		fmt.Fscanf(reader, "%f\n", &numbers[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if numbers[j] < numbers[i] {
				res[i] = math.Max(res[i], res[j]+1)
			}
		}
	}
	for _, v := range res {
		max = math.Max(v, max)
	}
	fmt.Fprintf(writer, "%0.f", float64(n)-max)
}
