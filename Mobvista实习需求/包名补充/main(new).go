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
	DefaultDay=1
	androidUrl="https://3s-creative.mobvista.com/batch-base-info/android?access_token=fhtp1plsjkagcd3uia8htaeuj5ojtria"
	iosUrl="https://3s-creative.mobvista.com/batch-base-info/ios?access_token=fhtp1plsjkagcd3uia8htaeuj5ojtria"
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


/*代码逻辑大致为：准备四个Map，分别用来储存 1要插入的安卓包，2要插入的ios包，3要更新的安卓包，4要更新的ios包。以安卓插入包Map为例：
Map中有10个key，代表10个top地区，每个key对应一个字符串数组（也就是说同一个数组中的包名都是同一个地区的），储存着要处理的包名，
若数组长度达到100，则进行一次接口数据抓取（url最多拼接100个package_name），并清空该数组（因为已经处理完了）
把抓到的数据存在sqlList数组中，sqlList数组每满5000个数据，就进行一次批量insert处理
*/

/*另外还准备一个map，把数据库中的数据都拿出来，形式为 key:包名+地区 value:app_name     （测试过用map存一亿个数据，没有出现内存问题，
而数据表中数据基量为10000多，每天新增1000个不到，基本不会出现内存问题）*/
func ReadFile(path string){
	//比如android_map中有10个key，分别是10个top国家，映射着一个字符串数组，里面存着安卓的该地区的包名，每满100进行处理
	android_Map:=loadMap()//记录要插入的安卓包名，
	ios_Map:=loadMap()//记录插入的ios包名
	android_update_Map:=loadMap()//记录更新的安卓包名
	ios_update_Map:=loadMap()//记录更新的ios包名
	sqlList:=[]SqlStruct{}//记录要插入的包信息，每满5000个insert一次
	updateSqlList:=[]SqlStruct{} //记录要更新的包信息，每满2000个就进行update操作
	file,_ := os.Open(path)
	defer file.Close()
	inputReader := bufio.NewReader(file)
	//一行行读文件
	for{
		str,err :=inputReader.ReadString('\n')
		//若读到了空的一行，则证明读完了，直接break
		if err == io.EOF{
			break
		}

		if err != nil{
			break //err or EOF
		}
		//把一行的数据分割为字符串数组，以一个tab为为边界来切割
		res:=strings.Split(str,"	")
		//包名在下标1
		packageName:=res[1]
		//地区在下标2
		area:=res[3]
		//判断该条数据的地区是否是top地区，因为map中10个key就是10个top地区
		_,flag:=android_Map[area]
		//如果是top地区，则要进行处理
		if(flag){
			//键值为包名+地区，因为这是数据辨别属性的唯一标识
			key:=packageName+area
			//查找数据库中是否已经有该包名+地区了
			_,have_this_data:=sqlMap[key]

			//若没有，则进行insert操作
			if(!have_this_data){
				//ios包名格式都大致为："id12344566"这样,所有不以id开头的 和 以id.开头的(因为有些安卓包名类如 id.xxx.xxx),都视为安卓包
				if !strings.HasPrefix(packageName,"id")||strings.HasPrefix(packageName,"id."){
					//往安卓map中对应的area中的数组填包名
					android_Map[area]=append(android_Map[area],packageName)
					//若包名达到100，说明已经达到url拼接最大长度了，可以进行一次抓取了
					if len(android_Map[area])>=100{
						//对这100个包名进行url拼接，然后去接口进行数据抓取，抓取后会放在sqlList中
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


				//若有该包名+地区，则 可能 要进行update操作，但是还要看从接口抓下来的app_name和数据表中的是否一致，不一致才要更新
			}else{
				//若是安卓
				if !strings.HasPrefix(packageName,"id")||strings.HasPrefix(packageName,"id."){
					android_update_Map[area]=append(android_update_Map[area],packageName)
					if len(android_update_Map[area])>=100{
						//操作类似，只不过储存包名的结构变成了updateSqlList
						androidMapSql(&updateSqlList,area,android_update_Map[area])
						android_update_Map[area]=make([]string,0)
						//若超过了500，则执行update
						if len(updateSqlList)>=1000{
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
						if len(updateSqlList)>=1000{
							updateSql(updateSqlList)
							updateSqlList=make([]SqlStruct,0)
						}
					}
				}

			}
		}
	}
	//若文件读完了，但是一般来说map和sqlList中还存有数据，只是分别还没达到100和5000，没有达到处理条件。这时候需要遍历一次四个map，把数据全部读出来
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
	//insert
	goSql(sqlList)
	//update
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
	//拼接完之后，取抓取数据
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
	/*读取url上的数据，以body来接收，body是json格式的字节数组，抓到的json解析后以包名为key,具体数据为value的map*/
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
		//若该包名+数据的app_name和在数据库中的不一致，则要更新，否则不用
		key:=updateSqlList[i].PackageName+updateSqlList[i].Area
		//若app_name更新了，才update，否则不进行此操作
		if sqlMap[key]!=updateSqlList[i].App_name{
			str:="UPDATE package_app_2 SET app_name=\""+updateSqlList[i].App_name+"\" WHERE package=\""+updateSqlList[i].PackageName+"\" AND country=\""+updateSqlList[i].Area+"\""
			db.Exec(str)
		}
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

//执行shell指令
func goShell(){
	path,_:=os.Getwd()
	cmd:=exec.Command("aws","s3","cp","s3://mob-emr-test/adn/adn_data/detect_stat_result_new/detect_detail_scenario_"+time.Now().AddDate(0,0,-DefaultDay).Format("20060102")+".xls",path+"/package/")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

//获取数据库中已有的数据，来判断文件中的数据是否已经在数据库中存在，来决定需要update还是需要insert
func getSqlData() map[string]string{
	db:=Get()
	rows, err := db.Query("SELECT package,country,app_name FROM package_app_2")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	myMap:=make(map[string]string,0)
	for rows.Next() {
		pack:=pack{}
		err = rows.Scan(&pack.Package,&pack.Country,&pack.App_name)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		//数据的唯一标识就是包名+地区
		key:=pack.Package+pack.Country
		//map的形式为  key:包名+地区 value：app_name
		myMap[key]=pack.App_name
	}
	return myMap
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
	//若输入的参数不能转为int，则继续默认执行
	if err!=nil{
		log.Fatal(err)
		DefaultDay=1
		return
	}
	//把默认值改为传入值
	DefaultDay=get
}


func main(){
	//监听命令行参数
	Listen()
	fmt.Println("start!")
	Path, _ := os.Getwd()
	Absolute_path := strings.Replace(Path, "\\", "/", -1)
	//执行shell指令，即把文件从s3下载到本地，再进行后续读取文件
	goShell()
	ReadFile(Absolute_path+"/package/detect_detail_scenario_"+time.Now().AddDate(0,0,-DefaultDay).Format("20060102")+".xls")
	fmt.Println("success!")
}
