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

	var a, b int

	for i := 0; i < length; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Printf("Case #%d: %d\n", i+1, a+b)
	}
	writer.Flush()
}
