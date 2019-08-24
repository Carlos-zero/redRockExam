package handler

import (
	"docter/util"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func DownloadFile(ctx *gin.Context,fileAddr string){
	//fsha1:=ctx.Request.FormValue("fileHash")
	//fmeta:=meta.GetFileMeta

	file,err:=os.Open(fileAddr)
	util.FailOnErr(err,"open file error!")
	if err!=nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"fileErr:":err,
		})
	}
	data,err:=ioutil.ReadAll(file)
	util.FailOnErr(err,"read file error")
	if err!=nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"dataErr:":err,
		})
	}
	fileName:=ctx.Query("file_name")
	ctx.Header("Content-Disposition", "attachment; filename=\""+fileName)
	ctx.Data(http.StatusOK,"application/octet-stream",data)
}
