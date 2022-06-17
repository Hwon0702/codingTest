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
	fmt.Fscanf(reader, "%d\n", &n)
	var length = make([]int, n+1)
	var oil = make([]int, n+1)
	var nowOil = 0
	var price int
	for i := 2; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &length[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &oil[i])
	}
	nowOil = oil[1]
	price += oil[1] * length[2]

	for i := 3; i < len(length); i++ {
		if oil[i-1] < nowOil {
			nowOil = oil[i-1]
		}
		price += nowOil * length[i]

	}
	fmt.Println(price)
}
