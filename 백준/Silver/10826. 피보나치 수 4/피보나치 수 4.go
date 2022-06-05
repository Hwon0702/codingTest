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
	fmt.Fscanf(reader, "%d\n", &n)
	var a, b big.Int
	var result big.Int
	a.SetInt64(1)
	b.SetInt64(1)
	if n == 0 {
		fmt.Fprintf(writer, "%d", 0)
	} else if n <= 2 {
		fmt.Fprintf(writer, "%d", 1)
	} else {
		for i := 0; i < n-2; i++ {
			result.Set(a.Add(&a, &b))
			a.Set(&b)
			b.Set(&result)
		}
		fmt.Fprintf(writer, "%d", &result)
	}
}
