package main

import (
	"encoding/xml"
	"fmt"
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
	fmt.Println(stu)  //{10 {GZ 20} [{home 118@qq.com} {work z@mintegral.com}] jay chou}
	buf,err:=xml.Marshal(stu)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(buf))  //<Student id="10"><City>GZ</City><Area>20</Area><Email where="home"><Addr>118@qq.com</Addr></Email><Email where="work"><Addr>z@mintegral.com</Addr></Email><name><firstNme>jay</firstNme><lastName>chou</lastName></name></Student>

	var stu2 Student
	err1:=xml.Unmarshal(buf,&stu2)
	if err1!=nil{
		fmt.Println(err1.Error())
	}
	fmt.Println(stu2)  //{10 {GZ 20} [{home 118@qq.com} {work z@mintegral.com}] jay chou}
}
