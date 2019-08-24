package service

import (
	"docter/dao"
	"docter/handler"
	"docter/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoginServiceByName(ctx *gin.Context) string {

	userName:=ctx.PostForm("user_name")
	password:=ctx.PostForm("password")
	fmt.Println("sadsadasdasdasdasdasdasdasd",userName,password)
	userID:=dao.UserLoginToName(userName,password)
	//fmt.Println(userID)
	if userID==0{
		ctx.Writer.Write([]byte("登录失败，请验证账号或密码！"))
		return ""
	}
	//cus:=&handler.CustomerClaim{}
	user:=handler.User{}
	user.UserId=userID
	user.UserName=userName
	user.Phone=0

	//生成token
	token,_:=handler.Encode(user)
	dao.InsertUserToken(userID,token)
	fmt.Println("token:",token)

	return token
}

func LoginServiceByPhone(ctx *gin.Context) string {
	phoneS:=ctx.PostForm("phone")
	password:=ctx.PostForm("password")
	phone:=util.StringToInt(phoneS)
	userID:=dao.UserLoginToPhone(phone,password)
	fmt.Println(userID)
	//cus:=&handler.CustomerClaim{}
	user:=handler.User{}
	user.UserId=userID
	user.UserName=""
	user.Phone=phone

	//生成token
	token,_:=handler.Encode(user)
	dao.InsertUserToken(userID,token)
	fmt.Println("token:",token)

	return token
}

func Register(ctx *gin.Context)  {
	userName:=ctx.PostForm("user_name")
	password:=ctx.PostForm("password")
	phoneS:=ctx.PostForm("phone")
	phone:=util.StringToInt(phoneS)
	userID:=dao.Register(userName,password,phone)
	dao.AddFolder(1,"我的资源",0,0,userID)
	ctx.Writer.Write([]byte("注册成功，请前往登录页面登录！"))
}

