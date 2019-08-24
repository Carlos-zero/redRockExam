package service

import (
	"docter/dao"
	"docter/handler"
	"docter/util"
	"github.com/gin-gonic/gin"
)

//添加好友
func AddFriendService(ctx *gin.Context)  {
	token := handler.GetToken(ctx.Request)
	_,err:= handler.Decode(token)
	if err!=nil {
		ctx.Writer.Write([]byte("您的身份已过期或身份错误，请重新登录！"))
		return
	}
	Aid:=util.StringToInt(ctx.PostForm("id"))
	Bid:=util.StringToInt(ctx.PostForm("BName"))

	dao.AddFriend(Aid,Bid)
}

//给予好友权限
func GiveRoleService(ctx *gin.Context)  {
	token := handler.GetToken(ctx.Request)
	_,err:= handler.Decode(token)
	if err!=nil {
		ctx.Writer.Write([]byte("您的身份已过期或身份错误，请重新登录！"))
		return
	}
	Aid:=util.StringToInt(ctx.PostForm("id"))
	Bid:=util.StringToInt(ctx.PostForm("BName"))
	role:=util.StringToInt("role")
	dao.GiveRole(Aid,Bid,role)
}

//删除好友
func DeleteFriendService(ctx *gin.Context)  {
	token := handler.GetToken(ctx.Request)
	_,err:= handler.Decode(token)
	if err!=nil {
		ctx.Writer.Write([]byte("您的身份已过期或身份错误，请重新登录！"))
		return
	}
	Aid:=util.StringToInt(ctx.PostForm("id"))
	Bid:=util.StringToInt(ctx.PostForm("BName"))
	dao.DeleteFriend(Aid,Bid)
}
