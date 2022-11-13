package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanln(reader, &n)
	fmt.Fprintln(writer, int(math.Pow(2, float64(n)))-1)
	Hanoi(writer, n, 1, 3, 2)
}

func Hanoi(writer *bufio.Writer, n, from, to, aux int) {
	if n == 1 {

		fmt.Fprintln(writer, from, to)
		return
	}
	Hanoi(writer, n-1, from, aux, to)
	fmt.Fprintln(writer, from, to)
	Hanoi(writer, n-1, aux, to, from)

}
