package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s, rem string
	fmt.Fscanf(reader, "%s\n%s", &s, &rem)
	var stack = []byte{}
	sArr := []byte(s)
	remArr := strings.Split(rem, "")
	var last = remArr[len(remArr)-1]
	var length = len(remArr)
	for _, s := range sArr {
		stack = append(stack, s)
		if len(stack) < length {
			continue
		} else if string(s) == last && string(stack[len(stack)-length:]) == rem {
			stack = stack[:len(stack)-length]
		}
	}
	if len(stack) > 0 {
		fmt.Println(string(stack))
	} else {
		fmt.Println("FRULA")
	}
}
