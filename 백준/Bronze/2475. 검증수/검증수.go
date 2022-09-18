package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, sum int
	for i := 0; i < 5; i++ {
		fmt.Fscanf(reader, "%d ", &a)
		sum += a * a
	}
	fmt.Println(sum % 10)
}
