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
	var wines = make([]int, n+1)
	var res = make([]int, n+1)
	res[0] = 0
	wines[0] = 0
	for i := 1; i <= n; i++ {
		res[i] = 0
		fmt.Fscanf(reader, "%d\n", &wines[i])
	}
	if n == 1 {
		fmt.Fprintf(writer, "%d\n", wines[1])
	} else if n == 2 {
		fmt.Fprintf(writer, "%d\n", wines[1]+wines[2])
	} else {
		res[1] = wines[1]
		res[2] = wines[1] + wines[2]
		for i := 3; i <= n; i++ {
			res[i] = int(math.Max(math.Max(float64(res[i-3]+wines[i-1]+wines[i]), float64(res[i-2]+wines[i])), float64(res[i-1])))
		}
		fmt.Fprintf(writer, "%d\n", res[n])
	}

}
