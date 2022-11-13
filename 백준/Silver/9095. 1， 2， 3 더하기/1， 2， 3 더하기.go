package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
https://www.acmicpc.net/problem/9095
*/
var SumToVal = map[int]int{
	1: 1,
	2: 2,
	3: 4,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int

	fmt.Fscanf(reader, "%d\n", &tc)

	for i := 0; i < tc; i++ {
		var num int
		fmt.Fscanf(reader, "%d\n", &num)
		fmt.Fprintf(writer, "%d\n", Calculate(num))
	}
}

func Calculate(num int) int {

	if _, ok := SumToVal[num]; ok {
		return SumToVal[num]
	} else {
		for idx := 1; idx <= num; idx++ {
			if _, ok := SumToVal[idx]; !ok { //저장돼있는 값이 없다면
				SumToVal[idx] = SumToVal[idx-1] + SumToVal[idx-2] + SumToVal[idx-3]
			}
		}
	}
	return SumToVal[num]
}
