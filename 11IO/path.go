package main

import (
	"fmt"
	"path"
)

func main(){
	fmt.Println("path.Base: ",path.Base("C:/Users/zzc/go/src/channelDemo"))     //path.Base:  channelDemo   获取路径的最后一个元素
	fmt.Println("path.Clean: ",path.Clean("C:/Users/zzc/./go///src/../../channelDemo/"))   //path.Clean:  C:/Users/zzc/channelDemo         Clean函数通过单纯的词法操作返回和path代表同一地址的最短路径。
	fmt.Println("path.Ext: ",path.Ext("C:/Users/zzc/go/src/channelDemo/main.go"))  //path.Ext:  .go     Ext函数返回path文件扩展名，若没有.  则返回空字符串
	fmt.Println("path.Dir: ",path.Dir("C:/Users/zzc/go/src/channelDemo/main.go"))  //path.Dir:  C:/Users/zzc/go/src/channelDemo     返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。
	fmt.Println("path.IsAbs: ",path.IsAbs("C:/Users/zzc/go/src/channelDemo/main.go"))  //path.IsAbs:  false    IsAbs返回路径是否是一个绝对路径
	dir, file := path.Split("C:/Users/zzc/go/src/channelDemo/main.go")   //Split函数将路径从最后一个斜杠后面位置分隔为两个部分（dir和file）并返回。如果路径中没有斜杠，函数返回值dir会设为空字符串，file会设为path。两个返回值满足path == dir+file。
	fmt.Println("path.Split: ",dir, file)   //path.Split:  C:/Users/zzc/go/src/channelDemo/ main.go
	fmt.Println("path.Join: ",path.Join("c:/Users", "zzc", "go", "src"))    //path.Join:  c:/Users/zzc/go/src   拼接
}
