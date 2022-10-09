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
	var tc, res int
	fmt.Fscanf(reader, "%d\n", &tc)
	var x1, y1, r1, x2, y2, r2, dist float64
	for i := 0; i < tc; i++ {
		res = 0
		fmt.Fscanf(reader, "%f %f %f %f %f %f\n", &x1, &y1, &r1, &x2, &y2, &r2)
		dist = math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
		if dist == 0 && r1 == r2 {
			res = -1
		} else if math.Abs(r1-r2) == dist || r1+r2 == dist {
			res = 1
		} else if math.Abs(r1-r2) < dist && dist < r1+r2 {
			res = 2
		}
		fmt.Fprintf(writer, "%d\n", res)
	}
}
