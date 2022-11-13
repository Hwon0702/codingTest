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
	var length = make([]int, n+1)
	var oil = make([]int, n+1)
	var result = make([]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &length[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &oil[i])
	}
	for i := 1; i < len(result); i++ {
		result[i] = 10000000
	}
	result[2] = oil[1] * length[2] //0부터 1까지 가는 비용. 처음엔 무조건 곱해야함
	for i := 3; i < len(length); i++ {
		result[i] = int(math.Min(float64(result[i-1]+oil[i-2]*length[i]), float64(result[i-1]+oil[i-1]*length[i])))
	}
	fmt.Println(result[n])
}
