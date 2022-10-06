package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, cnt int
	fmt.Fscanf(reader, "%d\n", &n)
	for true {
		if n%5 == 0 {
			cnt += int(n / 5)
			fmt.Println(cnt)
			break
		} else {
			n -= 2
			cnt += 1
		}
		if n < 0 {
			fmt.Println(-1)
			break
		}
	}
}
