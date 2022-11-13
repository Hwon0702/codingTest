package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	//var n int
	var x, y, w, h int

	defer writer.Flush()
	fmt.Fscanf(reader, "%d %d %d %d\n", &x, &y, &w, &h)
	var minX, minY int
	if x > w-x {
		minX = w - x
	} else {
		minX = x
	}
	if y > h-y {
		minY = h - y
	} else {
		minY = y
	}
	if minX > minY {
		fmt.Fprintf(writer, "%d", minY)
	} else {
		fmt.Fprintf(writer, "%d", minX)
	}
}