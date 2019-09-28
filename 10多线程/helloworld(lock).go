package main

import (
	"fmt"
	"sync"
)

func main(){
	a:=new(sync.Mutex)
	b:=new(sync.Mutex)
	b.Lock()
	go hello(a,b)
	go world(a,b)

	var str string
	fmt.Scan(&str)
}
func hello(a *sync.Mutex,b *sync.Mutex){
	for i:=0;i<10;i++{
		a.Lock()
		fmt.Println("hello")
		b.Unlock()
	}
}

func world(a *sync.Mutex,b *sync.Mutex){
	for i:=0;i<10;i++{
		b.Lock()
		fmt.Println("world")
		a.Unlock()
	}
}
