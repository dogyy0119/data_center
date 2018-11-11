package database

import (
	"data_center/config"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type DBase struct {
	db *sql.DB
}

var tableInfo = "(itemid,type_,hostid,name,key_,lastvalue,,lastclock,value_type,datatype,description) VALUES (?, ?, ?, ?, ?,?, ?, ?, ?, ?)"
var tablename = "datamodel"
var insert    = "INSERT INTO "

func (d *DBase)init(user , password ,ip , port string) (error){
	//d.db,err = sql.Open("mysql", user+"："+password+"@tcp("+ip+":"+port+")"+"/"+table)

	d.db,err = sql.Open("mysql", "root:12345678@tcp(localhost:3306)/datacenter")
	if err !=  nil{
		fmt.Println("connet faile")
		return err
	}
	return nil
}

func (d *DBase) intsert(datas []config.DataModel, table string)  {

	tx,_ := d.db.Begin()
	queryStr := insert + tablename + tableInfo
	for _,data := range datas {

		//每次循环用的都是tx内部的连接，没有新建连接，效率高
		tx.Exec(
			queryStr,
			data,
			)
	}
	//最后释放tx内部的连接
	tx.Commit()
}