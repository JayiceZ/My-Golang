package myMath

import "fmt"

func init(){
	fmt.Println("myMath init")
}

//想被其他包调用，函数名首字母必须大写，为public
func Add(x,y int) int{
	return x+y
}
