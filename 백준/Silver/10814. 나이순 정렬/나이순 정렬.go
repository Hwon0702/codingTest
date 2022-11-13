package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type User struct {
	age  int
	name string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanln(reader, &n)

	defer writer.Flush()
	userArr := make([]User, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &userArr[i].age, &userArr[i].name)
	}
	sort.SliceStable(userArr, func(i, j int) bool {
		return userArr[i].age < userArr[j].age
	})
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, userArr[i].age, userArr[i].name)
	}
}
