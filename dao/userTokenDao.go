package dao

import "docter/util"

func UpdateUserToken(userID int, token string) {
	db := util.Conn()

	_, err := db.Exec("update token set token=? where user_id=?", token, userID)
	util.FailOnErr(err, "gengxin token shi bai!")
}

func InsertUserToken(userID int, token string) {
	db := util.Conn()

	_, err := db.Exec("insert into token(token,user_id)value (?,?)", token, userID)
	util.FailOnErr(err, "插入token失败！")
}

func SelectUserToken(userID int) string {
	db := util.Conn()

	rows, err := db.Query("select token from token where user_id=?", userID)

	var token string
	for rows.Next() {
		err = rows.Scan(&token)
		util.FailOnErr(err, "写入数据失败！")
	}
	return token
}
