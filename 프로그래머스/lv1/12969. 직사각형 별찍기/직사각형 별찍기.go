package main

import "fmt"
func solution(a, b int) {
	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    solution(a, b)
}