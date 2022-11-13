package main

import (
	"bufio"
	"fmt"
	"os"
)

func move(x, y int) {
	dist := y - x
	var n = 0
	for {
		if dist <= n*(n+1) {
			break
		}
		n += 1
	}
	if dist <= n*n {
		fmt.Println(n*2 - 1)
	} else {
		fmt.Println(n * 2)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	var x, y int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d %d\n", &x, &y)
		move(x, y)
	}
}
