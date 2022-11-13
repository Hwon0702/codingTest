package main

func sum(a []int) int {
	sum := 0
	for _, num := range a {
		sum += num
	}
	return sum
}