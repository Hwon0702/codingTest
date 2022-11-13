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

	var n, m, tmp, sum int
	var x1, y1, x2, y2 int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var arrSum = make([][]int, n+1)
	var numbersSum = make([]int, n+1)
	for i := 1; i <= n; i++ {
		numbersSum = make([]int, n+1)
		for j := 1; j <= n; j++ {
			fmt.Fscanf(reader, "%d ", &tmp)
			numbersSum[j] = numbersSum[j-1] + tmp
		}
		arrSum[i] = make([]int, n+1)
		arrSum[i] = numbersSum
	}
	for i := 0; i < m; i++ {
		sum = 0
		fmt.Fscanf(reader, "%d %d %d %d\n", &y1, &x1, &y2, &x2)
		for j := y1; j <= y2; j++ {
			sum += arrSum[j][x2] - arrSum[j][x1-1]
		}
		fmt.Fprintf(writer, "%d\n", sum)
	}

}
