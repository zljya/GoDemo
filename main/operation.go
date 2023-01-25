package main

import (
	"database/sql"
	"fmt"
	"strings"

	//安装驱动 go get -u github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
)

const (
	name     = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbname   = "godemo"
)

var db *sql.DB

func main() {
	initDB()
	//deleteUser("Lily")
	//insertUser("Lily", "18790983652")
	selectUser("Lily")
	//updateUser("Lily", "15038369885")
	deleteUser("Lily")
	//selectUser("Lily")
	//db.Close() //不用关闭，因为是连接池共享的
}
func initDB() {
	dataSourceName := strings.Join([]string{name, ":", password, "@tcp(", ip, ":", port, ")/", dbname, "?charset=utf8"}, "")
	fmt.Println(dataSourceName)
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		println()
		fmt.Println("数据库连接失败")
		return
	}
	//open()在执行时不会真正的与数据库进行连接，只是设置连接数据库需要的参数
	//ping()方法才是连接数据库
	err = db.Ping()
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功")
	db.SetMaxIdleConns(10)

	return
}
func insertUser(name string, phone string) bool {
	ret, err := db.Exec("INSERT INTO user(name, phone) VALUES(?,?)", name, phone)

	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	fmt.Printf("insert %v successfully", id)
	return true
}

func deleteUser(name string) bool {
	ret, err := db.Exec("DELETE FROM user WHERE name = ?", name)
	if err != nil {
		println("删除失败", err)
		return false
	}
	changerow, err := ret.RowsAffected()
	if err != nil {

	}
	println("共删除", changerow, "行记录")
	return true
}
func updateUser(name string, phone string) {
	ret, err := db.Exec("UPDATE user set phone = ? WHERE name = ?", phone, name)
	if err != nil {
		println("删除失败")
	}
	changerow, err := ret.RowsAffected()
	if err != nil {

	}
	println("共删除", changerow, "行记录")
}

func selectUser(username string) {

	rows, err := db.Query("SELECT * FROM user WHERE name = ?", username)
	if err != nil {
		fmt.Println("查询失败")
	}
	for rows.Next() {

		var id int
		var name string
		var phone string

		err := rows.Scan(&id, &name, &phone)
		if err != nil {
			println("查询失败")
		}
		println(id, name, phone)
	}
	rows.Close()

}
