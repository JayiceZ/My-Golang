package main

import "fmt"


//当管道为满时，被堵塞，不能再进行生产。也就是不能执行c<-i
func producer(c chan<-int){
	defer close(c)
	for i:=0;i<10;i++{
			c<-i
	}
}

//当管道为空时，不能进行读操作，也就是不能执行<-c
func comsumer(c <-chan int,f chan<-int){
	for{
		if v,ok:=<-c;ok{
			fmt.Println(v)
		}else{
			break
		}
	}
	f<-1  //当执行完毕时，把1写到f中
}

func main(){
	c:=make(chan int)
	f:=make(chan int)

	go producer(c)
	go comsumer(c,f)
	//等f中有数据了，主线程才会结束
	<-f
}
