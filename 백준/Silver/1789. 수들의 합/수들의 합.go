package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	var sum, cnt int

	fmt.Fscanf(reader, "%d\n", &n)
	for i := 1; i <= n; i++ {
		sum += i
		cnt++
		if sum > n {
			cnt -= 1
			break
		}
	}
	if cnt <= 0 {
		fmt.Println(0)
	} else {
		fmt.Println(cnt)
	}
}
