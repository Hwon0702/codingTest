package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var left, right int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	end := false
	for !end {
		fmt.Fscanln(reader, &left, &right)

		end = left == 0 && right == 0
		if !end {
			fmt.Printf("%d\n", left+right)
		}
	}
	writer.Flush()
}
