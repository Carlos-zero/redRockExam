package service

import (
	"docter/dao"
	"docter/handler"
	"docter/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DownloadService(ctx *gin.Context)  {

	token := handler.GetToken(ctx.Request)
	_,err:= handler.Decode(token)
	if err!=nil {
		ctx.Writer.Write([]byte("您的身份已过期或身份错误，请重新登录！"))
		return
	}


	fileName:=ctx.Query("file_name")
	fileTableIdS:=ctx.Query("file_table_id")
	fileTableId:=util.StringToInt(fileTableIdS)
	fileUserIdS:=ctx.Query("file_user_id")
	fileUserId:=util.StringToInt(fileUserIdS)
	fmt.Println(fileName,fileTableId,fileUserId)
	fileAddr:=dao.GetFileAddr(fileName,fileTableId,fileUserId)
	//fmt.Println("fileAddr:",fileAddr)
	handler.DownloadFile(ctx,fileAddr)

}
