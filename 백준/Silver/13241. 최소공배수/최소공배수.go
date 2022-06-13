package main

import (
	"bufio"
	"fmt"
	"os"
)

func getGcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	for a > 0 {
		a, b = b%a, a
	}
	return b
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var a, b int
	fmt.Fscanf(reader, "%d %d", &a, &b)
	gcd := getGcd(a, b)
	gcm := a * b / gcd
	fmt.Fprintf(writer, "%d", gcm)
}
