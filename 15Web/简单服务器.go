package main

import (
	"fmt"
	"log"
	"net/http"
)

type mux struct {
}

//ServeHTTP的函数名不能变
func (m *mux) ServeHTTP	(w http.ResponseWriter,r *http.Request){
	if r.URL.Path=="/first"{
		First(w,r)
		return
	}else if r.URL.Path=="/second"{
		Second(w,r)
		return
	}
	http.NotFound(w,r)
	return
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

