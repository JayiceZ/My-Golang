package main

import "fmt"

func main() {
	var age int;      //声明，相当于:int age; 默认值为0
	fmt.Println(age); //0

	age = 10;
	fmt.Println(age);

	//自动判断类型

	var b1 = 10; //自动判断为10;
	fmt.Println(b1);

	//支持多个变量
	var a, b, c int = 1, 2, 3;
	//也可以同时打印
	fmt.Println(a, b, c); //1 2 3

	//组合声明
	var (
		name= "naveen"
		height =10;
	)
	fmt.Println(name, height);  //naveen 10

	//简短声明
	a1:=1;
	fmt.Println(a1);  //1

	a2,a3:=10,"jay";
	fmt.Println(a2,a3);  //10 jay
}
