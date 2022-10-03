package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s string
	var n int
	fmt.Fscanf(reader, "%s\n", &s)
	sArr := strings.Split(s, "")
	for i := 1; i < len(sArr); i++ {
		if sArr[i-1] != sArr[i] {
			n++
		}
	}
	fmt.Println(int((n + 1) / 2))
}
