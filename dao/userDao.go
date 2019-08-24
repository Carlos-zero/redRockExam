package dao

import (
	"docter/util"
)

func UserLoginToName(userName string,password string) int {
	db:=util.Conn()


	var userId int
	//rows,err:=db.Query("select id from user where user_name=? and password=?",userName,password)
	//fmt.Println("???????????",userName,password)
	err:=db.QueryRow("select id from user where user_name=? and password=?",userName,password).Scan(&userId)
	util.FailOnErr(err,"查询数据失败！UserLoginToName")

	/*for rows.Next() {
		err = rows.Scan(&userId)
		util.FailOnErr(err,"写入数据失败！")
	}*/
	return userId
}

func UserLoginToPhone(phone int,password string) int {
	db:=util.Conn()

	var userId int
	//rows,err:=db.Query("select id from user where phone=? and password=?",phone,password)
	err:=db.QueryRow("select id from user where phone=? and password=?",phone,password).Scan(&userId)
	util.FailOnErr(err,"查询数据失败！UserLoginToPhone")

	/*for rows.Next() {
		err = rows.Scan(&userId)
		util.FailOnErr(err,"写入数据失败！")
	}*/
	return userId
}

func Register(userName string,password string,phone int) int64 {
	db := util.Conn()

	new,err:=db.Exec("insert into user(user_name,password,phone)value (?,?,?)",userName,password,phone)
	util.FailOnErr(err,"注入失败！")

	userID,err:=new.LastInsertId()
	util.FailOnErr(err,"注册失败！")
	return userID
}

