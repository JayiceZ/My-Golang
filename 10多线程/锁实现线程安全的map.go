package main

import (
	"errors"
	"fmt"
	"sync"
)

type myMap struct{
	mp map[string]int
	lock sync.Locker
}

func (this *myMap) Set(key string,val int){
	this.lock.Lock()
	this.mp[key]=val
	this.lock.Unlock()
}

func (this *myMap) get(key string) (int,error){
	this.lock.Lock()
	i,ok:=this.mp[key]
	this.lock.Lock()
	if !ok{
		return i,errors.New("不存在")
	}else{
		return i,nil
	}
}

func(this *myMap)DisPlay(){
	this.lock.Lock()
	defer this.lock.Unlock()
	for key,val:=range this.mp{
		fmt.Println(key,"=",val)
	}
}

func setValue(this *myMap){
	var a rune
	a='a'
	for i:=0;i<10;i++{
		this.Set(string(a+rune(i)),i)
	}
}

func main(){
	m:=&myMap{map[string]int{},new(sync.Mutex)}
	go setValue(m)
	go m.DisPlay()

	var str string
	fmt.Scan(&str)
}
