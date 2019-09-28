package main

import "fmt"

type student struct{
	id int
	name string
	class string
}

func (stu *student) getId() int{
	return stu.id
}

func (stu *student) setId(i int){
	stu.id=i
}

func main(){
	stu:=student{}
	stu.id=10
	stu.name="jay"
	stu.class="15"
	fmt.Println(stu.getId())

	stu.setId(20)
	fmt.Println(stu.getId())
}
