package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io"
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
	androidUrl="https://3s-creative.mobvista.com/get-android-info?access_token=fhtp1plsjkagcd3uia8htaeuj5ojtria"
	iosUrl="https://3s-creative.mobvista.com/get-ios-info?access_token=fhtp1plsjkagcd3uia8htaeuj5ojtria"
	sqlMap=getSqlData()

)

type pack struct {
	Id int
	Platform int
	Package string
	Country string
	App_name string
	Ctime string
	Mtime string
}

type Data struct {
	App_name string `json:"app_name"`
	Package_name string `json:"package_name"`
	Lang string `json:"lang"`
	Short_desc string `json:"short_desc"`
	Updated_at int64 `json:"updated_at"`
}

type SqlStruct struct {
	Platform int
	PackageName string
	Area string
	App_name string
}

func ReadFile(path string){
	//初始化map，主要用来处理数据是否为map（top国家）中的内容
	//比如android_map中有10个key，分别是10个top国家，映射着一个字符串数组，里面存着安卓的该地区的包名，每满100进行处理
	android_Map:=loadMap()//记录插入的安卓包名
	ios_Map:=loadMap()//记录插入的ios包名
	android_update_Map:=loadMap()//记录更新的安卓包名
	ios_update_Map:=loadMap()//记录更新的ios包名
	sqlList:=[]SqlStruct{}//记录要插入的包信息，每满5000个insert一次
	updateSqlList:=[]SqlStruct{} //记录要更新的包信息，每满500个就进行update操作
	file,_ := os.Open(path)
	defer file.Close()
	inputReader := bufio.NewReader(file)
	for{
		//每次读一行
		str,err :=inputReader.ReadString('\n')
		//若读到了空的一行，则证明读完了，直接break
		if err == io.EOF{
			break
		}

		if err != nil{
			break //err or EOF
		}
		//把一行的数据分割为字符串数组
		res:=strings.Split(str,"	")
		//包名在下标1
		packageName:=res[1]
		//地区在下标2
		area:=res[3]
		_,flag:=android_Map[area]
		//如果是top地区，则要进行处理
		if(flag){
			//键值为包名+地区，因为这是数据辨别属性的唯一标识
			key:=packageName+area
			//查找数据库中是否已经有该包名+地区了
			_,have_this_data:=sqlMap[key]
			//若没有，则进行insert操作
			if(!have_this_data){
				//ios包名格式都大致为："id12344566"这样,所有不以id开头的和以id.开头的,都视为安卓包
				if !strings.HasPrefix(packageName,"id")||strings.HasPrefix(packageName,"id."){
					//往安卓map中对应的area中的数组填包名
					android_Map[area]=append(android_Map[area],packageName)
					//若包名达到100，说明已经达到url拼接最大长度了，可以进行一次抓取了
					if len(android_Map[area])>=100{
						//往安卓map中对应的area中的数组填包名
						androidMapSql(&sqlList,area,android_Map[area])
						//处理完后进行清空
						android_Map[area]=make([]string,0)
						//若抓取数据填入sqlList后，长度超过5000，那么可以进行一次数据库批量处理了
						if len(sqlList)>=5000 {
							goSql(sqlList)
							//处理后sqlList要清空
							sqlList=make([]SqlStruct,0)
						}
					}
					//ios的处理方法
				}else{
					ios_Map[area]=append(ios_Map[area],packageName)
					if len(ios_Map[area])>=100{
						iosMapSql(&sqlList,area,ios_Map[area])
						//清空数据
						ios_Map[area]=make([]string,0)
						if len(sqlList)>=5000 {
							goSql(sqlList)
							//清空数据
							sqlList=make([]SqlStruct,0)
						}
					}
				}
				//若有，则要进行update操作，更新app_name
			}else{
				//若是安卓
				if !strings.HasPrefix(packageName,"id")||strings.HasPrefix(packageName,"id."){
					android_update_Map[area]=append(android_update_Map[area],packageName)
					if len(android_update_Map[area])>=100{
						androidMapSql(&updateSqlList,area,android_update_Map[area])
						android_update_Map[area]=make([]string,0)
						//若超过了500，则执行update
						if len(updateSqlList)>=500{
							updateSql(updateSqlList)
							updateSqlList=make([]SqlStruct,0)
						}
					}
					//若是ios
				}else{
					ios_update_Map[area]=append(ios_update_Map[area],packageName)
					if len(ios_update_Map)>=100{
						iosMapSql(&updateSqlList,area,ios_update_Map[area])
						ios_update_Map[area]=make([]string,0)
						if len(updateSqlList)>=500{
							updateSql(updateSqlList)
							updateSqlList=make([]SqlStruct,0)
						}
					}
				}

			}
		}
	}
	//若文件中的行数读完了，但是一般来说map中还存有数据，只是还没到100，没有达到处理条件。这时候需要遍历一次四个map，把数据处理完
	for android_key:=range android_Map{
		androidMapSql(&sqlList,android_key,android_Map[android_key])
	}
	for ios_key:=range ios_Map{
		iosMapSql(&sqlList,ios_key,ios_Map[ios_key])
	}
	for android_update_key:=range android_update_Map{
		androidMapSql(&updateSqlList,android_update_key,android_update_Map[android_update_key])
	}
	for ios_update_key:=range ios_update_Map{
		iosMapSql(&updateSqlList,ios_update_key,ios_update_Map[ios_update_key])
	}
	goSql(sqlList)
	updateSql(updateSqlList)
}

