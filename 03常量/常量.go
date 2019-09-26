package main

import (
	"fmt"
	"math"
)

func main(){
	const a1=10; //相当于java的final修饰变量，就变成了常量
	//a1=20;  //error

	//不允许用常量来接受函数返回值
	var a2=math.Sqrt(4);
	//const a3=math.Sqrt(4);  //error
	fmt.Println(a2);

	//字符串常量
	var b1="sum";
	type myString string;
	var b2 myString ="sum";
	//b1=b2;  myString和string不能互相赋值
	fmt.Println(b1,b2);

	//布尔常量
	const c1=true;
	var c2=c1;
	type myBool bool;
	var c3 myBool =c1;
	//c3=c2;   //false
	fmt.Println(c2,c3);

	//数字常量
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)
	//intVar 5
	//int32Var 5
	//float64Var 5
	//complex64Var (5+0i)

	
}
