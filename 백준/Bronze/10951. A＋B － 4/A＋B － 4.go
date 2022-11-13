package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		var left, right int
		fmt.Fscanln(reader, &left, &right)
		if left == 0 || right == 0 {
			break
		}
		fmt.Printf("%d\n", left+right)
	}
	writer.Flush()
}
