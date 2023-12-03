package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type Stack struct {
	Numbers []int
}

var stack *Stack

func Init() *Stack {
	stack = new(Stack)
	return stack
}

func (s *Stack) Push(n int) {
	s.Numbers = append(s.Numbers, n)
}

func (s *Stack) Pop() (n int) {
	n = -1
	if len(s.Numbers) <= 0 {
		return
	}
	n = s.Numbers[len(s.Numbers)-1]
	if len(s.Numbers) <= 1 {
		s.Numbers = []int{}
	} else {
		s.Numbers = s.Numbers[:len(s.Numbers)-1]
	}
	return
}

func (s *Stack) Len() (n int) {
	n = len(s.Numbers)
	return
}

func (s *Stack) Empty() int {
	if len(s.Numbers) > 0 {
		return 0
	}
	return 1
}

func (s *Stack) Next() int {
	if len(s.Numbers) > 0 {
		return s.Numbers[len(s.Numbers)-1]
	}
	return -1
}

func Calc(s *Stack, method string, num ...int) {
	if method == "1" {
		s.Push(num[0])
	} else if method == "2" {
		fmt.Fprintf(writer, "%d\n", s.Pop())

	} else if method == "3" {
		fmt.Fprintf(writer, "%d\n", s.Len())

	} else if method == "4" {
		fmt.Fprintf(writer, "%d\n", s.Empty())

	} else if method == "5" {
		fmt.Fprintf(writer, "%d\n", s.Next())
	}
}

func main() {
	var n, num int
	defer writer.Flush()
	s := Init()
	fmt.Fscanf(reader, "%d\n", &n)

	for i := 0; i < n; i++ {
		m, _ := reader.ReadString('\n')
		m = strings.TrimSpace(m)
		methods := strings.Split(m, " ")
		if len(methods) > 1 {
			num, _ = strconv.Atoi(methods[1])
			Calc(s, methods[0], num)
		} else {
			Calc(s, methods[0])
		}
	}
}
