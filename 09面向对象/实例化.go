package main

import "fmt"

type student struct{
	id int
	name string
	class string
}

func main(){
	stu:=student{10,"jay","15"}
	stu.id=20
	fmt.Println(stu)  //{20 jay 15}

	stu1:=new(student)
	stu1.id=10
	stu1.name="ice"
	stu1.class="16"
	fmt.Println(stu1)  //&{10 ice 16}
}
