package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack struct {
	Nodes []string
}

func (s *Stack) Push(c string) {
	s.Nodes = append(s.Nodes, c)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s, rem string
	fmt.Fscanf(reader, "%s\n%s", &s, &rem)
	var stack = new(Stack)
	sArr := strings.Split(s, "")
	remArr := strings.Split(rem, "")
	var last = remArr[len(remArr)-1]
	var length = len(remArr)
	for _, s := range sArr {
		stack.Push(s)
		if s == last && len(stack.Nodes)-length >= 0 && strings.Join(stack.Nodes[len(stack.Nodes)-length:], "") == rem {
			stack.Nodes = stack.Nodes[:len(stack.Nodes)-length]
		}
	}
	if len(stack.Nodes) > 0 {
		fmt.Println(strings.Join(stack.Nodes, ""))
	} else {
		fmt.Println("FRULA")
	}
}
