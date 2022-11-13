package main

import "fmt"

func main(){
    var h, m, c int 
    fmt.Scanf("%d %d", &h, &m)
    fmt.Scanf("%d",&c)
    h = h+(c/60)
    m = m+(c%60)
    if m>=60{
        m = m-60
        h ++
    }
    if h>=24{
        h = h-24
    }
    fmt.Println(h, m)
    
}