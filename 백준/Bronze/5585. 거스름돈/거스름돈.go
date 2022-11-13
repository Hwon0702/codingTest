package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
500엔, 100엔, 50엔, 10엔, 5엔, 1
*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var changes = []int{500, 100, 50, 10, 5, 1}
	var cnt = 0
	var remaincnt int
	var remain = 1000 - n
	for true {
		if remain == 0 {
			break
		}
		remaincnt += (remain / changes[cnt])
		remain = remain - (changes[cnt] * (remain / changes[cnt]))
		cnt++
	}
	fmt.Fprintf(writer, "%d", remaincnt)
}
