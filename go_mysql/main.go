/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/29 9:05 下午
 **/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql包
	"time"
	//"time"
)

type Doctor struct {
	ID      int64
	Name    string
	Age     int
	Sex     int
	AddTime time.Time
}

func main() {
	db, err := sql.Open("mysql", "root:Liuxu!99454@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return
	}
	//延迟到函数结束关闭链接
	defer db.Close()

	var doc Doctor
	//执行单条查询
	rows := db.QueryRow("select * from doctor_tb where id = ?", 1)
	rows.Scan(&doc.ID, &doc.Name, &doc.Age, &doc.Sex, &doc.AddTime)
	fmt.Println("单条数据结果：", doc)

	result, err := db.Exec("insert into doctor_tb(name,age,sex,addTime) values(?,?,?,Now())", "刘医生", 20, 3)
	if err != nil {
		fmt.Println("新增数据错误", err)
		return
	}
	newID, _ := result.LastInsertId() //新增数据的ID
	i, _ := result.RowsAffected()     //受影响行数
	fmt.Printf("新增的数据ID：%d , 受影响行数：%d \n", newID, i)

}
