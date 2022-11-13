package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n, count int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for true {
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}
		var primeNumber = make(map[int]bool, n)
		if n == 1 {
			fmt.Fprintln(writer, 1)
			continue
		}
		for i := n + 1; i <= 2*n; i++ {
			primeNumber[i] = true
		}

		for i := n + 1; i <= 2*n; i++ {
			if primeNumber[i] == false {
				break
			}
			for j := 1; j < int(math.Sqrt(float64(i+1))); j++ {
				if (i)%(j+1) == 0 {
					for k := 1; k*i <= 2*n; k++ {
						primeNumber[k*i] = false
					}
					break
				}
			}
		}

		count = 0
		for _, v := range primeNumber {
			if v {
				count++
			}
		}
		fmt.Fprintln(writer, count)
	}
}
