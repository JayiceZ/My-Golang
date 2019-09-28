package main

import (
	"fmt"
	"strings"
)

func main(){
	str1:="jayice"


	//Contain
	fmt.Println(strings.Contains(str1,"jay"))  //true

	//判断是否包含多个字符
	fmt.Println(strings.ContainsAny(str1,"j & r"))  //true
	fmt.Println(strings.ContainsAny(str1,"q & s")) //false

	fmt.Println(strings.ContainsRune(str1,'c')) //true
	fmt.Println(strings.ContainsRune(str1,99))  //true 99是c的unicode编码

	fmt.Println(strings.Count("jayice","j")) //1
	fmt.Println(strings.Count("jayicieeee","e"))  //4
	fmt.Println(strings.Count("jayiceeee","ee"))  //2



	//Index
	fmt.Println(strings.Index("jay","j"))  //0
	fmt.Println(strings.Index("jay","a"))  //1
	fmt.Println(strings.Index("jay","y"))  //2
	fmt.Println(strings.Index("jayj","j")) //0
	fmt.Println(strings.Index("jay","ja"))  //0
	fmt.Println(strings.Index("jay","e")) //-1

	//IndexByte 和Index类似，只不过传入的是字符
	fmt.Println(strings.IndexByte("jay",'j')) //0
	fmt.Println(strings.IndexByte("jay",'r'))  //-1

	//LastIndex  从尾部开始检索
	fmt.Println(strings.LastIndex("jayyya","a"))  //5
	fmt.Println(strings.LastIndex("eeiwiwee","e")) //7

	//func EnqualFold(s string, t string) bool,两个字符串比较，忽略大小写，返回bool类型
	fmt.Println(strings.EqualFold("jay","JAY"))  //true
	fmt.Println(strings.EqualFold("jay","ice"))   //false

	//func Join(s []string, seq string) string将字符串数组按照指定的分隔符拼接成字符串
	s:=[]string{"jay","ice","niupi"}
	fmt.Println(strings.Join(s,","))   //jay,ice,niupi
	fmt.Println(strings.Join(s,""))   //jayiceniupi

	//func Map(mapping func(rune)rune, s string) string, 如果mapping方法返回个合法的字符串，改方法返回一个由mapping方法修改过的复制过来的字符串。
	//书写func函数来对每一个字符进行处理
	myFunc:=func(r rune) rune{
		if r>='A'&&r<='Z'{
			return 'A'
		}else if(r>='a'&&r<='z'){
			return 'a'
		}
		return 'b'
	}
	//调用myFunc函数，来处理后面的字符串
	fmt.Println(strings.Map(myFunc,"nihaoma,WNUDW...."))   //aaaaaaabAAAAAbbbb


	//func Repeat(s string, count int)string，改方法返回一个新的重复指定次数的字符串
	fmt.Println("ba"+strings.Repeat("na",2))  //banana
	fmt.Println("jay"+strings.Repeat("jayice",3))  //jayjayicejayicejayice


	//func Replace(s, old, new string, count int)string返回一个新的字符串，参数s是原来的字符串，old是需要被替换掉的字符串，new是要替代old的字符串，count是替换的次数，如果为-1，则为全部替换。
	fmt.Println(strings.Replace("jay jay","jay","newjay",1))  //newjay jay
	fmt.Println(strings.Replace("jay jay","jay","newjay",2))  //newjay newjay
	fmt.Println(strings.Replace("jay jay","jay","newjay",-1))  //newjay newjay


	//func Split(s, seq string)[]string将字符串按照指定的字符串分割生一个字符串数组
	fmt.Println(strings.Split("jay,ice,niupi",","))  //[jay ice niupi]
	fmt.Println(strings.Split("jay ice niupi,"," "))   //[jay ice niupi,]


	//func SplitN(s, seq string, count int)[]string将字符串按照指定的字符串分割生一个指定元素数量的字符串数组。该方法返回的数组将不保留分隔符。count参数为-1时效果如Split。
	fmt.Println(strings.SplitN("jay,ice,niupi",",",2))  //[jay ice,niupi]
	fmt.Println(strings.SplitN("jay,ice,niupi",",",3))  //[jay ice niupi]
	fmt.Println(strings.SplitN("jay,ice,niupi",",",-1))  //[jay ice niupi]  (效果和Split一样)


	//func SplitAfter(s, seq string)[]string将字符串按照指定的字符串分割生一个字符串数组。该方法返回的数组将保留分隔符，且至于每个元素的末端。
	fmt.Println(strings.SplitAfter("jay/ice/niupi","/"))  //[jay/ ice/ niupi]


	//func Title(s string)string该方法返回一个新的字符串，该字符串把原字符串的单词首字母改为大写，对中文没有效果。
	fmt.Println(strings.Title("jay plays basketball"))   //Jay Plays Basketball


	//func ToTitle(s string)string将字符串转为大写字母。
	fmt.Println(strings.ToTitle("jayice"))  //JAYICE

	//func ToLower(s string)string将字符串转为小写字母。
	fmt.Println(strings.ToLower("JAYICE"))  //jayice

	//func ToUpper(s string)string将字符串转为大写字母。
	fmt.Println(strings.ToUpper("jayice"))   //JAYICE


	//func Trim(s, cutset string)string去除字符串中首尾指定的字符。
	fmt.Println(strings.Trim("!!!jayice!!!","!"))  //jayice


	//func TrimLeft(s, cutset string)string去除字符串中左侧指定的字符。
	fmt.Println(strings.TrimLeft("!!!jayice!!!","!"))  //jayice!!!

	//func TrimSpace(s stirng)string去除字符串中首尾的空白部分。
	fmt.Println(strings.TrimSpace("   jayice   "))   //jayice





}
