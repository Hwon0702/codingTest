package main

import "fmt"

func main(){
    var h,m int
    fmt.Scanf("%d %d", &h, &m)
	if (m-45)<0{
		if h-1<0{
			fmt.Println(23 , 60-(45-m))
		}else{
			fmt.Println(h-1, 60-(45-m))
		}

	}else{
		fmt.Println(h, m-45)
	}
   
}