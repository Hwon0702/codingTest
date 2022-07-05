package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l, r string
	fmt.Fscanf(reader, "%s %s\n", &l, &r)
	lArr := strings.Split(l, "")
	rArr := strings.Split(r, "")
	var res int
	if len(l) == len(r) {
		for i := 0; i < len(lArr); i++ {
			if lArr[i] == rArr[i] && lArr[i] == "8" {
				res++
			} else if lArr[i] != rArr[i] {
				break
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
