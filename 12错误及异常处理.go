package main

import (
	"errors"
	"fmt"
)

func divide(x,y int) (int,error){
	if y<0{
		return 0,errors.New("除数不能小于0")   //返回错误类型
	}else if y==0{
		panic("被除数不能为0")    //直接抛出异常，这个异常需要被捕获，否则会报错
	}else{
		return x/y,nil
	}
}

func main(){
	defer func(){    //用来捕获异常
		if err:=recover();err!=nil{
			fmt.Println("err:",err)   //打印异常
		}
	}()

	if num,err:=divide(5,-1);err!=nil{
		fmt.Println(err)   //打印错误
	}else{
		fmt.Println(num)
	}

	fmt.Println("i am here")

	if num,err:=divide(5,0);err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(num)
	}

	fmt.Println("still main")   //错误出现后，直接执行defer，不会再执行这里的语句
	
	//除数不能小于0
	//i am here
	//err: 被除数不能为0
}
