package service

import (
	"docter/dao"
	"docter/handler"
	"docter/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func UploadService(ctx *gin.Context)  {
	token := handler.GetToken(ctx.Request)
	_,err:= handler.Decode(token)
	if err!=nil {
		ctx.Writer.Write([]byte("您的身份已过期或身份错误，请重新登录！"))
		return
	}

	fileNames,fileSizes:=handler.UploadFile(ctx)
	fileTableIdS:=ctx.PostForm("file_table_num")
	fmt.Println("file_table_num",fileTableIdS)
	fileTableId:=util.StringToInt(fileTableIdS)
	fileUserIdS:=ctx.PostForm("file_user_id")
	fileUserId:=util.StringToInt(fileUserIdS)
	filePreFolderId:=util.StringToInt(ctx.PostForm("file_pre_folder_id"))
	for index:=range fileNames{
		fileName:=fileNames[index]
		fileSize:=fileSizes[index]
		fileAddr:="C:\\test\\"+fileName
		uploadTime:=time.Now()
		//fmt.Println(uploadTime)
		dao.FileUploadFished(fileAddr,fileSize,fileName,fileTableId,fileUserId,uploadTime,filePreFolderId)
	}
	//util.Close()
}
