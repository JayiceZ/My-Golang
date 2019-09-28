package main

import (
	"fmt"
)

type student interface {
	shout()
}

type astudent struct{
	id int
	name string
}

type bstudent struct{
	id int
	 name string
}

func (atu *astudent) shout(){
	fmt.Println("astudnet is shouting")
}

func (bstu *bstudent) shout(){
	fmt.Println("bstudent is shouting")
}

func main(){
	var astu student=&astudent{}
	astu.shout()  //astudnet is shouting


	var bstu student=&bstudent{}
	bstu.shout()  //bstudnet is shouting

	/*只要是实现了接口中的一系列方法，都是该接口的子类*/
}
