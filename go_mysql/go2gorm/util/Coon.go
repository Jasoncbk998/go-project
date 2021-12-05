/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/2 11:02 上午
 **/
package util

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql包
	"github.com/jinzhu/gorm"
	"log"
)

type DatabaseSetting struct {
	UserName string
	Password string
	Host     string
	DBName   string
	Port     string
	Driver   string
}

func Coon(dataSet *DatabaseSetting) *gorm.DB {

	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dataSet.UserName, dataSet.Password, dataSet.Host, dataSet.Port, dataSet.DBName)
	instance, err := gorm.Open(dataSet.Driver, connArgs)
	if err != nil {
		log.Fatal(err)
	}
	//查询单表， 禁止复数表存在
	instance.SingularTable(true)
	instance.LogMode(true)
	return instance
}
