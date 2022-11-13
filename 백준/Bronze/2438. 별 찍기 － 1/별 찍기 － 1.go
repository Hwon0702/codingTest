package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var length int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &length)
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			fmt.Fprintf(writer, "*")
		}
		fmt.Fprintf(writer, "\n")
	}
	writer.Flush()
}
