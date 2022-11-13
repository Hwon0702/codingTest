package main

import "fmt"

func main(){
    var n int
    fmt.Scanf("%d", &n)
	values:=[][]int{}
	for idx:=0; idx<n; idx++{
		
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		values = append(values, []int{a,b})
	}
	for _, value:=range values{
		fmt.Println(value[0]+value[1])
	}
}