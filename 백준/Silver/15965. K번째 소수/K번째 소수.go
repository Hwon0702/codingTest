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
	var cnt = 0
	fmt.Fscanf(reader, "%d\n", &n)
	var res = make([]bool, 10000000)
	for i := 0; i < len(res); i++ {
		res[i] = true
	}

	for i := 2; i < 10000000; i++ {
		if res[i] == true {
			cnt++
			if cnt == n {
				fmt.Fprintf(writer, "%d\n", i)
				break
			}
			for j := 2; j*i < 10000000; j++ {
				if res[j*i] == true {
					res[j*i] = false

				}
			}
		}

	}
}
