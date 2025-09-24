package main

import (
	"fmt"
)

func main(){
	for i:=1;i<101;i++{
		if i%3==0 && i%5!=0{
			fmt.Println("Pin")
		}else if i%5==0 && i%3 != 0{
			fmt.Println("Pan")
		}else if i%5==0 && i%3 == 0{
			fmt.Println("PinPan")
		}else{
			fmt.Println(i)
		}
	}
}