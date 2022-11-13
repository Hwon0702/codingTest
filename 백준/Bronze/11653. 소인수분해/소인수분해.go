package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n)
	defer writer.Flush()
	i := 2
	for i <= n {
		if n%i == 0 {
			fmt.Fprintf(writer, "%d\n", i)
			n = n / i

		} else {
			i++
		}
	}
}
