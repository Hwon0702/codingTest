package main

import "fmt"

func main(){
    var n int
    fmt.Scanf("%d", &n)
    answer:=0
    for i:=1; i<=n; i++{
        answer+=i
    }
    fmt.Println(answer)
}