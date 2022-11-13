package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	2002
*/
func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	inputNumToCar := make(map[string]int)
	outputNumToCar := make([]int, n)
	for i := 0; i < n; i++ {
		var carNum string
		fmt.Fscanf(reader, "%s\n", &carNum)
		inputNumToCar[carNum] = i
	}
	for i := 0; i < n; i++ {
		var carNum string
		fmt.Fscanf(reader, "%s\n", &carNum)
		outputNumToCar[i] = inputNumToCar[carNum]
	}
	/*
		모든 input에 대하여 나보다 뒤에 들어온게 나보다 먼저 나가면 추월
		나보다 뒤에 들어온거 -> n+1
		나보다 앞에 있다 -> n-1
	*/
	var pass int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if outputNumToCar[i] > outputNumToCar[j] {
				pass++
				break
			}
		}
	}
	fmt.Fprintf(writer, "%d", pass)
}
