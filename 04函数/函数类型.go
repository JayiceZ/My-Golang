package main

import "fmt"

type fun func(int)bool

func isBiggerThan5(x int) bool{
	if x<=5{
		return false
	}
	return true
}

func play(array []int,f fun)  {
	for num:=range array{
		if f(num){
			fmt.Println(num,"isBiggerThan5")
		}
	}
}

func main(){
	array:=[]int{1,2,3,4,5,6,7,8,9}
	play(array,isBiggerThan5)
	//6 isBiggerThan5
	//7 isBiggerThan5
	//8 isBiggerThan5
}
