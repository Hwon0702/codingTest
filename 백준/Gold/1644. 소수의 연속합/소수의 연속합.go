package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var target int
	fmt.Fscanf(reader, "%d\n", &target)
	var primeNumber = make([]bool, 4000001)
	var primeNumbers = []int{}
	var cnt, sum, end int
	for i := 0; i < 4000001; i++ {
		primeNumber[i] = true
	}
	primeNumber[0] = false
	primeNumber[1] = false
	for i := 2; i < 2000; i++ {
		if primeNumber[i] == false {
			continue
		}
		var divCnt int
		for j := 1; j <= 2000; j++ {
			if i%j == 0 {
				divCnt++
			}
			if divCnt > 1 {
				for k := 2; k*i < 4000001; k++ {
					primeNumber[k*i] = false
				}
				break
			}
		}
	}
	for i, v := range primeNumber {
		if v {
			primeNumbers = append(primeNumbers, i)
		}
	}
	cnt = 0
	sum = 0
	end = 0
	for start := 0; start < len(primeNumbers); start++ {
		sum = 0
		end = start
		for sum < target && end < len(primeNumbers) {
			sum += primeNumbers[end]
			end += 1
		}
		if sum == target {
			cnt += 1
		}
		sum -= primeNumbers[start]
	}
	fmt.Println(cnt)
}
