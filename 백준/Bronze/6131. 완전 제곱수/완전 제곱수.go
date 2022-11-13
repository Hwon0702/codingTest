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

	var tc int
	fmt.Fscanf(reader, "%d\n", &tc)

	var cnt int
	for i := 1; i <= 500; i++ {
		for j := i; j <= 500; j++ {
			if j*j == i*i+tc {
				cnt++
			}
		}
	}
	fmt.Println(cnt)

}
