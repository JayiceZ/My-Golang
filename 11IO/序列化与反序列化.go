package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Student struct{
	//发现了一个神坑。这里的参数命名必须首字母大写，否则序列化失败
	Name string
	Age int
}

func main(){
	stu:=&Student{"jay",10}
	f,err:=os.Create("d:/test5.txt")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	encode:=gob.NewEncoder(f)
	//序列化
	encode.Encode(stu)

	f.Seek(0, os.SEEK_SET)
	decoder:=gob.NewDecoder(f)
	var stu2 Student
	//反序列化
	decoder.Decode(&stu2)
	fmt.Println(stu2)
}
