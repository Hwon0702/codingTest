package main

import "fmt"
import "sort"
func main(){
    var a,b,c int
    fmt.Scanf("%d %d %d", &a, &b, &c)
    if (a==b&& b==c){
        fmt.Println(a*1000 + 10000)
    }else if (a==b&&a!=c|| b==c&&a!=b ||a==c&&a!=b){
        eq :=0
        if a==b || a==c {
            eq = a
        }else if b==c{
            eq =c
        }
        fmt.Println(eq*100+1000)
    }else{
        arr:=[]int{a,b,c}
        sort.Sort(sort.IntSlice(arr))
		fmt.Println(arr[2]*100)
    }
}