package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	var mod5, mod25, mod125 int
	fmt.Fscanf(reader, "%d", &n)
	mod5 = n / 5
	mod25 = n / 25
	mod125 = n / 125
	fmt.Println(mod5 + mod25 + mod125)
}
