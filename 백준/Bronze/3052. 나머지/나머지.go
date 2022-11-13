package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	modulerMap := make(map[int]int)
	for i := 0; i < 10; i++ {
		var num int
		fmt.Fscanln(reader, &num)
		modulerMap[num%42]++
	}

	fmt.Fprintf(writer, "%d\n", len(modulerMap))
	writer.Flush()

}
