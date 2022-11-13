package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	nodes []int
}

func (s *Stack) Push(n int) {
	s.nodes = append(s.nodes, n)
}
func (s *Stack) Pop() (ret int) {
	if len(s.nodes) < 1 {
		ret = 0
	} else {
		ret = s.nodes[len(s.nodes)-1]
		if len(s.nodes) <= 1 {
			s.nodes = []int{}
		} else {
			s.nodes = s.nodes[:len(s.nodes)-1]
		}
	}
	return ret
}
func (s *Stack) Size() int {
	return len(s.nodes)
}

func (s *Stack) Empty() int {
	if len(s.nodes) > 0 {
		return 0
	} else {
		return 1
	}
}

func (s *Stack) Top() int {
	if len(s.nodes) < 1 {
		return -1
	} else {
		return s.nodes[len(s.nodes)-1]
	}
}

func main() {
	S := new(Stack)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscanln(reader, &num)
		if num > 0 {
			S.Push(num)
		} else {
			_ = S.Pop()
		}
	}
	var sum int
	for S.Size() > 0 {
		sum += S.Pop()
	}
	fmt.Fprintln(writer, sum)
	writer.Flush()
}
