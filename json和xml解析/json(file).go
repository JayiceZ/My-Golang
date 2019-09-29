package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	file,err:=os.Create("d:/try.txt")
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer file.Close()

	encoder:=json.NewEncoder(file)
	encoder.Encode(stu)

	file.Seek(0,os.SEEK_SET)
	var stu2 Student
	decoder:=json.NewDecoder(file)
	decoder.Decode(&stu2)
	fmt.Println(stu2)  //{10 {GZ 20} [{home 118@qq.com} {work z@mintegral.com}] jay chou}


}
