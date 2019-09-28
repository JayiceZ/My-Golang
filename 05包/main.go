package main

import (
	"awesomeProject/src/myMath"   //！注意路径
	"fmt"
	"math"
)
//在go中，不允许import进来的包不被访问
//在共有变量区用空白符来调用一下调进来的包，就可以了

var _=math.Abs(1)
func init(){
	fmt.Println("main init")
}

func main(){
	fmt.Println(myMath.Add(1,2))
	/*先加载包的init，然后加载main的init*/
	//myMath init
	//main init
	//3
}

