package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int

	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	var res = make([][]int, k+1)
	for i := 0; i <= k; i++ {
		res[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			res[i][j] = 1
		}
	}
	for i := 2; i <= k; i++ {
		for j := 1; j <= n; j++ {
			res[i][j] = (res[i-1][j] + res[i][j-1]) % 1000000000
		}
	}
	fmt.Println(res[k][n])
}
