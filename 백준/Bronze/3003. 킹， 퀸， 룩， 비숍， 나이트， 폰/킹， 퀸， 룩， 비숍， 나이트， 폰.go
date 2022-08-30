package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var right = []int{1, 1, 2, 2, 2, 8}
	var search = make([]int, 6)
	for i := 0; i < 6; i++ {
		fmt.Fscanf(reader, "%d ", &search[i])
	}
	for i := 0; i < 6; i++ {
		right[i] = right[i] - search[i]
		fmt.Printf("%d ", right[i])
	}
}
