package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql" //导入连接MySQL数据库的驱动包

// User 结构体的每个字段表示数据表的列，结构体的字段首字母必须是大写的。
type User struct {
	ID    int
	Name  string
	Phone string
}

const (
	name     = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbname   = "godemo"
	driver   = "mysql"
)

var db *gorm.DB

// TableName gorm框架默认表名为struct名的复数
func (User) TableName() string {
	return "user"
}
func main() {
	dsn := name + ":" + password + "@tcp(" + ip + ":" + port + ")/" + dbname + "?charset=utf8"
	var err error
	db, err = gorm.Open(driver, dsn)
	if err != nil {
		println("连接失败")
		panic(err)
	}
	println("数据库连接成功")
	user := User{Name: "李四", Phone: "19806728736"}
	addUser(user)
	defer db.Close()
}
func addUser(user User) {
	//查询逐渐是否存在，不存在返回true
	if db.NewRecord(user) {
		res := db.Select("name", "phone").Create(&user)
		if res.RowsAffected != 0 {
			println("插入成功")
		} else {
			println("插入失败")
		}
	}
}

func addUsers(users []User) {
	//查询逐渐是否存在，不存在返回true
	if err := db.Create(&users).Error; err != nil {
		fmt.Println("插入失败", err)
	}

}
func selectUser(name string) []User {
	var users []User
	db.Where("name = ?", name).Order("phone desc").Find(&users)
	return users
}
func updateUser(user User) {
	db.Where("name = ?", user.Name).Updates(user)
}
func deleteUser(name string) {
	users := User{}
	db.Where("name = ?", name).Find(&users)
	db.Delete(&users)
}
