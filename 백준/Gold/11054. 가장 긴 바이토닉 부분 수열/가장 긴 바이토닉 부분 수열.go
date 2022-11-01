package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	var f, res float64
	fmt.Fscanf(reader, "%d\n", &n)
	var nums = make([]float64, n)
	var reverse = make([]float64, n)
	var increase = make([]float64, n)
	var decrease = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%f ", &f)
		nums[i] = f
		reverse[n-i-1] = f
		increase[i] = 1
		decrease[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				increase[i] = math.Max(increase[i], increase[j]+1)
			}
			if reverse[i] > reverse[j] {
				decrease[i] = math.Max(decrease[i], decrease[j]+1)
			}
		}
	}
	for i := 0; i < n; i++ {
		res = math.Max(res, increase[i]+decrease[n-i-1]-1)
	}
	fmt.Fprintf(writer, "%0.f\n", res)
}
