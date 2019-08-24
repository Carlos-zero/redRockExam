package dao

import "docter/util"

//添加好友
func AddFriend(Aid int,Bid int)  {
	db:=util.Conn()

	_,err:=db.Exec("insert into friend(a_id,b_id,aTob_role)value (?,?,2)",Aid,Bid)
	util.FailOnErr(err,"注入失败！")
	_,err=db.Exec("insert into friend(a_id,b_id,aTob_role)value (?,?,2)",Bid,Aid)

	util.FailOnErr(err,"注册失败！")
}

//给好友设置权限
func GiveRole(Aid int,Bid int,role int)  {
	db:=util.Conn()
	_,err:=db.Exec("update friend set aTob_role=? where a_id=? and b_id=?",role,Aid,Bid)
	util.FailOnErr(err, "更改权限失败！")
}

//删除好友
func DeleteFriend(Aid int,Bid int){
	db:=util.Conn()

	_,err:=db.Exec("delete from friend where a_id=? and b_id=?",Aid,Bid)
	util.FailOnErr(err,"删除好友失败！")
}

//查看权限
func QueryRole(Aid int,Bid int)int{
	db:=util.Conn()

	var role int
	err:=db.QueryRow("select aTob_role from friend where a_id=? and b_id=?",Aid,Bid).Scan(&role)
	if err!=nil {
		return -1
	}

	return role
}
