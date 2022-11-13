package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReverseArr(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var a, b string

	fmt.Fscanf(reader, "%s %s\n", &a, &b)
	aArr := strings.Split(a, "")
	bArr := strings.Split(b, "")

	maxLength := 0
	if len(aArr) >= len(bArr) {
		maxLength = len(aArr)
	} else {
		maxLength = len(bArr)
	}
	ret := make([]int, maxLength+1)
	aArr = ReverseArr(aArr)
	bArr = ReverseArr(bArr)
	for i := 0; i < maxLength; i++ {
		var numA, numB int
		if i < len(aArr) {
			numA, _ = strconv.Atoi(aArr[i])
		} else {
			numA = 0
		}
		if i < len(bArr) {
			numB, _ = strconv.Atoi(bArr[i])
		} else {
			numB = 0
		}
		ret[maxLength-i] = ret[maxLength-i] + numA + numB
		if ret[maxLength-i] >= 10 {
			ret[maxLength-i-1] = ret[maxLength-i] / 10
			ret[maxLength-i] = ret[maxLength-i] % 10

		}
	}

	for i := 0; i <= len(ret)-1; i++ {
		if i == 0 && ret[i] > 0 || i != 0 {
			fmt.Print(ret[i])
		}
	}

}
