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

	var n, k int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	var numbers = make([]bool, n+1)
	if n <= 2 {
		fmt.Fprintf(writer, "%d", 2)
	} else {
		for i := 2; i <= n; i++ {
			numbers[i] = true
		}
		var removeCnt int

		for i := 2; i <= n; i++ {
			if numbers[i] == true {
				removeCnt++
				numbers[i] = false
				if removeCnt == k {
					fmt.Println(i)
					break
				}
				for j := 2; j*i <= n; j++ {
					if numbers[j*i] == true {
						numbers[j*i] = false
						removeCnt++
						if removeCnt == k {
							fmt.Println(j * i)
							break
						}
					}
				}
			}

		}
	}
}
