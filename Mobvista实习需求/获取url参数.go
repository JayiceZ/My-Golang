package main

import (
	"fmt"
	"net/url"
)

func main(){
	str:="http://www.baidu.com?id=1&name=jay"
	u,_:=url.Parse(str)
	res,_:=url.ParseQuery(u.RawQuery)
	fmt.Println(res.Get("id")) //1
	fmt.Println(res.Get("name")) //jay
}
