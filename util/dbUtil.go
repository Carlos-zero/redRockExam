package util

import (
	"database/sql"
	"docter/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init()  {
	db,_=sql.Open("mysql",config.MySQLSource)

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	err:=db.Ping()
	if err!=nil {
		fmt.Println("Failed to connect to MySQL,err:" + err.Error())
		os.Exit(1)
	}
}

//返回数据库连接对象
func Conn() *sql.DB  {
	return db
}

func Close() interface{}{
	return db.Close()
}


