package main

import "fmt"

func add(x,y int) int{
	return x+y
}

//多返回值
func double(x,y int) (int,int){
	return x,y
}

//命名返回值
func circle(r int)(area,len float32){
	area = float32(r) * float32(r) * 3.14
	len = float32(r)*2*3.14
	return   //自动把area和len填充到返回值中
}
//多参数
func addTotal(args...int) int{
	s:=0
	for num:=range args{
		s+=num
	}
	return s
}
func main(){
	fmt.Println(add(10,20))  //30
	//多返回值
	x,y:=double(10,20)
	fmt.Println(x,y)  //10 20
	//命名返回值
	myArea,myLen:=circle(2)
	fmt.Println(myArea,myLen)  //12.56 12.56

	//空白符
	//比如说此时只需要获得面积，不需要获得周长，便可以用空白符来接收周长
	_,len:=circle(2)
	fmt.Println(len)  //12.56
	
	fmt.Println(addTotal(1,2,3,4,5))
	fmt.Println(1,2)
}
