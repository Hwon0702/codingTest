package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	Node []int
}

func (q *Queue) Push(n int) {
	q.Node = append(q.Node, n)
}
func (q *Queue) Pop() (val int) {
	if len(q.Node) <= 0 {
		val = -1
	} else {
		val = q.Node[0]
		if len(q.Node) > 1 {
			q.Node = q.Node[1:]

		} else {
			q.Node = []int{}
		}
	}
	return val
}
func (q *Queue) Size() int {
	return len(q.Node)
}
func (q *Queue) Empty() int {
	if len(q.Node) > 0 {
		return 0
	} else {
		return 1
	}
}

func (q *Queue) Front() int {
	if len(q.Node) > 0 {
		return q.Node[0]
	} else {
		return -1
	}
}

func (q *Queue) Back() int {
	if len(q.Node) > 0 {
		return q.Node[len(q.Node)-1]
	} else {
		return -1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	var queue = new(Queue)
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)
		arr := strings.Split(op, " ")
		var num int
		if len(arr) > 1 {
			num, _ = strconv.Atoi(arr[1])
		}
		switch arr[0] {
		case "push":
			queue.Push(num)
		case "pop":
			fmt.Fprintf(writer, "%d\n", queue.Pop())
		case "size":
			fmt.Fprintf(writer, "%d\n", queue.Size())
		case "empty":
			fmt.Fprintf(writer, "%d\n", queue.Empty())
		case "front":
			fmt.Fprintf(writer, "%d\n", queue.Front())
		case "back":
			fmt.Fprintf(writer, "%d\n", queue.Back())
		}
	}
}
