package main

import "fmt"

func divide(x,y int) int{
	return x/y
}

func main(){
	//recover作用相当于java中的catch，用于捕获异常，并且要放在被defer修饰的方法中，defer即为压进栈中稍后运行，当出现错误时，就会出栈操作捕获函数
	defer func(){
		if err:=recover();err!=nil {
			fmt.Println(err)
		}
	}()

	divide(5,1)
	fmt.Println("successful run")
}
