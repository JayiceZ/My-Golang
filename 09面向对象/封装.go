package main

import (
	"fmt"
)

type student struct{
	id int
}



func (stu *student) getId() int{
	return stu.id
}

func (stu *student) setId(i int){
	stu.id=i
}

func main(){
	stu:=student{}
	stu.setId(10)
	fmt.Println(stu.getId())  //10
}
