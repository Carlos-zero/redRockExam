package dao

import "docter/util"

func AddShare(shareAddr string,sharePassword string,shareUri string)  {
	db:=util.Conn()

	_,err:=db.Exec("insert into share(share_addr,share_password,share_uri)value (?,?,?)",shareAddr,sharePassword,shareUri)
	util.FailOnErr(err,"注入失败！")
}

func AuthShare(sharePassword string,shareUri string) string {
	db:=util.Conn()

	var addr string
	err:=db.QueryRow("select share_addr from share where share_uri=? and share_password=?",shareUri,sharePassword)
	if err!=nil {
		return "提取码错误！"
	}else {
		return addr
	}
}
