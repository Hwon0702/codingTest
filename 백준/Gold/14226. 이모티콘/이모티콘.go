package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	Emoji     []int
	Clipboard []int
}

func (q *Queue) Enqueue(e, c int) {
	q.Emoji = append(q.Emoji, e)
	q.Clipboard = append(q.Clipboard, c)
}
func (q *Queue) Dequeue() (e, c int) {
	if len(q.Emoji) <= 0 || len(q.Clipboard) <= 0 {
		e = -1
		c = -1
	} else {
		e = q.Emoji[0]
		c = q.Clipboard[0]
		if len(q.Emoji) == 1 {
			q.Emoji = []int{}
			q.Clipboard = []int{}
		} else {
			q.Emoji = q.Emoji[1:]
			q.Clipboard = q.Clipboard[1:]
		}
	}
	return e, c
}
func (q *Queue) Clear() {
	q.Emoji = []int{}
	q.Clipboard = []int{}
}
func (q *Queue) Size() int {
	return len(q.Clipboard)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, emojiCnt, clibboard int
	fmt.Fscanf(reader, "%d\n", &n)
	var q = new(Queue)
	var already = make(map[string]int)
	q.Enqueue(1, 0)
	already[fmt.Sprintf("%d,%d", 1, 0)] = 0
	for q.Size() > 0 {
		emojiCnt, clibboard = q.Dequeue()
		if emojiCnt == n {
			fmt.Println(already[fmt.Sprintf("%d,%d", emojiCnt, clibboard)])
			q.Clear()
			break
		} else {
			//화면에있는거 복사
			if _, exists := already[fmt.Sprintf("%d,%d", emojiCnt, emojiCnt)]; !exists {
				already[fmt.Sprintf("%d,%d", emojiCnt, emojiCnt)] = already[fmt.Sprintf("%d,%d", emojiCnt, clibboard)] + 1
				q.Enqueue(emojiCnt, emojiCnt)
			}
			//클립보드 붙여넣기
			if _, exists := already[fmt.Sprintf("%d,%d", emojiCnt+clibboard, clibboard)]; !exists {
				already[fmt.Sprintf("%d,%d", emojiCnt+clibboard, clibboard)] = already[fmt.Sprintf("%d,%d", emojiCnt, clibboard)] + 1
				q.Enqueue(emojiCnt+clibboard, clibboard)
			}
			//화면에있는거 지우기
			if _, exists := already[fmt.Sprintf("%d,%d", emojiCnt-1, clibboard)]; !exists {
				already[fmt.Sprintf("%d,%d", emojiCnt-1, clibboard)] = already[fmt.Sprintf("%d,%d", emojiCnt, clibboard)] + 1
				q.Enqueue(emojiCnt-1, clibboard)
			}
		}
	}
}
