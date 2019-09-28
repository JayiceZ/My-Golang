package main

import (
	"fmt"
)

func main(){
	//Array 是值类型，Slice 和 Map 是引用类型。他们是有很大区别的，尤其是在参数传递的时候。
	//另外，Slice 和 Map 的变量 仅仅声明是不行的，必须还要分配空间（也就是初始化，initialization） 才可以使用。
	//第三，Slice 和 Map 这些引用变量 的 内存分配，不需要你操心，因为 golang 是存在 gc 机制的（垃圾回收机制）


	//array 数组
		//声明
	var a [10]int
	fmt.Println(a)   //[0 0 0 0 0 0 0 0 0 0]

	var b=[5]int{1,2,3,4,5}
	fmt.Println(b)  //[1 2 3 4 5]

	var c=[...]int{1,2,3,4,5,6,7,8,9}  //数组长度由初始化列表来确定。若不写...  则成为了切片
	fmt.Println(c)  //[1 2 3 4 5 6 7 8 9]

	//通过下表来初始化
	var d=[10]int{1:1,2:2}
	fmt.Println(d)   //[0 1 2 0 0 0 0 0 0 0]

	var array=[5]int{1,2,3,4,5}
	fmt.Println(len(array))  //5
	//遍历
	for x:=0;x< len(array);x++{
		fmt.Print(array[x])  //12345
	}




	//slice 切片

		//直接声明
	s1:=[]int{1,2,3,4,5}
	s2:=s1
	s2[0]=100
	//因为是引用传递，传递的是地址，直接会对原数据造成影响
	fmt.Println(s1)  //[100 2 3 4 5]
	fmt.Println(s2)  //[100 2 3 4 5]

		//根据底层数据进行创建
	f:=[5]int{1,2,3,4,5}
	//包左不包右
	s3:=f[0:4]
	fmt.Println(s3)  //[1 2 3 4]

		//通过make来构建
	s4:=make([]int,5)
	s5:=make([]int,3,4)
	fmt.Println(s4)  //[0 0 0 0 0]
	fmt.Println(len(s5), cap(s5))  //3 4


	//对slice进行操作
	//添加
	b1:=[]int{1,2,3,4,5}
	b2:=append(b1,1,2)
	fmt.Println(b2)   //[1 2 3 4 5 1 2]
	fmt.Println(b1)   //[1 2 3 4 5]   append后不会影响原切片
	b2[0]=100;
	fmt.Println(b1)  //[1 2 3 4 5]
	fmt.Println(b2)  //[100 2 3 4 5 1 2]


	//copy  返回进行复制的数，为两个切片长度的最小值
	c1:=[]int{1,2,3,4,5,6,7,8}
	c2:=make([]int,3,6)
	num:=copy(c2,c1)
	fmt.Println(num)   //3
	fmt.Println(c1)     //[1 2 3 4 5 6 7 8]
	fmt.Println(c2)    //[1 2 3]
	c3:=make([]int,10,11)
	num1:=copy(c3,c1)
	fmt.Println(num1)   //  8
	fmt.Println(c3)   //[1 2 3 4 5 6 7 8 0 0]



	//map
	//声明之后的map必须初始化或创建才能使用的，否则就是nil


	//map的创建
	var myMap=map[string]string{}
	myMap["jay"]="ice"

	m1:=make(map[string]string)
	m1["ice"]="jay"

	//查找
	v,ok:=myMap["jay"]
	fmt.Println(v)  //ice
	fmt.Println(ok)  //true

	v1,ok1:=myMap["haha1"]
	fmt.Println(v1)
	fmt.Println(ok1)  //false

	//删除
	delete(myMap,"jay")
	v2,ok2:=myMap["jay"]
	fmt.Println(v2)
	fmt.Println(ok2)  //false

}
