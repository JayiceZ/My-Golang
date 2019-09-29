package src

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	dbhost = "127.0.0.1:3366"
	dbusername = "root"
	dbpassword = "xxx"
	dbname = "mygo"
)

/*获取连接*/
func getDb() *sql.DB{
	db,err:=sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbusername, dbpassword, dbhost, dbname))
	if err!=nil{
		fmt.Println(err.Error())
	}
	return db
}

/*插入*/

func insert(username, departname string) bool{
	db:=getDb()
	defer db.Close()

	_,err:=db.Exec("INSERT INTO userinfo(username,departname,created) VALUES(?,?,?)",username,departname,time.Now())
	if err!=nil{
		fmt.Println(err.Error())
		return false;
	}
	return true
}


/*修改*/
func update(id int,username string) bool{
	db:=getDb()
	defer db.Close()

	stmt,err:=db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}

	_,err2:=stmt.Exec(id,username)
	if err2!=nil{
		fmt.Println(err2.Error())
		return false
	}

	return true
}

/*删除*/
func delete(id int)bool{
	db:=getDb()
	defer db.Close()

	stmt,err:=db.Prepare("DELECT FROM userinfo WHERE id=?")
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}
	_,err2:=stmt.Exec(id)
	if(err2!=nil){
		fmt.Println(err2.Error())
		return false
	}
	return true
}

func queryOne(id int){
	db:=getDb()
	defer db.Close()

	var username, departname, created string
	err:=db.QueryRow("SELECT * FROM userinfo WHERE uid=?",id).Scan(&username,&departname,&created)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("username: ", username, "departname: ", departname, "created: ", created)
}

func queryAll(){
	db:=getDb()
	defer db.Close()

	rows,err:=db.Query("SELECT * FROM userinfo")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}

	for rows.Next(){
		var username, departname, created string
		err2:=rows.Scan(&username,&departname,&created)
		if err2==nil{
			fmt.Println("username: ", username, "departname: ", departname, "created: ", created)

		}
	}

}


/*事务性的删除（）*/
func newdelete(id int) bool{
	db:=getDb()
	defer db.Close()
	tx, err :=db.Begin()
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}
	_,err2:=tx.Exec("DELECT FROM userinfo WHERE id=?",id)
	if err2!=nil{
		tx.Rollback()
		fmt.Println(err2.Error())
		return false
	}
	tx.Commit()
	return true
}
/*事务性的删除*/
func newinsert(username, departname string) bool{
	db:=getDb()
	defer db.Close()
	tx, err :=db.Begin()
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}
	_,err2:=tx.Exec("INSERT INTO userinfo(username,departname,created) VALUES(?,?,?)",username,departname,time.Now())
	if err2!=nil{
		tx.Rollback()
		fmt.Println(err2.Error())
		return false
	}
	tx.Commit()
	return true
}
