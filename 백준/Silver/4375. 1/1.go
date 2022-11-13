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
	var err error
	for err == nil {
		_, err = fmt.Fscanf(reader, "%d\n", &n)
		if n != 0 && err == nil {
			var moduler = 0
			var counter = 0
			for true {
				counter++
				moduler = moduler*10 + 1
				moduler %= n
				if moduler == 0 {
					fmt.Fprintf(writer, "%d\n", counter)
					break
				}
			}
		}
	}
}
