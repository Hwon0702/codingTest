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

func (s *Stack) Size() int {
	return len(s.Nodes)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var s, rem string
	fmt.Fscanf(reader, "%s\n%s", &s, &rem)
	var stack = new(Stack)
	sArr := strings.Split(s, "")
	remArr := strings.Split(rem, "")
	var last = remArr[len(remArr)-1]
	var len = len(remArr)
	for _, s := range sArr {
		stack.Push(s)
		if s == last && stack.Size()-len >= 0 && strings.Join(stack.Nodes[stack.Size()-len:], "") == rem {
			stack.Nodes = stack.Nodes[:stack.Size()-len]
		}
	}
	if stack.Size() > 0 {
		fmt.Println(strings.Join(stack.Nodes, ""))
	} else {
		fmt.Println("FRULA")
	}
}
