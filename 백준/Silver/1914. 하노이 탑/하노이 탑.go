package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanln(reader, &n)
	var i, e = big.NewInt(2), big.NewInt(int64(n))
	i.Exp(i, e, nil)
	sub := new(big.Int)
	sub = sub.Sub(i, big.NewInt(1))
	fmt.Println(sub)
	if n <= 20 {
		Hanoi(writer, n, 1, 3, 2)
	}
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
