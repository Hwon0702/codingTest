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
	var res = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		res[i] =1
	}
	res[1]=3
	for i := 2; i < n+1; i++ {
		res[i] = (res[i-1]*2 + res[i-2]) %9901
	}
	fmt.Println(res[n])
}
