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
	var e, s, m int
	//(1 ≤ E ≤ 15, 1 ≤ S ≤ 28, 1 ≤ M ≤ 19)
	fmt.Fscanf(reader, "%d %d %d\n", &e, &s, &m)
	if e == s && s == m {
		fmt.Fprintf(writer, "%d\n", e)
	} else {
		var year = 1
		e %= 15
		s %= 28
		m %= 19
		for true {
			if year%15 == e && year%28 == s && year%19 == m {
				fmt.Fprintf(writer, "%d\n", year)
				break
			} else {
				year++
			}
		}
	}
}
