package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, f int
	fmt.Fscanf(reader, "%d\n%d", &n, &f)
	n = int(n/100) * 100
	for n%f != 0 {
		n += 1
	}
	if n%100 < 10 {
		fmt.Printf("0%d", n%100)
	} else {
		fmt.Printf("%d", n%100)
	}
}
