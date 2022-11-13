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
	var n int
	fmt.Fscanf(reader, "%d \n", &n)

	person := make(map[int][]int)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		person[i] = []int{a, b}
	}
	for i := 0; i < n; i++ {
		var rank int
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if person[i][0] < person[j][0] && person[i][1] < person[j][1] {
				rank++
			}
		}
		fmt.Printf("%d ", rank+1)
	}
}
