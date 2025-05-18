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
	Node []string
}

func (s *Stack) Push(n string) {
	s.Node = append(s.Node, n)
}

func (s *Stack) Last() string {
	return s.Node[s.Len()-1]
}

func (s *Stack) Len() int {
	return len(s.Node)
}

func New() *Stack {
	st := new(Stack)
	return st
}

func (s *Stack) Pop() (error, string) {
	if s.Len() <= 0 {
		return fmt.Errorf("empty"), ""
	}
	last := s.Node[s.Len()-1]
	if s.Len() == 1 {
		s.Node = []string{}
	} else {
		s.Node = s.Node[:s.Len()-1]
	}
	return nil, last
}
func main() {
	defer writer.Flush()
	var n, res int
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Fscanf(reader, "%s\n", &str)
		strArr := strings.Split(str, "")
		s := New()
		for j := 0; j < len(strArr); j++ {
			if s.Len() <= 0 {
				s.Push(strArr[j])
			} else if s.Last() == strArr[j] {
				s.Pop()
			} else {
				s.Push(strArr[j])
			}
		}
		if s.Len() == 0 {
			res += 1
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
