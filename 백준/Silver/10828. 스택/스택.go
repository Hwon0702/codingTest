package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	nodes []int
}

func (s *Stack) Push(n int) {
	s.nodes = append(s.nodes, n)
}
func (s *Stack) Pop() (ret int) {
	if len(s.nodes) < 1 {
		ret = -1
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

		var funcstr string
		var num int
		funcstr, _ = reader.ReadString('\n')

		funcstr = strings.TrimSpace(funcstr)
		strArr := strings.Split(funcstr, " ")
		if strArr[0] == "push" {
			num, _ = strconv.Atoi(strArr[1])
		}

		switch strArr[0] {
		case "push":
			S.Push(num)
		case "pop":
			fmt.Fprintf(writer, "%d\n", S.Pop())
		case "size":
			fmt.Fprintf(writer, "%d\n", S.Size())
		case "empty":
			fmt.Fprintf(writer, "%d\n", S.Empty())
		case "top":
			fmt.Fprintf(writer, "%d\n", S.Top())
		}
	}
	writer.Flush()
}
