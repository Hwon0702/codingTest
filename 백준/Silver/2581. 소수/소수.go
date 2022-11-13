package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &m)
	fmt.Fscanln(reader, &n)
	var sum int
	var min = n //일단 최소 소수는 들어온 수의 최솟값으로 잡음
	for i := m; i <= n; i++ {
		var divisorCount int
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				divisorCount++ //이렇게 하면 1과 자기 자신이 들어오면 소수
			}

		}
		if divisorCount == 2 {
			sum += i
			if min > i {
				min = i
			}
		}
	}
	if sum == 0 {
		fmt.Fprintf(writer, "%d\n", -1)
	} else {
		fmt.Fprintf(writer, "%d\n%d", sum, min)
	}
}