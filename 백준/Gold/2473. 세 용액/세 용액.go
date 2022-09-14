package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, start, end int
	fmt.Fscanf(reader, "%d\n", &n)
	var requids = make([]float64, n)
	var result = make([]float64, 3)
	var min = math.MaxFloat64
	var tmpMin float64
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%f ", &requids[i])
	}
	sort.Float64s(requids)
	for mid := 0; mid < n-2; mid++ {
		start = mid + 1
		end = n - 1
		for start < end {
			tmpMin = requids[mid] + requids[start] + requids[end]
			if math.Abs(tmpMin) < min {
				min = math.Abs(tmpMin)
				result = []float64{requids[mid], requids[start], requids[end]}
			}
			if tmpMin < 0 {
				start++
			} else if tmpMin > 0 {
				end--
			} else {
				break
			}
		}
	}
	sort.Float64s(result)
	fmt.Fprintf(writer, "%0.f %0.f %0.f", result[0], result[1], result[2])
}
