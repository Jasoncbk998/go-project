/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/29 9:39 下午
 **/
package demo01

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-practice/go_mysql/go2gorm/util"
	"log"
)

const (
	TableName = "student"
	User      = "root"
	Password  = "Liuxu!99454"
	Host      = "127.0.0.1"
	Port      = "3306"
	Db_Name   = "test"
)

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

type Student struct {
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
	Sex  string `gorm:"column:sex"`
	High int    `gorm:"column:high"`
}

func main() {
	fmt.Println(User)
	dbset := util.DatabaseSetting{
		DBName:   Db_Name,
		UserName: User,
		Password: Password,
		Port:     Port,
		Host:     Host,
	}
	db := util.Coon(&dbset)
	db.LogMode(true)
	defer db.Close()
	data := &Student{Name: "jason", Age: 20, Sex: "男", High: 10}

	insert_data(db, data)
	//sql_test(db)

}

func insert_data(db *gorm.DB, data *Student) {
	create := db.Create(data)
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
	Num_       int64
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

func sql_test(db *gorm.DB) {
	var rows []XzAutoServerConf
	//输入 表中的字段
	db.Raw("select  group_zone ,server_id,server_name,status ,username from xz_auto_server_conf where group_zone=?", "26").Scan(&rows)
	for k, v := range rows {
		fmt.Println(k, v)
	}
}
