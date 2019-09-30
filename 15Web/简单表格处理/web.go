package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type mux struct {
}

//ServeHTTP的函数名不能变
//路由
func (m *mux) ServeHTTP	(w http.ResponseWriter,r *http.Request){
	if r.URL.Path=="/first"{
		First(w,r)
		return
	}else if r.URL.Path=="/second"{
		Second(w,r)
		return
	}else if r.URL.Path=="/login"{
		Login(w,r)
		return
	}
		http.NotFound(w,r)
	return
}

func Login(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	m:=r.Method
	if m=="GET"{
		//载入界面
		t,_:=template.ParseFiles("./src/view/login.html")
		t.Execute(w,nil)
	}else if m=="POST"{
		fmt.Println(r.Form["username"])
		fmt.Println(r.Form["password"])  //[123]
		fmt.Println(r.Form.Get("password"))  //123
	}
}

func First(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.URL)  ///first?id=10&name=jay
	for k,v:=range r.Form{
		fmt.Println(k,v)  //id [10]
						 //name [jay]
	}
	fmt.Fprintf(w,"i am first")
}


func Second(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.URL)
	for k,v:=range r.Form{
		fmt.Println(k,v)
	}
	fmt.Fprintf(w,"i am second")
}


func Start(){
	m:=&mux{}
	err:=http.ListenAndServe(":9090",m)
	if err!=nil{
		log.Fatal(err.Error())
	}
}

func main(){
	Start()
}

