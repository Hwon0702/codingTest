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
	fmt.Fscanf(reader, "%d\n", &n)
	var points = make([]int, n+1)
	var results = make([]int, n+1)
	results[0] = 0
	points[0] = 0
	for i := 1; i <= n; i++ {
		results[i] = 0
		fmt.Fscanf(reader, "%d\n", &points[i])
	}
	if n == 1 {
		fmt.Fprintf(writer, "%d\n", points[1])
	} else if n == 2 {
		fmt.Fprintf(writer, "%d\n", points[1]+points[2])
	} else {
		results[1] = points[1]
		results[2] = points[1] + points[2]
		results[3] = int(math.Max(float64(points[1]+points[3]), float64(points[2]+points[3])))
		for i := 4; i <= n; i++ {
			results[i] = int(math.Max(float64(results[i-3]+points[i-1]+points[i]), float64(results[i-2]+points[i])))
		}
		fmt.Fprintf(writer, "%d\n", results[n])
	}

}
