package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var res = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		res[i] = []int{0, 0, 0}
	}
	res[1][0] = 1
	res[1][1] = 1
	res[1][2] = 1
	for i := 2; i < n+1; i++ {
		res[i][0] = (res[i-1][0] + res[i-1][1] + res[i-1][2]) % 9901
		res[i][1] = (res[i-1][0] + res[i-1][2]) % 9901
		res[i][2] = (res[i-1][0] + res[i-1][1]) % 9901
	}
	fmt.Println((res[n][0] + res[n][1] + res[n][2]) % 9901)
}
