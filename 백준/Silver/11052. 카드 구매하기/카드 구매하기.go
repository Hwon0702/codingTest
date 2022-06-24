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
	var numbers = make([]int, n+1)
	var results = make([]int, n+1)

	//var count int
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	if n == 1 {
		fmt.Fprintf(writer, "%d", numbers[1])
	} else {

		results[1] = numbers[1]
		for i := 2; i <= n; i++ {
			results[i] = numbers[i] //기존값을 넣어두고
			for j := 1; j < i; j++ {
				//result를 만들때까지의 모든 경우의 수에 대하여 검사해서 최댓값을 넣음
				//ex. 5를 만드는 방법은 1*5, 2+1*2, 3+2, 3+1*2 등 다양함
				results[i] = int(math.Max(float64(results[i]), float64(results[i-j]+numbers[j])))
			}
		}
		fmt.Fprintf(writer, "%d\n", results[n])
	}
}
