package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type Stack struct {
	Nodes []string
}

func (s *Stack) Push(n string) {
	s.Nodes = append(s.Nodes, n)
}
func (s *Stack) Size() int {
	return len(s.Nodes)
}
func (s *Stack) Pop() (r string) {
	if len(s.Nodes) <= 0 {
		return ""
	}
	r = s.Nodes[0]
	if len(s.Nodes) == 1 {
		s.Nodes = []string{}
	} else {
		s.Nodes = s.Nodes[1:]
	}
	return
}
func main() {
	defer writer.Flush()
	var bar string
	var res int
	bar, _ = reader.ReadString('\n')
	barArr := strings.Split(strings.TrimSpace(bar), "")
	var stack = new(Stack)
	for i := 0; i < len(barArr); i++ {
		if barArr[i] == "(" {
			stack.Push(barArr[i])
		} else if stack.Size() > 0 {
			if barArr[i-1] == "(" {
				stack.Pop()
				res += len(stack.Nodes)
			} else {
				stack.Pop()
				res += 1
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}