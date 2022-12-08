package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n, j int
	var flag bool
	var primes = make([]bool, 1000001)
	for i := 2; i < 1000001; i++ {
		primes[i] = true
	}
	for i := 2; i < 1000001; i++ {
		j = 2
		if primes[i] {
			for i*j < 1000001 {
				primes[i*j] = false
				j++
			}
		}
	}
	for {
		fmt.Fscanf(reader, "%d\n", &n)
		if n == 0 {
			break
		}
		flag = true
		for i := 3; i < n; i++ {
			if primes[i] && primes[n-i] {
				fmt.Fprintf(writer, "%d = %d + %d\n", n, i, n-i)
				flag = false
				break
			}
		}
		if flag {
			fmt.Fprintf(writer, "%s\n", "Goldbach's conjecture is wrong.")
		}
	}
}
