package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dequeue struct {
	Nodes []int
}

func (q *Dequeue) PushBack(n int) {
	q.Nodes = append(q.Nodes, n)
}

func (q *Dequeue) PushFront(n int) {
	q.Nodes = append([]int{n}, q.Nodes...)

}

func (q *Dequeue) PopFront() (ret int) {
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

func (q *Dequeue) PopBack() (ret int) {
	ret = -1
	if q.Size() > 1 {
		ret = q.Nodes[len(q.Nodes)-1]
		q.Nodes = q.Nodes[:len(q.Nodes)-1]
	} else if q.Size() == 1 {
		ret = q.Nodes[len(q.Nodes)-1]
		q.Nodes = []int{}
	}

	return ret
}

func (q *Dequeue) Size() int {
	return len(q.Nodes)

}
func (q *Dequeue) Empty() int {
	if q.Size() > 0 {
		return 0
	}
	return 1

}
func (q *Dequeue) Front() int {
	if q.Size() > 0 {
		return q.Nodes[0]
	}
	return -1
}

func (q *Dequeue) Back() int {
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
	var dequeue = new(Dequeue)
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
		case "push_front":
			dequeue.PushFront(number)
		case "push_back":
			dequeue.PushBack(number)
		case "pop_front":
			fmt.Fprintf(writer, "%d\n", dequeue.PopFront())
		case "pop_back":
			fmt.Fprintf(writer, "%d\n", dequeue.PopBack())
		case "size":
			fmt.Fprintf(writer, "%d\n", dequeue.Size())
		case "empty":
			fmt.Fprintf(writer, "%d\n", dequeue.Empty())
		case "front":
			fmt.Fprintf(writer, "%d\n", dequeue.Front())
		case "back":
			fmt.Fprintf(writer, "%d\n", dequeue.Back())
		}
	}

}
