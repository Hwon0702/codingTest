package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func getGcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	for a > 0 {
		a, b = b%a, a
	}
	return b
}
func makeUnique(arr []int) []int {
	var numToMap = make(map[int]bool)
	var uniqueResults = []int{}
	for _, num := range arr {
		if _, exists := numToMap[num]; !exists {
			numToMap[num] = true
			uniqueResults = append(uniqueResults, num)
		}
	}
	return uniqueResults
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var numbers = make([]int, n)
	var diff = []int{}
	var gcd = 0
	var results = []int{}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	sort.Ints(numbers)
	for i := 1; i < n; i++ {
		diff = append(diff, numbers[i]-numbers[i-1])
	}
	sort.Ints(diff)
	gcd = diff[0]
	for i := 1; i < len(diff); i++ {
		gcd = getGcd(gcd, diff[i])
	}
	for i := 2; i <= int(math.Sqrt(float64(gcd))); i++ {
		if gcd%i == 0 {
			results = append(results, i)
			results = append(results, int(gcd/i))
		}
	}
	results = append(results, gcd)
	uniqueArr := makeUnique(results)
	sort.Ints(uniqueArr)
	for _, v := range uniqueArr {
		if v != 1 {
			fmt.Fprintf(writer, "%d ", v)
		}
	}
}
