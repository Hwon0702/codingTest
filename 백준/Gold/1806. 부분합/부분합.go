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

	var n, s, min, sum int
	min = 100000002
	fmt.Fscanf(reader, "%d %d\n", &n, &s)
	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	for i := 0; i < n; i++ {
		sum = numbers[i]
		if sum >= s {
			min = 1
			break
		}
		for j := i + 1; j < n; j++ {
			sum += numbers[j]
			if sum >= s {
				if j-i < min {
					min = j - i + 1
				}
				break
			}
		}
	}
	if min >= 100000002 {
		fmt.Fprintf(writer, "%d\n", 0)
	} else {
		fmt.Fprintf(writer, "%d\n", min)
	}
}
