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

func Conn(user, password, host, db, port string) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, db)
	db_instance, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatalf("连接mysql err:", err)
	}
	return db_instance
}

type XzAutoServerConf struct {
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
	orm_op(db)
}

func orm_op(db *gorm.DB) {
	//var rows []XzAutoServerConf
	add_data := XzAutoServerConf{GroupZone: "20", ServerName: "jason", ServerId: 811}
	db.Create(&add_data)

}
