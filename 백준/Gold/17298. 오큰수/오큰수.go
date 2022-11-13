package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	Values  []int
	Indexes []int
}

func (s *Stack) Push(value, index int) {
	s.Values = append(s.Values, value)
	s.Indexes = append(s.Indexes, index)
}
func (s *Stack) Pop() (value, index int) {
	if len(s.Values) <= 0 {
		return -1, -1
	}
	value = s.Values[len(s.Values)-1]
	index = s.Indexes[len(s.Indexes)-1]
	if len(s.Values) == 1 {
		s.Values = []int{}
		s.Indexes = []int{}
	} else {
		s.Values = s.Values[:len(s.Values)-1]
		s.Indexes = s.Indexes[:len(s.Indexes)-1]
	}
	return value, index
}

func (s *Stack) Size() int {
	return len(s.Values)
}
func (s *Stack) Last() (int, int) {
	if len(s.Values) <= 0 {
		return -1, -1
	}
	return s.Values[len(s.Values)-1], s.Indexes[len(s.Indexes)-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	var stack = new(Stack)
	fmt.Fscanf(reader, "%d\n", &tc)
	var result = make([]int, tc)
	for i := 0; i < tc; i++ {
		result[i] = -1
	}
	for i := 0; i < tc; i++ {
		var num int
		fmt.Fscanf(reader, "%d ", &num)
		values, index := stack.Last()
		if values < num {
			for values < num && values > 0 {
				stack.Pop()
				result[index] = num
				values, index = stack.Last()
			}
		}

		stack.Push(num, i)

	}
	for i := 0; i < len(result); i++ {
		fmt.Fprintf(writer, "%d ", result[i])
	}
}
