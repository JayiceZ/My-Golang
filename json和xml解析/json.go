package main

import (
	"encoding/json"
	"fmt"
)

type Email struct{
	Where string `json:"where"`
	Addr string   `json:"addr"`
}

type Address struct{
	City string  `json:"city"`
	Area int   `json:"area"`
}

type Student struct{
	Id int   `json:"id"`
	Address   `"json:address"`
	Email []Email    `json:"email"`
	FirstName string   `json:"firstname"`
	LastName string    `json:"lastname"`
}

func main(){
	stu:=Student{10,Address{"GZ",20},[]Email{Email{"home","118@qq.com"},Email{"work","z@mintegral.com"}},"jay","chou"}
	fmt.Println(stu)  //{10 {GZ 20} [{home 118@qq.com} {work z@mintegral.com}] jay chou}
	buf,err:=json.Marshal(stu)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf))     //{"id":10,"city":"GZ","area":20,"email":[{"where":"home","addr":"118@qq.com"},{"where":"work","addr":"z@mintegral.com"}],"firstname":"jay","lastname":"chou"}

	var stu2 Student
	err2:=json.Unmarshal(buf,&stu2)
	if err2!=nil{
		fmt.Println(err2.Error())
	}
	fmt.Println(stu2)   //{10 {GZ 20} [{home 118@qq.com} {work z@mintegral.com}] jay chou}
}
