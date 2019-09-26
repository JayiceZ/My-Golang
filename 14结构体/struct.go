package main

import "fmt"

//普通结构体
type student struct{
	old int;
	name string;
}

//匿名结构
type teacher struct{
	string;
	int;
}

//嵌套结构体
type Address struct{
	city string;
	state string;
}
type person struct{
	age int;
	name string;
	address Address;
}

//提升字段
type add struct{
	pop int;
	area int;
}

type city struct{
	state string;
	add;
}

func main(){
	s:=student{18,"jay"};
	//匿名结构体
	s2:=struct{old int;name string}{20,"jj"};
	fmt.Println(s);  //{18 jay}
	fmt.Println(s2);  //{20 jj};


	//结构体的0值
	var s3 student;
	fmt.Println(s3);   //{0 }

	s4:=student{name:"jay"};
	fmt.Println(s4);  //{0 jay}


	var s5 student;
	s5.name="jayjay";
	fmt.Println(s5);  //{0 jayjay}


	//结构体指针
	s6:=&student{10,"jay"};
	fmt.Println((*s6).name);  //jay
	fmt.Println(s6.name);  //jay

	//默认结构体的使用
	s7:=teacher{"ll",20};
	s7.string="jay";
	s7.int=10;
	fmt.Println(s7);   //{jay 10}

	//嵌套结构体
	var s8 person;
	s8.name="jay";
	s8.age=20;
	s8.address=Address{"gz","gd"};
	fmt.Println(s8);  //{20 jay {gz gd}}

	//提升结构体
	var s9 city;
	s9.area=10;
	s9.pop=1000;
	s9.state="gz";
	fmt.Println(s9);  //{gz {1000 10}}

	//结构体相等性
	a1:=student{10,"jay"};
	a2:=student{10,"jay"};
	a3:=student{20,"jay"};
	fmt.Println(a1==a2);  //true
	fmt.Println(a2==a3);  //false;
}
