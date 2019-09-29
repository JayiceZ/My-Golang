package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Email struct{
	Where string `xml:"where,attr"`  //这样就是说写成  <Email where="work">  的形式
	Addr string
}

type Address struct{
	City string
	Area int
}

type Student struct{
	Id int `xml:"id,attr"`  //<Student id="10">
	Address
	Email []Email
	FirstName string `xml:"name>firstNme"`  //><name><firstNme>jay</firstNme><lastName>chou</lastName></name>
	LastName string `xml:"name>lastName"`
}

func main(){
	stu:=Student{10,Address{"GZ",20},[]Email{Email{"home","118@qq.com"},Email{"work","z@mintegral.com"}},"jay","chou"}
	file,err:=os.Create("d:/stu.xml")
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer file.Close()
	encoder:=xml.NewEncoder(file)
	err2:=encoder.Encode(stu)
	if err2!=nil{
		fmt.Println(err2.Error())
	}

	file.Seek(0, os.SEEK_SET)
	var stu2 Student
	decoder:=xml.NewDecoder(file)
	err3:=decoder.Decode(&stu2)
	if err3!=nil{
		fmt.Println(err3.Error())
	}
	fmt.Println(stu2) //{10 {GZ 20} [{home 118@qq.com} {work z@mintegral.com}] jay chou}
}
