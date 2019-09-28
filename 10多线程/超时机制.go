package main

import (
	"fmt"
	"time"
)

func main(){
	//缓冲channel
	c:=make(chan int,2)
	select {
	case <-c:fmt.Println("有数据")
	case <-time.After(5* time.Second):
		fmt.Println("超时退出")
	}


	fmt.Println("you get here")
}
