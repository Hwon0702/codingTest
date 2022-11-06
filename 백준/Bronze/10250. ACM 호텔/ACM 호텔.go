package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var tc int
	var h, w, n, floor, room int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &h, &w, &n)
		floor = 0
		room = 0
		if n%h == 0 {
			floor = h * 100
			room = int(n / h)
		} else {
			floor = (n % h) * 100
			room = 1 + n/h
		}
		fmt.Fprintf(writer, "%d\n", room+floor)
	}
}
