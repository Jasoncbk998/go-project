/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/29 9:39 下午
 **/
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//连接实例
func Conn(user, password, host, db, port string) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, db)
	db_instance, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatalf("连接mysql err:", err)
	}
	db_instance.SingularTable(true)
	return db_instance
}

//表结构体
type XzAutoServerConf struct {
	Id         string `gorm:"column:id"`
	GroupZone  string `gorm:"column:group_zone"`
	ServerId   int    `gorm:"column:server_id"`
	OpenTime   string `gorm:"coloumn:open_time"`
	ServerName string `gorm:"coloumn:server_name"`
	Status     int    `gorm:"coloumn:status"`
	Username   string `gorm:"coloumn:username"`
}

func main() {
	user := "root"
	password := "Liuxu!99454"
	host := "127.0.0.1"
	port := "3306"
	db_name := "test"
	db := Conn(user, password, host, db_name, port)
	db.LogMode(true)
	defer db.Close()
	//insert_data(db)
	groupby_data(db)

}

func insert_data(db *gorm.DB) {
	add_data := XzAutoServerConf{GroupZone: "26", ServerId: 10, ServerName: "abc", Username: "aaaaaaa", Status: 26}
	create := db.Create(&add_data)
	if create.Error != nil {
		log.Fatal(create.Error)
	}
}

var rows []XzAutoServerConf

func delete_data(db *gorm.DB) {

	err := db.Model(&rows).Where("group_zone=?", "21").Delete(&XzAutoServerConf{}).Error
	if err != nil {
		log.Fatal(err)
	}
}

func search_data(db *gorm.DB) {
	db.Table("xz_auto_server_conf").Where("status=?", "22").Select([]string{"group_zone", "server_id", "open_time", "server_name", "username"}).Find(&rows)
	for _, v := range rows {
		fmt.Println(v.Username)
	}
}

func order_data(db *gorm.DB) {
	db.Table("xz_auto_server_conf").Where("status>=?", 23).Select([]string{"username", "id"}).Order("status desc").Find(&rows)
	for _, v := range rows {
		fmt.Println(v.Id)
	}
}

//封装对象必须大小写
type Result struct {
	Group_Zone string
	Num_        int64
}

func groupby_data(db *gorm.DB) {
	var results []Result
	db.Table("xz_auto_server_conf").Select("group_zone,count(*) as num_").Group("group_zone").Scan(&results)
	fmt.Println(results)
}

func update_data(db *gorm.DB) {
	err := db.Model(&rows).Where("status=?", 26).Update("group_zone", 26).Error
	if err != nil {
		log.Fatal(err)
	}
}
