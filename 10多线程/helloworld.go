package main

import "fmt"

func hello(c chan<-int){
	defer close(c)
	for i:=0;i<10;i++{
		//fmt.Println("hello")
		c<-i
		fmt.Println("hello",i)
	}
}

func world(c <-chan int,f chan<-int){
	for{
		if v,ok:=<-c;ok{
			fmt.Println("world",v)
		}else{
				break
		}
	}

	f<-1
}

func main(){
	c:=make(chan int)
	d:=make(chan int)

	go hello(c)
	go world(c,d)

	<-d

}
