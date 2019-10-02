package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	dbhost = "127.0.0.1:3366"
	dbusername = "root"
	dbpassword = "123"
	dbname = "mygo"
)

type User struct {
	Uid int   `json:"uid"`
	Username string `json:"username"`
	Departname string  `json:"departname"'`
	Created string  `"json:created"`
}

func Get() *sqlx.DB{
	db,err:=sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbusername, dbpassword, dbhost, dbname))
	if err!=nil{
		fmt.Println(err.Error())
	}
	return db
}

func Search(id int) User{
	db:=Get()
	myUser:=User{}
	//Unsafe用来解决字段名不匹配的问题
	err:=db.Unsafe().Get(&myUser,"SELECT * FROM userinfo WHERE uid=? LIMIT 1",id)
	if err!=nil{
		fmt.Println(err.Error())
	}

	return myUser
}

func SearchAll() []User{
	db:=Get()
	myUser:=[]User{}
	err:=db.Select(&myUser,"select * from userinfo")
	if err!=nil{
		fmt.Println(err.Error())
	}
	return myUser
}


func Delete(id int) bool{
	db:=Get()
	tx,_:=db.Begin()
	_,err:=tx.Exec("delete from userinfo where uid=?",id)
	if err!=nil{
		fmt.Println(err.Error())
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func Insert(user *User) bool{
	db:=Get()
	tx,_:=db.Begin()
	_,err:=tx.Exec("INSERT INTO userinfo (username.departname,created) VALUES (?,?,?)",user.Username,user.Departname,user.Created)
	if err!=nil{
		fmt.Println(err.Error())
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

