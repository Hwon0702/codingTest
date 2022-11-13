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
	xCnt := make(map[int]int, 2)
	yCnt := make(map[int]int, 2)

	for i := 0; i < 3; i++ {
		var x, y int
		fmt.Fscanf(reader, "%d %d\n", &x, &y)
		xCnt[x]++
		yCnt[y]++
	}
	var x, y int
	for key, val := range xCnt {
		if val < 2 {
			x = key
		}
	}
	for key, val := range yCnt {
		if val < 2 {
			y = key
		}
	}
	fmt.Println(x, y)
}
