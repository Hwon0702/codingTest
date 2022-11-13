package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Reverse(s string) int {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	num, _ := strconv.Atoi(result)
	return num
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b string
	fmt.Fscanln(reader, &a, &b)
	over := 0
	if Reverse(a) > Reverse(b) {
		over = Reverse(a)
	} else {
		over = Reverse(b)
	}
	fmt.Println(over)
}
