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
	var a, b, c float64
	var min, max float64
	fmt.Fscanf(reader, "%d\n", &n)
	var maxRes, minRes, nowMax, nowMin = make([]float64, 3), make([]float64, 3), make([]float64, 3), make([]float64, 3)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%f %f %f\n", &a, &b, &c)
		for j := 0; j < 3; j++ {
			if j == 0 {
				nowMax[0] = a + math.Max(maxRes[0], maxRes[1])
				nowMin[0] = a + math.Min(minRes[0], minRes[1])
			} else if j == 1 {
				nowMax[1] = b + math.Max(maxRes[0], math.Max(maxRes[1], maxRes[2]))
				nowMin[1] = b + math.Min(minRes[0], math.Min(minRes[1], minRes[2]))
			} else if j == 2 {
				nowMax[2] = c + math.Max(maxRes[1], maxRes[2])
				nowMin[2] = c + math.Min(minRes[1], minRes[2])
			}
		}
		for j := 0; j < 3; j++ {
			maxRes[j] = nowMax[j]
			minRes[j] = nowMin[j]
		}
	}
	max = math.Max(maxRes[0], math.Max(maxRes[1], maxRes[2]))
	min = math.Min(minRes[0], math.Min(minRes[1], minRes[2]))
	fmt.Fprintf(writer, "%0.f %0.f\n", max, min)
}
