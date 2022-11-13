package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += S1065(i)
	}
	fmt.Println(sum)
}

func S1065(n int) int {
	numStr := strings.Split(fmt.Sprintf("%d", n), "")
	d := 0
	if len(numStr) > 2 {
		a, _ := strconv.Atoi(numStr[0])
		b, _ := strconv.Atoi(numStr[1])
		d = b - a
		for j := 1; j < len(numStr)-1; j++ {
			a, _ = strconv.Atoi(numStr[j])
			b, _ = strconv.Atoi(numStr[j+1])
			if d != b-a {
				return 0
			}
		}
	}
	return 1
}
