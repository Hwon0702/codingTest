package main

import (
	"bufio"
	"fmt"
	"os"
)

func Gcd(a, b int) int {
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var a, b, res int
	fmt.Fscanf(reader, "%d %d\n", &a, &b)
	res = Gcd(a, b)
	for i := 0; i < res; i++ {
		fmt.Fprintf(writer, "%d", 1)
	}
}
