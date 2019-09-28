package main

import "fmt"

type student struct{
	id int
	name string
	class string
}

type Astudent struct{
	stu student
	age int
}

func (stu *student) getId() int{
	return stu.id
}

func (stu *student) setId(i int){
	stu.id=i
}

func main(){
	astu:=Astudent{}
	astu.stu.id=1
	fmt.Println(astu.stu.getId())   //1
}
