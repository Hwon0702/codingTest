package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	Nodes []int
}

func (q *Queue) Push(n int) {
	q.Nodes = append(q.Nodes, n)
}

func (q *Queue) Pop() (ret int) {
	ret = -1
	if q.Size() > 1 {
		ret = q.Nodes[0]
		q.Nodes = q.Nodes[1:]
	} else if q.Size() == 1 {
		ret = q.Nodes[0]
		q.Nodes = []int{}
	}

	return ret
}

func (q *Queue) Size() int {
	return len(q.Nodes)

}
func (q *Queue) Empty() int {
	if q.Size() > 0 {
		return 0
	}
	return 1

}
func (q *Queue) Front() int {
	if q.Size() > 0 {
		return q.Nodes[0]
	}
	return -1
}

func (q *Queue) Back() int {
	if q.Size() > 0 {
		return q.Nodes[len(q.Nodes)-1]
	}
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	var queue = new(Queue)
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		var str string
		var operation string
		var number int
		str, _ = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		arr := strings.Split(str, " ")
		operation = arr[0]
		if len(arr) > 1 {
			number, _ = strconv.Atoi(arr[1])
		}
		switch operation {
		case "push":
			queue.Push(number)
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
