package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var reservation = make([][]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		reservation[i] = []int{a, b}
	}
	//이용할 수 있는 최대 개수
	//종료되자마자 시작하는 경우->종료되지 않으면 시작할 수 없음
	//따라서 종료를 기준으로 문제 세워야 함
	sort.Slice(reservation, func(a, b int) bool {
		if reservation[a][1] <= reservation[b][1] {
			if reservation[a][1] == reservation[b][1] {
				return reservation[a][0] < reservation[b][0]
			} else {
				return reservation[a][1] < reservation[b][1]
			}
		} else {
			return false
		}
	})
	var cnt = 1
	var time = reservation[0][1]
	for i := 1; i < n; i++ {
		if time <= reservation[i][0] {
			cnt++
			time = reservation[i][1]
		}
	}
	fmt.Fprintf(writer, "%d\n", cnt)
}
