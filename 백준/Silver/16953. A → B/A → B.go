package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mod(n int) int {
	return n / 2
}

func minus(n int) (ret int) {
	num := strconv.Itoa(n)
	numArr := strings.Split(num, "")
	if numArr[len(numArr)-1] == "1" {
		numArr = numArr[:len(numArr)-1] //1을 더함

		ret, _ = strconv.Atoi(strings.Join(numArr, ""))
		return ret
	} else {
		return 0
	}

}

/*
끝이 2의 배수 -> 마지막 연산이 *2
끝이 2의 배수가 아님 -> 마지막 연산은 +1
*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var start, result, count int
	count = 1
	fmt.Fscanf(reader, "%d %d\n", &start, &result)
	for result == start || result > 0 {
		count++
		if result%2 == 0 {
			result = mod(result)
		} else {
			result = minus(result)
		}
		if result == start {
			break
		}
	}
	if result == 0 {
		count = -1
	}
	fmt.Println(count)
}
