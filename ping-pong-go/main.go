package main

import (
	"fmt"
	"time"
	"sync"
)

func Ping(c chan<- string){
	c <- "Ping!"		
}

func Pong(c chan<- string){
	c <- "Pong!"
}

func manageSending(c chan<- string){
	for i:=0;i<3;i++{
		Ping(c)
		Pong(c)
	}
	close(c)
}

func Puts(c <-chan string, wg *sync.WaitGroup){
	defer wg.Done()
	for msg:= range c{
		fmt.Println(msg)
		time.Sleep(350*time.Millisecond)
	}
}

func main(){
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go manageSending(c)
	go Puts(c,&wg)

	var input string
	fmt.Scanln(&input)
}