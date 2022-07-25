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
	var len, max int
	fmt.Fscanf(reader, "%d\n", &len)
	var numbers = make([]int, len)
	var res = make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
		res[i] = 1
	}
	for i := 0; i < len; i++ {
		for j := 0; j < i; j++ {
			if numbers[i] < numbers[j] {
				if res[i] < res[j]+1 {
					res[i] = res[j] + 1
				}
			}
		}
	}
	for _, v := range res {
		if max < v {
			max = v
		}
	}
	fmt.Fprintf(writer, "%d\n", max)
}
