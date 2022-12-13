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
	var n, m, x, res int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var frimes = make([]bool, m-n+1)
	for i := 0; i < len(frimes); i++ {
		frimes[i] = true
	}
	for i := 2; i < int(math.Sqrt(float64(m)))+1; i++ {
		x = int(n / (i * i))
		if n%(i*i) != 0 {
			x += 1
		}
		for x*(i*i) <= m && 0 <= x*(i*i)-n && x*(i*i)-n < m {
			frimes[x*(i*i)-n] = false
			x += 1
		}
	}

	for i := 0; i < m-n+1; i++ {
		if frimes[i] {
			res++
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
