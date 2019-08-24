package service

import (
	"docter/dao"
	"docter/handler"
	"docter/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//分享连接生成
func ShareFileService(ctx *gin.Context)  {
	token := handler.GetToken(ctx.Request)
	_,err:= handler.Decode(token)
	if err!=nil {
		ctx.Writer.Write([]byte("您的身份已过期或身份错误，请重新登录！"))
		return
	}

	fileName:=ctx.PostForm("file_name")
	fileTableIdS:=ctx.PostForm("file_table_id")
	fileTableId:=util.StringToInt(fileTableIdS)
	fileUserIdS:=ctx.PostForm("file_user_id")
	fileUserId:=util.StringToInt(fileUserIdS)
	fmt.Println(fileName,fileTableId,fileUserId)
	sharePassword:=ctx.PostForm("share_password")
	fileAddr:=dao.GetFileAddr(fileName,fileTableId,fileUserId)
	timeInt:=time.Now().Unix()
	timeStr:=strconv.FormatInt(timeInt,10)
	path:="/share/"
	shareUri:=path+timeStr

	dao.AddShare(fileAddr,sharePassword,shareUri)
	ctx.JSON(http.StatusOK,gin.H{
		"shareUri":shareUri,
	})
}

//权限验证
func AuthShareService(ctx *gin.Context)  {
	token := handler.GetToken(ctx.Request)
	_,err:= handler.Decode(token)
	if err!=nil {
		ctx.Writer.Write([]byte("您的身份已过期或身份错误，请重新登录！"))
		return
	}

	uris:=ctx.Request.RequestURI
	uri:=strings.Split(uris,"?")[0]
	fmt.Println(uri)
	password:=ctx.Query("share_password")
	res:=dao.AuthShare(password,uri)
	if res=="提取码错误！" {
		ctx.Writer.Write([]byte("提取码错误!"))
	}else {
		handler.DownloadFile(ctx,res)
	}

}
