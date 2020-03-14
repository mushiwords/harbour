package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Users struct {
	Id int `json:"id"`
	Pin string `json:"pin"`
	Name string `json:"name"`
	Password string `json:"-"`
	Gender string `json:"gender"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	SignTime string `json:"signTime"`
}

type DbHandler struct {
	dbPool *sql.DB
}

var db  *DbHandler

func (dbH *DbHandler) Exec(cmd string, value ...interface{}) (err error) {
	if _, err := dbH.dbPool.Exec(cmd, value...); err != nil {
		return err
	}
	return nil
}

/**
 * func: db.Query
 * note: should call rows.Close()
 **/
func (dbH *DbHandler) Query(cmd string, value ...interface{}) (*sql.Rows, error) {
	rows, err := dbH.dbPool.Query(cmd, value...)
	if err != nil {
		return nil,   err
	}
	return rows, nil
}

/**
 * func: db.QueryRow
 **/
func (dbH *DbHandler) QueryRow(cmd string, value ...interface{}) *sql.Row {
	return dbH.dbPool.QueryRow(cmd, value...)
}

/**
 * transaction
 **/
func (dbH *DbHandler) Begin() (*sql.Tx, error) {
	tx, err := dbH.dbPool.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func init(){
	dbPool, err := sql.Open("mysql", "root:Captain_10654@tcp(www.yycaptain.com:3306)/db_harbour?charset=utf8")
	if err != nil {
		fmt.Print("db open error :",err.Error())
		return
	}
	dbPool.SetMaxIdleConns(10)
	dbPool.SetMaxOpenConns(30)
	dbPool.SetConnMaxLifetime(time.Minute * 10)

	db = &DbHandler{dbPool: dbPool}

	//1.连接数据库
	//orm.RegisterDataBase("default","mysql","root:Captain_10654@tcp(www.yycaptain.com:3306)/db_harbour?charset=utf8")
	//2.注册表
	//orm.RegisterModel(new(Users))
	//3.生成表
	//1.数据库别名
	//2.是否强制更新
	//3.创建表过程是否可见
	//orm.RunSyncdb("default",false,true)
}

func SelectUserInfo(pin string ) (*Users, error  ) {
	var user Users
	cmd := `select pin, name, gender, phone, email, sign_time from users where pin = ? or name = ?`
	err := db.QueryRow(cmd,pin,pin).Scan(&user.Pin,&user.Name,&user.Gender,&user.Password,&user.Email,&user.SignTime)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		fmt.Println("SelectUserInfo error :",err.Error())
		return nil, err

	}

	return &user, nil
}

