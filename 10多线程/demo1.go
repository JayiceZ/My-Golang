package main

import "fmt"

func main(){
	ch:=make(chan int,10)
	for i:=0;i<10;i++{
		go func (){
			var a int
			for j:=0;j<10000;j++{
				a++
			}
			ch<-a
		}()
	}

	var sum int
	func (){
		for i:=0;i<10;i++{
			sum+=<-ch
		}
	}()

	fmt.Println(sum)
}
