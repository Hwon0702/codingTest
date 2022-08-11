package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, M int
	var S string
	var res, i, count int
	fmt.Fscanf(reader, "%d\n%d\n", &N, &M)
	fmt.Fscanf(reader, "%s\n", &S)
	for i < (M - 1) {
		if i+3 < len(S) && S[i:i+3] == "IOI" {
			i += 2
			count += 1
			if count == N {
				res += 1
				count -= 1
			}
		} else {
			i += 1
			count = 0
		}
	}
	fmt.Println(res)
}
