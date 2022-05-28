package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	Node []int
}

/*
func (s *Stack) Reverse() {
	for i, j := 0, len(s.Node)-1; i < j; i, j = i+1, j-1 {
		s.Node[i], s.Node[j] = s.Node[j], s.Node[i]
	}
}
*/
func (s *Stack) DeleteLast() string {
	if len(s.Node) <= 0 {
		return "error"
	} else {
		if len(s.Node) > 1 {
			s.Node = s.Node[:len(s.Node)-1]
		} else {
			s.Node = []int{}
		}
	}
	return ""
}

func (s *Stack) DeleteFront() string {
	if len(s.Node) <= 0 {
		return "error"
	} else {
		if len(s.Node) > 1 {
			s.Node = s.Node[1:]
		} else {
			s.Node = []int{}
		}
	}
	return ""
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		var operation, numbers string
		var num int
		var stack = new(Stack)
		fmt.Fscanf(reader, "%s\n", &operation)
		fmt.Fscanf(reader, "%d\n", &num)
		fmt.Fscanf(reader, "%s\n", &numbers)
		stack.Node = make([]int, num)
		numbers = strings.ReplaceAll(numbers, "[", "")
		numbers = strings.ReplaceAll(numbers, "]", "")
		numArr := strings.Split(numbers, ",")
		for j := 0; j < num; j++ {
			stack.Node[j], _ = strconv.Atoi(numArr[j])
		}
		operations := strings.Split(operation, "")
		var result string
		Reverse := false
		for k := 0; k < len(operations); k++ {

			switch operations[k] {
			case "R":
				if !Reverse {
					Reverse = true
				} else {
					Reverse = false
				}
			case "D":

				if Reverse {
					result = stack.DeleteLast()
					if result != "" {
						break
					}
				} else {
					result = stack.DeleteFront()
					if result != "" {
						break
					}
				}
			}
		}
		if result != "" {
			fmt.Fprintf(writer, "%s\n", result)
		} else {
			fmt.Fprintf(writer, "[")
			if !Reverse {
				for l := 0; l < len(stack.Node); l++ {
					fmt.Fprintf(writer, "%d", stack.Node[l])
					if l != len(stack.Node)-1 {
						fmt.Fprintf(writer, ",")
					}
				}
			} else {
				for l := 1; l <= len(stack.Node); l++ {
					fmt.Fprintf(writer, "%d", stack.Node[len(stack.Node)-l])
					if l != len(stack.Node) {
						fmt.Fprintf(writer, ",")
					}
				}
			}
			fmt.Fprintf(writer, "]\n")
		}

	}

}
