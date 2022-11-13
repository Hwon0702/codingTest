package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc, n int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d\n", &n)
		var numToCnt = make(map[int]int)
		var keys = []int{}
		for n != 1 {
			if n%2 == 0 {
				n = n / 2
				if _, exists := numToCnt[2]; !exists {
					numToCnt[2] = 1
					keys = append(keys, 2)
				} else {
					numToCnt[2] += 1
				}
			} else {
				var i = 3
				for (n % i) != 0 {
					i++
				}
				if _, exists := numToCnt[i]; !exists {
					numToCnt[i] = 1
					keys = append(keys, i)
				} else {
					numToCnt[i] += 1
				}
				n = n / i
			}
		}
		sort.Ints(keys)
		for _, v := range keys {
			fmt.Fprintf(writer, "%d %d\n", v, numToCnt[v])
		}
	}
}