//进行安卓url的拼接
func androidMapSql(sqlList *[]SqlStruct,area string,package_list []string){
	if len(package_list)==0{
		return
	}
	android_url:=androidUrl+"&country="+area+"&package_name="
	for i:=0;i< len(package_list);i++{
		if i!=0{
			android_url+=","
		}
		android_url=android_url+package_list[i]
	}
	//拼接完之后，取抓取数据
	getUrlData(sqlList,android_url,area,1)

}
//进行ios url的拼接(同理)
func iosMapSql(sqlList *[]SqlStruct,area string,package_list []string){
	if len(package_list)==0{
		return
	}
	ios_url:=iosUrl+"&country="+area+"&package_name="
	for i:=0;i< len(package_list);i++{
		if i!=0{
			ios_url+=","
		}
		ios_url=ios_url+package_list[i]
	}
	getUrlData(sqlList,ios_url,area,2)
}

//通过url对数据进行抓取
func getUrlData(sqlList *[]SqlStruct,url string,area string,platform int){
	res,err1:=http.Get(url)
	/*若出错，记录*/
	if err1!=nil{
		log.Fatal(err1)
	}
	defer res.Body.Close()
	/*读取url上的数据，以body来接收，body是json格式的字节数组*/
	body,err2:=ioutil.ReadAll(res.Body)
	if err2!=nil{
		log.Fatal(err2)
	}
	dataMap:=make(map[string]Data)
	json.Unmarshal(body,&dataMap)
	//把抓取的数据放在sqlList中
	for key:=range dataMap{
		*sqlList=append(*sqlList,SqlStruct{platform,key,area,dataMap[key].App_name})
	}
}

//若满了5000，则批量执行一次insert
func goSql(sqlList []SqlStruct){
	if len(sqlList)==0{
		return
	}
	sqlStr:="INSERT INTO package_app_2 (platform,package,country,app_name) VALUES "
	db:=Get()
	for i:=0;i<len(sqlList);i++{
		if i!=0{
			sqlStr+=","
		}
		sqlStr+="("+strconv.Itoa(sqlList[i].Platform)+",\""+sqlList[i].PackageName+"\",\""+sqlList[i].Area+"\",\""+sqlList[i].App_name+"\")"
	}
	db.Exec(sqlStr)
}

//执行update
func updateSql(updateSqlList []SqlStruct){
	if len(updateSqlList)==0{
		return
	}
	db:=Get()
	for i:=0;i<len(updateSqlList);i++{
		str:="UPDATE package_app_2 SET app_name=\""+updateSqlList[i].App_name+"\" WHERE package=\""+updateSqlList[i].PackageName+"\" AND country=\""+updateSqlList[i].Area+"\""
		db.Exec(str)
	}
}

func loadMap() map[string][]string{
	loadMap:=make(map[string][]string)
	loadMap["CN"]=make([]string,0)
	loadMap["US"]=make([]string,0)
	loadMap["JP"]=make([]string,0)
	loadMap["IN"]=make([]string,0)
	loadMap["ID"]=make([]string,0)
	loadMap["KR"]=make([]string,0)
	loadMap["MY"]=make([]string,0)
	loadMap["TH"]=make([]string,0)
	loadMap["UK"]=make([]string,0)
	loadMap["BR"]=make([]string,0)
	return loadMap
}

func Get() *sqlx.DB{
	db,err:=sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbusername, dbpassword, dbhost, dbname))
	if err!=nil{
		fmt.Println(err.Error())
	}
	return db
}


func goShell(){
	str,_:=os.Getwd()
	exec.Command("aws","s3","cp","s3://mob-emr-test/adn/adn_data/detect_stat_result_new/detect_detail_"+time.Now().AddDate(0,0,-1).Format("20060102")+".xls",str+"/package/")
}

//获取数据库中已有的数据，来判断文件中的数据是否已经在数据库中存在，来决定需要update还是需要insert
func getSqlData() map[string]int{
	db:=Get()
	rows, err := db.Query("SELECT package,country FROM package_app_2")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	myMap:=make(map[string]int,0)
	for rows.Next() {
		pack:=pack{}
		err = rows.Scan(&pack.Package,&pack.Country)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		//数据的唯一标识就是包名+地区
		key:=pack.Package+pack.Country
		myMap[key]=1
	}
	return myMap
}

func test(path string){
	file,_ := os.Open(path)
	defer file.Close()
	inputReader := bufio.NewReader(file)
	sum:=0
	for{
		sum++
		str,err :=inputReader.ReadString('\n')

		if err == io.EOF{
			break
		}

		if err != nil{
			return //err or EOF
		}
		fmt.Print(sum)
		fmt.Println(str)
	}
	fmt.Println("finished")
}

func main(){
	/*
	fmt.Println(time.Now())
	Path, _ := os.Getwd()
	Absolute_path := strings.Replace(Path, "\\", "/", -1)
	readFile(Absolute_path + "/detect_detail_20181001.xls")
	fmt.Println(time.Now())
	*/
	Path, _ := os.Getwd()
	Absolute_path := strings.Replace(Path, "\\", "/", -1)
	test(Absolute_path+"/package/detect_detail_scenario_"+time.Now().AddDate(0,0,-1).Format("20060102")+".xls")

}
