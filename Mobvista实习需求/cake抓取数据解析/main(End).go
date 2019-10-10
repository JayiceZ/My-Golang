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
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	dbhost = "13.228.176.169:3306"
	dbusername = "root"
	dbpassword = "TNKq6de8ttjGq4aB"
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
	DefaultDay=1
)

func WhichDate(){

}

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
	defer db.Close()
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
		USA, err := time.LoadLocation("America/Los_Angeles")
		if err != nil {
			log.Fatal(err)
		}
		//从数据库拿下来的url,先拼接上两个参数：开始时间和结束时间
		url=url+"&start_date="+time.Now().In(USA).AddDate(0,0,-DefaultDay).Format("01/02/2006")+"&end_date="+time.Now().In(USA).AddDate(0,0,1-DefaultDay).Format("01/02/2006")
		//start作为start_at_row
		start:=0
		for{
			//拼接，每次都读500条
			tempUrl:=url+"&start_at_row="+strconv.Itoa(start)+"&row_limit="+strconv.Itoa(500)
			res,err1:=http.Get(tempUrl)
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
				log.Fatal(err3) //
			}
			//dataSum用来记录本条url本次所读出来的数据数，在以下的for循环中，每循环一次就是一个数据。若循环结束dataSum小于500
			//则说明已经读完了，直接break.否则进入下一次循环
			dataSum:=0
			for Single:=range crr.Clicks.Click{
				//在当天文件夹路径下加上文件
				dataSum+=1
				FileName:=DirName+"/"+crr.Clicks.Click[Single].Uuid+".xls"
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
			//若小于500，则说明读完了，
			if dataSum<500{
				break
			}
			//否则start自增501,接着读
			start+=501
		}

	}
}

/*日常任务，每天凌晨0点进行操作*/
func Daily(){
		/*创建当天的文件夹*/
		//获取绝对路径
		str,_:=os.Getwd()
		//文件路径为项目文件中的form文件夹，里面存放着每一天对应的文件夹
		USA, err := time.LoadLocation("America/Los_Angeles")
		if err != nil {
			log.Fatal(err)
		}
		str2:=str+"/form/"+time.Now().In(USA).AddDate(0,0,-DefaultDay).Format("2006-01-02")
		DirName=strings.Replace(str2,"\\","/",-1)
		_,err2:=os.Stat(DirName)
		//若文件已经存在，则先删除文件
		if err2==nil{
			os.RemoveAll(DirName)
		}

		os.MkdirAll(DirName,0777)
		fmt.Println(DirName)
		/*开始抓取数据*/
		fmt.Println("Start")
		Travel()
		fmt.Println("Success")
}

func GoShell(){
	str,_:=os.Getwd()
	//文件路径为项目文件中的form文件夹，里面存放着每一天对应的文件夹
	USA, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatal(err)
	}
	str2:=str+"/form/"+time.Now().In(USA).AddDate(0,0,-DefaultDay).Format("2006-01-02")
	DirName=strings.Replace(str2,"\\","/",-1)
	cmd := exec.Command("aws","s3","sync",DirName,"s3://mob-emr-test/adn/cake/"+time.Now().In(USA).AddDate(0,0,-DefaultDay).Format("2006-01-02")+"/")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func Listen(){
	//监听命令行参数,若什么都不输入，参数有一个（路径）。若输入了，如:go run xxx.go 123  则会有一个123参数
	args:=os.Args
	//若不输入，默认为取前一天
	if len(args)<2{
		DefaultDay=1
		return
	}
	var ele string
	ele=args[1]
	//转int
	get,err:=strconv.Atoi(ele)
	if err!=nil{
		//若输入的参数不能转为int，则继续默认执行
		log.Fatal(err)
		DefaultDay=1
	}
	//把默认值改为传入值
	DefaultDay=get
}

func main(){
	//在运行之前监听命令行参数，以确定DefaultDay的值
	Listen()
	Daily()
	GoShell()
}
