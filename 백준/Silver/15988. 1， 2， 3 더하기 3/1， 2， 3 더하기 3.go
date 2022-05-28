package main

import (
	"bufio"
	"fmt"
	"os"
)

var numToCnt = map[int]int{
	1: 1,
	2: 2,
	3: 4,
	4: 7,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		var n int
		fmt.Fscanf(reader, "%d\n", &n)
		fmt.Fprintf(writer, "%d\n", Find(n))
	}
}

func Find(n int) int {
	if v, find := numToCnt[n]; find {
		return v
	} else {
		for i := 1; i <= n; i++ {
			if _, find := numToCnt[i]; !find {
				numToCnt[i] = (numToCnt[i-3] + numToCnt[i-2] + numToCnt[i-1]) % 1000000009
			}
		}
	}
	return numToCnt[n]
}
