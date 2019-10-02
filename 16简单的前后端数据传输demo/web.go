package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
)

func Start(){
	//这里主要用于打开界面
	http.HandleFunc("/login",Login)
	http.HandleFunc("/index",Index)

	http.ListenAndServe(":8080",nil)
}

func Login(w http.ResponseWriter,r *http.Request){
	f,_:=os.Open("./src/view/login.html")
	defer f.Close()
	io.Copy(w,f)
}

func Index(w http.ResponseWriter,r *http.Request){
	f,_:=os.Open("./src/view/index.html")
	defer f.Close()
	io.Copy(w,f)
}


//这里的函数用于处理前端请求，由前端指定具体由哪一个函数处理
func add(w httptest.ResponseRecorder,r http.Request){
	user:=&User{}
	r.ParseForm()
	user.Username=r.Form.Get("username")
	user.Departname=r.Form.Get("departname")
	user.Created=r.Form.Get("created")

	flag:=Insert(user)
	if flag==true{
		w.Write([]byte("insert successful"))
	}else{
		w.Write([]byte("fail"))
	}
}

func query(w httptest.ResponseRecorder,r http.Request){
	r.ParseForm()
	idstr:=r.Form.Get("uid")
	id,_:=strconv.ParseInt(idstr,10,64)
	user:=Search(int(id))
	buf,_:=json.Marshal(user)

	w.Write(buf)
}
