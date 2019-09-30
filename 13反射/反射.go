package main

import (
	"fmt"
	"reflect"
)

type Address struct{
  	City string
  	Area string
}

type Student struct{
	Address
	Name string
	Age int
}

func (this Student) Say(){
	fmt.Println("hello , i am",this.Name,"i am ",this.Age)
}

func (this Student)Hello(world string){
	fmt.Println("hello",world)
}

func StructInfo(o interface{}){
	t:=reflect.TypeOf(o)
	fmt.Println(t.Name(),"object type:",t.Name())

	if k:=t.Kind();k!=reflect.Struct{
		fmt.Println("the object is not a struct, but it is",k)
		return
	}

	//获取对象的值
	v:=reflect.ValueOf(o)
	fmt.Println(v)

	//获取对象的各个字段
	for i:=0;i<t.NumField();i++{
		f:=t.Field(i)
		val:=v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v \n", f.Name, f.Type, val)

		//
		t1:=reflect.ValueOf(val)
		if k:=t1.Kind();k==reflect.Struct{
			StructInfo(val)
		}
	}

	//获取对象方法
	 for i:=0;i<t.NumMethod();i++{
	 	m:=t.Method(i)
	 	fmt.Printf("%10s:%v \n", m.Name, m.Type)
	 }
}

func Set(o interface{}){
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet(){
		fmt.Println("修改失败")
		return
	}
	v = v.Elem()
	//获取字段
	f := v.FieldByName("Name")
	if !f.IsValid(){
		fmt.Println("修改失败")
		return
	}
	//设置值
	if f.Kind() == reflect.String{
		f.SetString("chairis")
	}
}

func RunMethod(o interface{}){
	v:=reflect.ValueOf(o)

	m1:=v.MethodByName("Say")
	//无参
	m1.Call([]reflect.Value{})

	m2:=v.MethodByName("Hello")
	//有参
	m2.Call([]reflect.Value{reflect.ValueOf("world")})
}

func main(){
	stu:=Student{Address:Address{City:"Shanghai", Area:"Pudong"}, Name:"chain", Age:23}
	fmt.Println(stu)
	StructInfo(stu)

	//Set(stu)
	fmt.Println(stu)

	RunMethod(stu)
}
