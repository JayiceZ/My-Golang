package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	str:="i am jayice "
	err:=ioutil.WriteFile("d:/test2.txt",[]byte(str),0666)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("write succesful")
	}

	buf1,err1:=ioutil.ReadFile("d:/test2.txt")
	if err1!=nil{
		fmt.Println("err")
	}else{
		fmt.Println(string(buf1))
	}

	//通过ReadAll方法来读
	file,err2:=os.OpenFile("d:/test2.txt", os.O_RDONLY, 0666)
	if err2!=nil{
		fmt.Println(err2)
	}
	defer file.Close()
	buf2,err3:=ioutil.ReadAll(file)
	if err2!=nil{
		fmt.Println(err3)
	}else{
		fmt.Println(string(buf2))
	}

}
