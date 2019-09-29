package main

import (
	"fmt"
	"os"
)

func  main(){
	//打开文件，若没有就创建
	file,err:=os.OpenFile("d:/test.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	//传入字符串
	file.WriteString("i am jayice")
	
	buf:=make([]byte,1024)
	var str string
	file.Seek(0, os.SEEK_SET)   //重置文件指针，使其指向头部
	for{
		n,ferr:=file.Read(buf)
		if ferr!=nil{
			fmt.Println(ferr)
			break
		}
		//n就是写入了多少个字符，这一点和java类似
		if n==0{
			break
		}
		//覆盖前n个，所以要取0:n
		str+=string(buf[0:n])
	}

	fmt.Println(str)
}
