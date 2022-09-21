package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var Xlist = make([]float64, n+1)
	var Ylist = make([]float64, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%f %f\n", &Xlist[i], &Ylist[i])
	}
	Xlist[n] = Xlist[0]
	Ylist[n] = Ylist[0]
	var sum float64
	for i := 0; i < len(Ylist)-1; i++ {
		sum += Xlist[i] * Ylist[i+1]
		sum -= Xlist[i+1] * Ylist[i]
	}
	fmt.Printf("%.1f", math.Abs(sum/2))
}
