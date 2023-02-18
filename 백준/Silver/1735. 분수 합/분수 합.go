package main

import (
	"bufio"
	"fmt"
	"os"
)

func getGcd(a, b int) int {
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var a, b, c, d int
	fmt.Fscanf(reader, "%d %d\n%d %d\n", &a, &b, &c, &d)
	gcd := getGcd(b, d)
	gcm := b * d / gcd
	a = a * (gcm / b)
	c = c * (gcm / d)
	newGcd := getGcd(a+c, gcm)
	fmt.Fprintf(writer, "%d %d\n", (a+c)/newGcd, gcm/newGcd)
}
