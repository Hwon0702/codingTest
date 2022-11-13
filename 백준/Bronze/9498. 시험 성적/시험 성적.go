package main

import "fmt"

func main(){
    var record int
    fmt.Scanf("%d", &record)
    
    if record<60{
        fmt.Println("F")
    }else if record<70{
        fmt.Println("D")
    }else if record<80{
        fmt.Println("C")
    }else if record<90{
        fmt.Println("B")
    }else {
        fmt.Println("A")
    }
}