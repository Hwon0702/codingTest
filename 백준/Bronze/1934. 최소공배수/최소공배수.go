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
	var n int
	var a, b int

	defer writer.Flush()
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {

		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		gcd := Gcd(a, b)
		gcm := a * b / gcd
		fmt.Fprintf(writer, "%d\n", gcm)
	}
}