package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
		//若输入用户名长度为0
		if len(r.Form["username"][0])==0{
			fmt.Fprintf(w, "username: null or empty \n")
		}
		//把表单中的age转为int类型，若无法转换，说明输入内容有错误
		age, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil{
			fmt.Fprintf(w, "age: The format of the input is not correct \n")
		}
		//若年龄小于18
		if age < 18{
			fmt.Fprintf(w, "age: Minors are not registered \n")
		}
		//正则表达式检验邮箱输入
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`,
			r.Form.Get("email")); !m {
			fmt.Fprintf(w, "email: The format of the input is not correct \n")
		}
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

