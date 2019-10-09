package main

import (
	_ "database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	dbhost = "13.228.176.169:3306"
	dbusername = "root"
	dbpassword = "xxx"
	dbname = "mob_adn"
)

type Value struct{
	Id int `json:"id"`
	Ckey string `json:"ckey"`
	Cvalue string `json:"cvalue"`
	Cname string `json:"cname"`
}


type Url struct{
	Url string `json:"url"`
}

type Click_report_response struct{
	XMLName xml.Name `xml:"click_report_response"`
	Clicks Clicks
	Description string   `xml:",innerxml"`
}

type Clicks struct{
	XMLName xml.Name `xml:"clicks"`
	Click []Click `xml:"click"`
	Description string   `xml:",innerxml"`
}

type Click struct{
	XMLName xml.Name `xml:"click"`
	Uuid string `xml:"sub_id_1"`
	ClickTime string `xml:"click_date"`
	Clickcid string `xml:"sub_id_2"`
	Clicks int `xml:"total_clicks"`
	Ua string `xml:"user_agent"`
}

var(
	DirName=""
)

func Get() *sqlx.DB{
	db,err:=sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbusername, dbpassword, dbhost, dbname))
	if err!=nil{
		fmt.Println(err.Error())
	}
	return db
}

/*数据库中获取到对应的cvalue，该值是一个json格式的字符串，解析后获得map*/
func SelectValue()[]byte{
	db:=Get()
	value:=Value{}
	db.Get(&value,"SELECT cvalue FROM config_system WHERE id=440")
	return []byte(value.Cvalue)
}

/*对字节数组及进行解析*/
func Analysis(buf []byte)map[string]Url{
	mapS := make(map[string]Url)
	json.Unmarshal(buf,&mapS)
	return mapS
}
/*遍历map*/
func Travel(){
	buf:=SelectValue()
	mapS:=Analysis(buf)
	for key:=range mapS{
		//获取每一条url
		url:=mapS[key].Url
		//访问url
		res,err1:=http.Get(url)
		/*若出错，记录*/
		if err1!=nil{
			log.Fatal(err1)
		}
		defer res.Body.Close()
		/*读取url上的数据，以body来接收，body是xml格式的字节数组*/
		body,err2:=ioutil.ReadAll(res.Body)
		if err2!=nil{
			log.Fatal(err2)
		}
		crr:=Click_report_response{}
		/*解析数据*/
		err3:=xml.Unmarshal(body,&crr)
		if err3!=nil{
			log.Fatal(err3)
		}

		for Single:=range crr.Clicks.Click{
			//在当天文件夹路径下加上文件
			FileName:=DirName+"/"+crr.Clicks.Click[Single].Uuid+".txt"
			_,err4:=os.Stat(FileName)
			if err4!=nil{
				file,_:=os.Create(FileName)
				defer file.Close()
				file,_=os.OpenFile(FileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
				file.WriteString("uuid")
				file.Write([]byte("\t"))
				file.WriteString("clickTime")
				file.Write([]byte("\t"))
				file.WriteString("clickid")
				file.Write([]byte("\t"))
				file.WriteString("clicks")
				file.Write([]byte("\t"))
				file.WriteString("ua")
				file.Write([]byte("\n"))
			}
			file,_:=os.OpenFile(FileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
			file.WriteString(crr.Clicks.Click[Single].Uuid)
			file.Write([]byte("\t"))
			file.WriteString(crr.Clicks.Click[Single].ClickTime)
			file.Write([]byte("\t"))
			file.WriteString(crr.Clicks.Click[Single].Clickcid)
			file.Write([]byte("\t"))
			file.WriteString(strconv.Itoa(crr.Clicks.Click[Single].Clicks))
			file.Write([]byte("\t"))
			file.WriteString(crr.Clicks.Click[Single].Ua)
			file.Write([]byte("\n"))
		}
	}
}

/*日常任务，每天凌晨0点进行操作*/
func Daily(){
	for {
		/*创建当天的文件夹*/
		//获取绝对路径
		str,_:=os.Getwd()
		str2:=str+"/"+time.Now().Format("2006-01-02")
		DirName=strings.Replace(str2,"\\","/",-1)
		os.Mkdir(DirName,0777)
		fmt.Println(os.Getwd())
		/*开始抓取数据*/
		Travel()
		fmt.Println("Success")
		now := time.Now()  //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24) //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))//计算当前时间到凌晨的时间间隔，设置一个定时器
		//等待计时结束后，才会进入下一次循环
		<-t.C
	}
}

func main(){
	Daily()
}
