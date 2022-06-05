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
	var count int
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	for n >= 0 {
		if n%5 == 0 {
			count += int(n / 5)
			fmt.Fprintf(writer, "%d", count)
			break
		}
		n -= 3
		count++
	}
	if n < 0 {
		fmt.Fprintf(writer, "%d", -1)
	}
}
