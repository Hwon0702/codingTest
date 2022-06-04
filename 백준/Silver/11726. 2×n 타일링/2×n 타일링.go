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
	var num int
	fmt.Fscanf(reader, "%d\n", &num)
	var results = make([]int, 1001)
	for i := 0; i < 1001; i++ {
		results[i] = 0
	}
	results[1] = 1
	results[2] = 2
	if num >= 3 {
		for i := 3; i <= num; i++ {
			results[i] = results[i-2]%10007 + results[i-1]%10007
		}
	}
	fmt.Fprintf(writer, "%d\n", results[num]%10007)
}
