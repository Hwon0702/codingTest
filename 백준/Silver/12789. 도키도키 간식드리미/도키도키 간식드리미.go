package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type Stack struct {
	v []int
}

func NewStack() *Stack {
	q := Stack{make([]int, 0)}
	return &q
}
func (s *Stack) Push(n int) {
	s.v = append(s.v, n)
}

func (s *Stack) Pop() int {
	if len(s.v) <= 0 {
		return -1
	}
	f := s.v[len(s.v)-1]
	s.v = s.v[:len(s.v)-1]
	return f
}

func (s *Stack) Last() int {
	if len(s.v) <= 0 {
		return -1
	}
	f := s.v[len(s.v)-1]
	return f
}
func (s *Stack) Empty() bool {
	if len(s.v) <= 0 {
		return true
	}
	return false
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	people := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &people[i])
	}
	stack := NewStack()
	last := 1
	for i := 0; i < len(people); i++ {
		stack.Push(people[i])
		for !stack.Empty() && stack.Last() == last {
			stack.Pop()
			last += 1
		}
	}
	if stack.Empty() {
		fmt.Fprintf(writer, "%s", "Nice")
	} else {
		fmt.Fprintf(writer, "%s", "Sad")
	}
}
