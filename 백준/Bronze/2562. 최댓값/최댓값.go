package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	cnt := 1
	max := 0
	for i := 1; i < 10; i++ {
		var num int
		fmt.Fscanln(reader, &num)
		if num > max {
			cnt = i
			max = num
		}

	}
	fmt.Fprintf(writer, "%d\n%d \n", max, cnt)

	writer.Flush()

}
